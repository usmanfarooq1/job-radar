package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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

type UpdateTaskHandler decorator.CommandHandler[AddTask]

type updateTaskHandler struct {
	engine   engine.Engine
	taskRepo engine.ScraperTaskRepository
}

func NewUpdateTaskHandler(
	engine engine.Engine,
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
		log.Err(err)
		return err
	}
	task.Execute()
	return nil
}
