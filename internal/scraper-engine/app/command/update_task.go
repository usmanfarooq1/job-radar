package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/usmanfarooq1/job-radar/internal/common/decorator"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type UpdateTask struct {
	TaskId         uuid.UUID
	DelayInSeconds uint32
	SearchKeyword  string
	LocationId     string
	DistanceRadius string
	TaskLocation   string
}

type UpdateTaskHandler decorator.CommandHandler[UpdateTask]

type updateTaskHandler struct {
	engine   engine.Engine
	logger   zerolog.Logger
	taskRepo engine.ScraperTaskRepository
}

func NewUpdateTaskHandler(
	engine engine.Engine,
	logger zerolog.Logger,
	taskRepo engine.ScraperTaskRepository,
) updateTaskHandler {
	return updateTaskHandler{
		engine:   engine,
		taskRepo: taskRepo,
	}
}

func (h updateTaskHandler) Handle(ctx context.Context, cmd UpdateTask) error {
	manager := h.engine.Manager()
	task, err := manager.UpdateScraperTask(cmd.TaskId, cmd.DelayInSeconds, cmd.SearchKeyword, cmd.LocationId, cmd.DistanceRadius, cmd.TaskLocation)
	if err != nil {
		h.logger.Err(err)
		return err
	}
	task.Execute()
	return nil
}
