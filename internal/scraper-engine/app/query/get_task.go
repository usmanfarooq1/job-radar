package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/usmanfarooq1/job-radar/internal/common/decorator"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type TaskQuery struct {
	TaskId uuid.UUID
}

type GetTaskHandler decorator.QueryHandler[TaskQuery, Task]

type getTaskHandler struct {
	taskRepo engine.ScraperTaskRepository
}

func NewGetTaskkHandler(
	engine engine.Engine,
	taskRepo engine.ScraperTaskRepository,
) getTaskHandler {
	return getTaskHandler{

		taskRepo: taskRepo,
	}
}
func (h getTaskHandler) Handle(ctx context.Context, cmd TaskQuery) (Task, error) {
	_, err := h.taskRepo.GetScraperTask(ctx, cmd.TaskId)
	if err != nil {
		log.Err(err)
		return Task{}, err
	}

	return Task{}, nil
}
