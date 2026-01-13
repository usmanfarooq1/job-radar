package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/usmanfarooq1/job-radar/internal/common/decorator"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type RunTask struct {
	TaskId uuid.UUID
}

type RunTaskHandler decorator.CommandHandler[RunTask]

type runTaskHandler struct {
	engine   engine.Engine
	logger   zerolog.Logger
	taskRepo engine.ScraperTaskRepository
}

func NewRunTaskHandler(
	engine engine.Engine,
	logger zerolog.Logger,
	taskRepo engine.ScraperTaskRepository,
) runTaskHandler {
	return runTaskHandler{
		engine:   engine,
		logger:   logger,
		taskRepo: taskRepo,
	}
}

func (h runTaskHandler) Handle(ctx context.Context, cmd RunTask) error {
	manager := h.engine.Manager()
	err := manager.ExecuteScraperTask(cmd.TaskId)
	if err != nil {
		h.logger.Err(err)
		return err
	}

	return nil
}
