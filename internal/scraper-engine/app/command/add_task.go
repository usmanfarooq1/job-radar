package command

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/usmanfarooq1/job-radar/internal/common/decorator"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type AddTask struct {
	DelayInSeconds uint32
	SearchKeyword  string
	LocationId     string
	TaskType       string
	DistanceRadius string
	TaskLocation   string
}

type AddTaskHandler decorator.CommandHandler[AddTask]

type addTaskHandler struct {
	engine   engine.Engine
	taskRepo engine.ScraperTaskRepository
}

func NewAddTaskHandler(
	engine engine.Engine,
	taskRepo engine.ScraperTaskRepository,
) addTaskHandler {
	return addTaskHandler{
		engine:   engine,
		taskRepo: taskRepo,
	}
}

func (h addTaskHandler) Handle(ctx context.Context, cmd AddTask) error {
	manager := h.engine.Manager()
	task, err := engine.MakeTask(cmd.DelayInSeconds, cmd.SearchKeyword, cmd.LocationId, cmd.TaskType, cmd.DistanceRadius, cmd.TaskLocation)
	if err != nil {
		log.Err(err)
		return err
	}
	task, err = manager.AddScraperTask(*task)
	if err != nil {
		log.Err(err)
		return err
	}

	h.taskRepo.AddScraperTask(ctx, task)
	fmt.Println(task)
	return nil
}
