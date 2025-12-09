package query

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/usmanfarooq1/job-radar/internal/common/decorator"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type ListTasksQuery struct{}

type ListTasksHandler decorator.QueryHandler[ListTasksQuery, []Task]

type listTasksHandler struct {
	taskRepo engine.ScraperTaskRepository
}

func NewListTasksHandler(
	taskRepo engine.ScraperTaskRepository,
) listTasksHandler {
	return listTasksHandler{
		taskRepo: taskRepo,
	}
}
func (h listTasksHandler) Handle(ctx context.Context, cmd ListTasksQuery) ([]Task, error) {
	_, err := h.taskRepo.ListScraperTasks(ctx)
	if err != nil {
		log.Err(err)
		return []Task{}, err
	}

	return []Task{}, nil
}
