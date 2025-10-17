package engine

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type ScraperTaskType string

const (
	LinkedIn ScraperTaskType = "linkedin"
)

// type ScraperTaskId uuid.UUID

// func (s ScraperTaskId) IsValidId(scraperTaskId string) uuid.UUID {
// 	uuid, err := uuid.Parse(scraperTaskId)
// 	if err != nil {

// 	}
// 	return ScraperTaskId(uuid)
// }

type ScraperTask struct {
	/*
		 Task will have a certain state which will be the main execution context for the go routine to process,
		 it contains the parameters for the searching and the status what is happening regarding the task.
		106430557
	*/
	id               uuid.UUID
	delayInSeconds   uint32
	searchKeyword    string
	taskLocationId   string
	distanceRadius   uint8
	taskLocation     string
	taskType         ScraperTaskType
	exectionHandler  ExecutionStrategy
	executionChannel chan (bool)
}

func ParseTaskType(in string) (ScraperTaskType, error) {
	switch in {
	case strings.ToLower(LinkedIn.String()):
		return LinkedIn, nil
	default:
		return LinkedIn, ErrInvalidTaskType
	}

}

func (tt ScraperTaskType) String() string {
	switch tt {
	case "linkedin":
		return "LinkedIn"
	default:
		return ""
	}

}

func (t *ScraperTask) isValidDelay(delayInSeconds uint32) error {
	if delayInSeconds < 1800 {
		return ErrInvalidDelay
	}
	return nil
}

func (t *ScraperTask) SetTaskType(taskType string) error {
	taskTypeEnum, err := ParseTaskType(strings.ToLower(taskType))
	if err != nil {
		return err
	}
	t.taskType = taskTypeEnum
	return nil
}

func (t *ScraperTask) SetDelay(delayInSeconds uint32) error {
	if err := t.isValidDelay(delayInSeconds); err != nil {
		return err
	}
	t.delayInSeconds = delayInSeconds
	return nil
}

func (t *ScraperTask) isValidRadiusDistance(distanceInString string) (*uint8, error) {
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

func (t *ScraperTask) SetDistance(distanceInString string) error {
	distance, err := t.isValidRadiusDistance(distanceInString)
	if err != nil {
		return err
	}
	t.distanceRadius = *distance
	return nil
}

func (t *ScraperTask) SetSearchKeywords(search string) error {
	search = strings.Trim(search, " ")
	if search == "" {
		return ErrEmptySearchKeyword
	}
	t.searchKeyword = search
	return nil
}

func (t *ScraperTask) SetTaskLocation(location string) error {
	location = strings.Trim(location, " ")
	if location == "" {
		return ErrEmptyTaskLocation
	}
	t.taskLocation = location
	return nil
}
func (t *ScraperTask) StopExecution() {
	t.executionChannel <- true
}

func (t ScraperTask) Execute() {
	go t.exectionHandler.JobExtractor(&t)
}

func (t *ScraperTask) generateExecutionChannel() {
	t.executionChannel = make(chan bool)
}

func (t *ScraperTask) Equal(task ScraperTask) bool {
	return t.taskLocationId == task.taskLocationId &&
		t.taskLocation == task.taskLocation &&
		t.searchKeyword == task.searchKeyword &&
		t.taskType == task.taskType
}

func (t *ScraperTask) SetTaskLocationId(locationId string) error {
	locationId = strings.Trim(locationId, " ")
	if locationId == "" {
		return ErrEmptyTaskLocationId
	}
	var re = regexp.MustCompile(`^[0-9]+$`)
	if !re.MatchString(locationId) {
		return ErrInvalidTaskLocationId
	}
	t.taskLocationId = locationId
	return nil
}
func MakeTask(
	delayInSeconds uint32,
	searchKeyword string,
	locationId string,
	taskType string,
	distanceRadius string,
	taskLocation string) (*ScraperTask, error) {
	task := ScraperTask{}
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
	if err := task.SetTaskType(taskType); err != nil {
		return nil, err
	}
	task.generateExecutionChannel()
	handler, err := GenerateExecutionStrategy(&task)
	if err != nil {
		return nil, err
	}
	task.exectionHandler = handler
	task.id = uuid.New()

	return &task, nil
}
