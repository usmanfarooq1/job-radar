package engine

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/playwright-community/playwright-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/usmanfarooq1/job-radar/internal/common/mq"
)

type Manager struct {
	scraperTasks   map[uuid.UUID]ScraperTask
	pBrowser       playwright.Browser
	logger         zerolog.Logger
	mq             ScraperTaskPublishRepository
	pubKillChannel chan (int)
	/*
		The Manager contains the list of Tasks and it will contain the behaviour for
		- adding a task.
		- removing a task.
		- stoping a task.
	*/
}

func MakeManager(mq ScraperTaskPublishRepository, logger zerolog.Logger) Manager {
	scraperList := make(map[uuid.UUID]ScraperTask)
	driver, err := playwright.NewDriver(&playwright.RunOptions{
		SkipInstallBrowsers: true,
	},
	)
	if err != nil {
		fmt.Println("unable to create playwright driver setting object")
	}
	err = driver.Install()
	if err != nil {
		fmt.Println("unable to install playwright drivers for communication")
	}
	pw, err := playwright.Run()

	if err != nil {
		log.Err(err).Msg("can't start playwright")
	}

	browser, err := pw.Chromium.Connect(os.Getenv("PLAYWRIGHT_CONNECTION_STRING"))
	if err != nil {
		log.Err(err).Msg("can't connect to chromium")
	}

	return Manager{scraperTasks: scraperList, pBrowser: browser, logger: logger, mq: mq, pubKillChannel: make(chan int)}
}

func (m *Manager) fanInToPublish() <-chan mq.JobLinkMessagePayload {
	out := make(chan mq.JobLinkMessagePayload)
	for _, st := range m.scraperTasks {
		go func(c <-chan mq.JobLinkMessagePayload) {
			for val := range c {
				out <- val
			}
		}(st.ResultChannel())
	}
	return out
}

func (m *Manager) publishMessages() {
	merged := m.fanInToPublish()
	_, ok := <-m.pubKillChannel
	for ok {
		if len(merged) > 0 {
			message := <-merged
			m.mq.Publish(context.Background(), message)
		}
	}
}

func (m *Manager) reCreatePubKillChannel() {
	m.pubKillChannel = make(chan int)
}

func (m *Manager) killPublisher() {
	_, ok := <-m.pubKillChannel
	if ok {
		close(m.pubKillChannel)
	}
}

func (m *Manager) startPublishing() {
	m.killPublisher()
	m.reCreatePubKillChannel()
	go m.publishMessages()
}
func (m *Manager) getScraperTask(taskId uuid.UUID) *ScraperTask {
	task, ok := m.scraperTasks[taskId]
	if ok {
		return &task
	}
	return nil
}

func (m *Manager) AddScraperTask(task ScraperTask) (*ScraperTask, error) {
	t, ok := m.scraperTasks[task.id]
	if !ok {
		m.scraperTasks[task.id] = task
	}
	m.startPublishing()
	return &t, nil
}

func (m *Manager) GetManagerTasksCount() int {
	return len(m.scraperTasks)
}

func (m *Manager) StopScraperTask(taskId uuid.UUID) error {
	task := m.getScraperTask(taskId)
	if task == nil {
		return ErrNotFound
	}
	go task.StopExecution()
	m.startPublishing()
	return nil
}
func (m *Manager) RemoveScraperTask(taskId uuid.UUID) error {
	task := m.getScraperTask(taskId)
	if task == nil {
		return ErrNotFound
	}
	go task.StopExecution()
	delete(m.scraperTasks, taskId)
	return nil
}

func (m *Manager) ExecuteScraperTask(taskId uuid.UUID) error {
	task := m.getScraperTask(taskId)
	if task == nil {
		return ErrNotFound
	}
	task.Execute()
	return nil
}

func (m *Manager) UpdateScraperTask(
	taskId uuid.UUID,
	delayInSeconds uint32,
	searchKeyword string,
	locationId string,
	distanceRadius string,
	taskLocation string) (*ScraperTask, error) {
	task := m.getScraperTask(taskId)
	if task == nil {
		return nil, ErrNotFound
	}
	if err := task.SetDelay(delayInSeconds); err != nil {
		return nil, err
	}
	if err := task.SetSearchKeywords(searchKeyword); err != nil {
		return nil, err
	}
	if err := task.SetTaskLocation(taskLocation); err != nil {
		return nil, err
	}
	if err := task.SetTaskLocationId(locationId); err != nil {
		return nil, err
	}
	if err := task.SetDistance(distanceRadius); err != nil {
		return nil, err
	}
	if err := task.SetDelay(delayInSeconds); err != nil {
		return nil, err
	}
	task.StopExecution()
	task.generateExecutionChannel()
	task.Execute()
	m.startPublishing()
	return task, nil
}
