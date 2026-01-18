package query

import (
	"context"
	"fmt"
	"strconv"
	"time"

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
	var tasks []Task
	for _, task := range list {
		tasks = append(tasks, Task{
			TaskId:         task.Id().String(),
			TaskType:       task.TaskType().String(),
			DelayInSeconds: task.DelayInSeconds(),
			SearchKeyword:  task.SearchKeyword(),
			LocationId:     task.LocationId(),
			DistanceRadius: strconv.Itoa(int(task.DistanceRadius())),
			TaskLocation:   task.TaskLocation(),
			CreatedAt:      time.Now().String(),
			UpdatedAt:      time.Now().String(),
		})
	}
	return tasks, nil
}
