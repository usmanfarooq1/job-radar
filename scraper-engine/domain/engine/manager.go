package engine

import (
	"github.com/google/uuid"
)

type Manager struct {
	scraperTasks []ScraperTask

	/*
		The Manager contains the list of Tasks and it will contain the behaviour for
		- adding a task.
		- removing a task.
		- stoping a task.
	*/
}

func MakeManager() Manager {
	scraperList := make([]ScraperTask, 0)
	return Manager{scraperTasks: scraperList}
}

func (m *Manager) getScraperTask(taskId uuid.UUID) *ScraperTask {
	for _, scraperTask := range m.scraperTasks {
		if taskId == scraperTask.id {
			return &scraperTask
		}
	}

	return nil
}

func (m *Manager) AddScraperTask(delayInSeconds uint32,
	searchKeyword string,
	locationId string,
	taskType string,
	distanceRadius string,
	taskLocation string) (*ScraperTask, error) {
	task, err := MakeTask(delayInSeconds,
		searchKeyword,
		locationId,
		taskType,
		distanceRadius,
		taskLocation)
	if err != nil {
		return nil, err
	}
	m.scraperTasks = append(m.scraperTasks, *task)
	return task, nil
}

func (m *Manager) StopScraperTask(taskId uuid.UUID) error {
	task := m.getScraperTask(taskId)
	if task == nil {
		return ErrNotFound
	}
	task.StopExecution()
	return nil
}

func (m *Manager) UpdateScraperTask(
	taskId uuid.UUID,
	delayInSeconds uint32,
	searchKeyword string,
	locationId string,
	distanceRadius uint8,
	taskLocation string) (*ScraperTask, error) {
	task := m.getScraperTask(taskId)
	if task == nil {
		return nil, ErrNotFound
	}
	if err := task.SetDelay(delayInSeconds); err != nil {
		return nil, err
	}
	task.searchKeyword = searchKeyword
	task.taskLocationId = locationId
	task.distanceRadius = distanceRadius
	task.taskLocation = taskLocation
	task.StopExecution()
	task.generateExecutionChannel()
	handler, err := GenerateExecutionStrategy(task)
	if err != nil {
		return nil, err
	}
	task.exectionHandler = handler
	return task, nil
}
