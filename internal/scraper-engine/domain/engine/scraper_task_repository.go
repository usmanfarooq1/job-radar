package engine

import (
	"context"

	"github.com/google/uuid"
)

type ScraperTaskRepository interface {
	AddScraperTask(ctx context.Context, st *ScraperTask, exec func() error) (*ScraperTask, error)
	UpdateScraperTask(ctx context.Context, st *ScraperTask, exec func() error) (*ScraperTask, error)
	RemoveScraperTask(ctx context.Context, id uuid.UUID, exec func() error) error
	GetScraperTask(ctx context.Context, id uuid.UUID) (*ScraperTask, error)
	ListScraperTasks(ctx context.Context) ([]ScraperTask, error)
}

type ExecutionStrategy interface {
	JobExtractor(t *ScraperTask)
}
