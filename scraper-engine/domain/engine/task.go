package engine

import (
	"strconv"
	"strings"
)

type TaskType string

const (
	LinkedIn TaskType = "linkedIn"
)

type Task struct {
	/*
		 Task will have a certain state which will be the main execution context for the go routine to process,
		 it contains the parameters for the searching and the status what is happening regarding the task.
		106430557
	*/
	delayInSeconds   uint32
	searchKeyword    string
	taskLocationId   string
	distanceRadius   uint8
	taskLocation     string
	taskType         TaskType
	exectionHandler  ExecutionStrategy
	executionChannel chan (bool)
}

func ParseTaskType(in string) (TaskType, error) {
	switch in {
	case LinkedIn.String():
		return LinkedIn, nil
	}
	return LinkedIn, ErrInvalidTaskType
}

func (tt TaskType) String() string {
	switch tt {
	case "linkedIn":
		return "LinkedIn"
	}
	return ""
}

func (t Task) isValidDelay(delayInSeconds uint32) error {
	if delayInSeconds < 1800 {
		return ErrInvalidDelay
	}
	return nil
}

func (t Task) SetTaskType(taskType string) error {
	taskTypeEnum, err := ParseTaskType(strings.ToLower(taskType))
	if err != nil {
		return err
	}
	t.taskType = taskTypeEnum
	return nil
}

func (t Task) SetDelay(delayInSeconds uint32) error {
	if err := t.isValidDelay(delayInSeconds); err != nil {
		return err
	}
	t.delayInSeconds = delayInSeconds
	return nil
}

func (t Task) isValidRadiusDistance(distanceInString string) (*uint8, error) {
	convertedDistance, err := strconv.Atoi(distanceInString)
	if err != nil {
		return nil, ErrInvalidDistanceType
	}
	if convertedDistance < 25 || convertedDistance > 100 {
		return nil, ErrInvalidDistance
	}
	distance := uint8(convertedDistance)
	return &distance, nil
}

func (t Task) SetDistance(distanceInString string) error {
	distance, err := t.isValidRadiusDistance(distanceInString)
	if err != nil {
		return err
	}
	t.distanceRadius = *distance
	return nil
}

func (t Task) StopExecution() {
	t.executionChannel <- true
}

func (t Task) Execute() {

}

func MakeTask(delayInSeconds uint32,
	searchKeyword string,
	locationId string,
	taskType string,
	distanceRadius string,
	taskLocation string) (*Task, error) {
	task := Task{
		searchKeyword:  searchKeyword,
		taskLocationId: locationId,
		taskLocation:   taskLocation,
	}
	if err := task.SetDistance(distanceRadius); err != nil {
		return nil, err
	}
	if err := task.SetDelay(delayInSeconds); err != nil {
		return nil, err
	}
	if err := task.SetTaskType(taskType); err != nil {
		return nil, err
	}
	task.executionChannel = make(chan bool)
	handler, err := GenerateExecutionStrategy(&task)
	if err != nil {
		return nil, err
	}
	task.exectionHandler = handler
	return &task, nil
}
