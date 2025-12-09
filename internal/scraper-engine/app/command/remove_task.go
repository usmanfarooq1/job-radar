package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/usmanfarooq1/job-radar/internal/common/decorator"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type RemoveTask struct {
	TaskId uuid.UUID
}

type RemoveTaskHandler decorator.CommandHandler[RemoveTask]

type removeTaskHandler struct {
	engine   engine.Engine
	taskRepo engine.ScraperTaskRepository
}

func NewRemoveTaskHandler(
	engine engine.Engine,
	taskRepo engine.ScraperTaskRepository,
) removeTaskHandler {
	return removeTaskHandler{
		engine:   engine,
		taskRepo: taskRepo,
	}
}

func (h removeTaskHandler) Handle(ctx context.Context, cmd RemoveTask) error {
	manager := h.engine.Manager()
	err := manager.RemoveScraperTask(cmd.TaskId)
	if err != nil {
		log.Err(err)
		return err
	}

	return nil
}
