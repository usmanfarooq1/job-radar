package adapters

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/usmanfarooq1/job-radar/internal/common/db"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type SQLScraperTaskRepository struct {
	db      *pgx.Conn
	queries *db.Queries
}

func NewSQLScraperTaskRepository(dbConn *pgx.Conn) *SQLScraperTaskRepository {
	if dbConn == nil {
		fmt.Println("unable to connect to database")
	}
	queries := db.New(dbConn)
	return &SQLScraperTaskRepository{db: dbConn, queries: queries}
}
func (r SQLScraperTaskRepository) AddScraperTask(ctx context.Context, st *engine.ScraperTask) (*engine.ScraperTask, error) {

	_, err := r.queries.CreateTask(ctx, db.CreateTaskParams{
		TaskID:         st.Id(),
		SearchLocation: st.SearchLocation(),
		LocationID:     st.LocationId(),
		DelayInSeconds: st.DelayInSeconds(),
		TaskState:      db.TaskStateEnum(st.TaskStatus()),
		SearchKeyword:  st.SearchKeyword(),
		DistanceRadius: st.DistanceRadius(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	if err != nil {
		return nil, errors.New("unable to create a scraper task")
	}
	return st, nil
}

func (r SQLScraperTaskRepository) UpdateScraperTask(ctx context.Context, st *engine.ScraperTask) (*engine.ScraperTask, error) {
	err := r.queries.UpdateTask(ctx, db.UpdateTaskParams{
		TaskID:         st.Id(),
		SearchLocation: st.SearchLocation(),
		LocationID:     st.LocationId(),
		DelayInSeconds: st.DelayInSeconds(),
		TaskState:      db.TaskStateEnum(st.TaskStatus()),
		SearchKeyword:  st.SearchKeyword(),
		DistanceRadius: st.DistanceRadius(),
		UpdatedAt:      time.Now(),
	})
	if err != nil {
		return nil, errors.New("unable to update scraper task")
	}
	return st, nil
}
func (r SQLScraperTaskRepository) RemoveScraperTask(ctx context.Context, id uuid.UUID) error {
	err := r.queries.DeleteTask(ctx, id)
	if err != nil {
		return errors.New("unable to delete scraper task")
	}
	return nil
}
func (r SQLScraperTaskRepository) GetScraperTask(ctx context.Context, id uuid.UUID) (*engine.ScraperTask, error) {
	task, err := r.queries.GetTask(ctx, id)
	if err != nil {
		return nil, errors.New("unable to fetch a task")
	}

	scraperTask, err := engine.UnmarshallTaskFromDatabase(task)
	if err != nil {
		return nil, errors.New("unable unmarshall task from database")
	}
	return scraperTask, nil
}
func (r SQLScraperTaskRepository) ListScraperTasks(ctx context.Context) ([]engine.ScraperTask, error) {
	tasks, err := r.queries.ListTasks(ctx)
	if err != nil {
		return nil, errors.New("unable to fetch the tasks from database")
	}
	var scraperTasks []engine.ScraperTask
	for _, task := range tasks {
		scraperTask, err := engine.UnmarshallTaskFromDatabase(task)
		if err != nil {
			return nil, errors.New("unable unmarshall task from database")
		}
		scraperTasks = append(scraperTasks, *scraperTask)
	}

	return scraperTasks, nil
}
