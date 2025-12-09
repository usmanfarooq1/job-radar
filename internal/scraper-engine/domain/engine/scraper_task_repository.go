package engine

import (
	"context"

	"github.com/google/uuid"
)

type ScraperTaskRepository interface {
	AddScraperTask(ctx context.Context, st *ScraperTask) (*ScraperTask, error)
	UpdateScraperTask(ctx context.Context, st *ScraperTask) (*ScraperTask, error)
	RemoveScraperTask(ctx context.Context, id uuid.UUID) error
	GetScraperTask(ctx context.Context, id uuid.UUID) (*ScraperTask, error)
	ListScraperTasks(ctx context.Context) ([]ScraperTask, error)
}
