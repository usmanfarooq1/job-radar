package engine

import (
	"log"

	"github.com/google/uuid"
	"github.com/playwright-community/playwright-go"
)

type Manager struct {
	scraperTasks map[uuid.UUID]ScraperTask
	pBrowser     *playwright.Browser

	/*
		The Manager contains the list of Tasks and it will contain the behaviour for
		- adding a task.
		- removing a task.
		- stoping a task.
	*/
}

func MakeManager() Manager {
	scraperList := make(map[uuid.UUID]ScraperTask)
	pw, err := playwright.Run()

	if err != nil {
		log.Fatalf("can't start playwright : %v", err)
	}
	// TODO Add the environment variable here
	browser, err := pw.Chromium.Connect("ws://127.0.0.1:3000/")
	if err != nil {
		log.Fatalf("can't start chromium  : %v", err)
	}

	return Manager{scraperTasks: scraperList, pBrowser: &browser}
}

func (m *Manager) getScraperTask(taskId uuid.UUID) *ScraperTask {
	task, ok := m.scraperTasks[taskId]
	if ok {
		return &task
	}
	return nil
}

func (m *Manager) AddScraperTask(task ScraperTask) (*ScraperTask, error) {
	task.SetPBrowser(m.pBrowser)
	t, ok := m.scraperTasks[task.id]
	if !ok {
		m.scraperTasks[task.id] = task
	}
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
	handler, err := GenerateExecutionStrategy(task)
	if err != nil {
		return nil, err
	}
	task.exectionHandler = handler
	task.Execute()
	return task, nil
}
