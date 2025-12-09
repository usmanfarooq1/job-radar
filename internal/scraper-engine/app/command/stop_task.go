package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/usmanfarooq1/job-radar/internal/common/decorator"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type StopTask struct {
	TaskId uuid.UUID
}

type StopTaskHandler decorator.CommandHandler[StopTask]

type stopTaskHandler struct {
	engine   engine.Engine
	taskRepo engine.ScraperTaskRepository
}

func NewStopTaskHandler(
	engine engine.Engine,
	taskRepo engine.ScraperTaskRepository,
) stopTaskHandler {
	return stopTaskHandler{
		engine:   engine,
		taskRepo: taskRepo,
	}
}

func (h stopTaskHandler) Handle(ctx context.Context, cmd StopTask) error {
	manager := h.engine.Manager()
	err := manager.StopScraperTask(cmd.TaskId)
	if err != nil {
		log.Err(err)
		return err
	}

	return nil
}
