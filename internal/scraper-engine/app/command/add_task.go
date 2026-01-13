package command

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/usmanfarooq1/job-radar/internal/common/decorator"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/adapters"
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
	logger   zerolog.Logger
	taskRepo engine.ScraperTaskRepository
}

func NewAddTaskHandler(
	engine engine.Engine,
	logger zerolog.Logger,
	taskRepo engine.ScraperTaskRepository,
) addTaskHandler {
	return addTaskHandler{
		engine:   engine,
		logger:   logger,
		taskRepo: taskRepo,
	}
}

func (h addTaskHandler) transaction(task *engine.ScraperTask) func() error {
	manager := h.engine.Manager()
	return func() error {
		strategy, err := adapters.GenerateExecutionStrategy(task)
		if err != nil {
			return err
		}
		task.SetExecutionStrategy(strategy)
		_, err = manager.AddScraperTask(*task)
		if err != nil {
			h.logger.Err(err).Msg("unable to add the scraper task to the manager")
			return err
		}

		return nil
	}
}

func (h addTaskHandler) Handle(ctx context.Context, cmd AddTask) error {
	task, err := engine.MakeTask(cmd.DelayInSeconds, cmd.SearchKeyword, cmd.LocationId, cmd.TaskType, cmd.DistanceRadius, cmd.TaskLocation)
	if err != nil {
		h.logger.Err(err).Msg("unable to create scraper task")
		return err
	}
	_, err = h.taskRepo.AddScraperTask(ctx, task, h.transaction(task))
	if err != nil {
		h.logger.Err(err).Msg("unble to persist the scraper task")
		return err
	}
	return nil
}
