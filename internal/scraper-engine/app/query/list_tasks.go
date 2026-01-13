package query

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/usmanfarooq1/job-radar/internal/common/decorator"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type ListTasksQuery struct{}

type ListTasksHandler decorator.QueryHandler[ListTasksQuery, []Task]

type listTasksHandler struct {
	taskRepo engine.ScraperTaskRepository
	logger   zerolog.Logger
}

func NewListTasksHandler(
	taskRepo engine.ScraperTaskRepository,
	logger zerolog.Logger,
) listTasksHandler {
	return listTasksHandler{
		taskRepo: taskRepo,
		logger:   logger,
	}
}
func (h listTasksHandler) Handle(ctx context.Context, cmd ListTasksQuery) ([]Task, error) {
	list, err := h.taskRepo.ListScraperTasks(ctx)
	h.logger.Info().Msg(fmt.Sprintf("Data %d", len(list)))
	if err != nil {
		h.logger.Err(err)
		return []Task{}, err
	}

	return []Task{}, nil
}
