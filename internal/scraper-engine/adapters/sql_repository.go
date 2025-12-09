package adapters

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/usmanfarooq1/job-radar/internal/common/db"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type SQLScraperTaskRepository struct {
	db      *pgx.Conn
	queries *db.Queries
	logger  zerolog.Logger
}

func NewSQLScraperTaskRepository(dbConn *pgx.Conn, logger zerolog.Logger) SQLScraperTaskRepository {
	if dbConn == nil {
		error := errors.New("null database connnection passed")
		logger.Error().Stack().Err(error).Msg("unable to connect to database")
	}
	queries := db.New(dbConn)
	return SQLScraperTaskRepository{db: dbConn, queries: queries, logger: logger}
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
		r.logger.Error().Err(err).Stack().Dict("task", zerolog.Dict().
			Str("location", st.SearchLocation()).
			Str("keyword", st.SearchKeyword()).Str("locationId", st.LocationId())).
			Msg("unable to create a scraper task")
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
		r.logger.Error().Err(err).Stack().Dict("task", zerolog.Dict().
			Str("scraper_task_id", st.Id().String()).
			Str("location", st.SearchLocation()).
			Str("keyword", st.SearchKeyword()).
			Str("locationId", st.LocationId())).
			Msg("unable to update scraper task")
		return nil, errors.New("unable to update scraper task")
	}
	return st, nil
}
func (r SQLScraperTaskRepository) RemoveScraperTask(ctx context.Context, id uuid.UUID) error {
	err := r.queries.DeleteTask(ctx, id)
	if err != nil {
		r.logger.Error().Err(err).Stack().Dict("task", zerolog.Dict().
			Str("scraper_task_id", id.String())).
			Msg("unable to delete scraper task")
		return errors.New("unable to delete scraper task")
	}
	return nil
}
func (r SQLScraperTaskRepository) GetScraperTask(ctx context.Context, id uuid.UUID) (*engine.ScraperTask, error) {
	task, err := r.queries.GetTask(ctx, id)
	if err != nil {
		r.logger.Error().Err(err).Stack().Dict("task", zerolog.Dict().
			Str("scraper_task_id", id.String())).
			Msg("unable to fetch a task from database")
		return nil, errors.New("unable to fetch a task from database")
	}

	scraperTask, err := engine.UnmarshallTaskFromDatabase(task)
	if err != nil {
		r.logger.Error().Err(err).Stack().Dict("task", zerolog.Dict().
			Str("scraper_task_id", id.String())).
			Msg("unable to unmarshall task from database")
		return nil, errors.New("unable to unmarshall task from database")
	}
	return scraperTask, nil
}
func (r SQLScraperTaskRepository) ListScraperTasks(ctx context.Context) ([]engine.ScraperTask, error) {
	tasks, err := r.queries.ListTasks(ctx)
	if err != nil {
		r.logger.Error().Err(err).Stack().
			Msg("unable to fetch the tasks from database")
		return nil, errors.New("unable to fetch the tasks from database")
	}
	var scraperTasks []engine.ScraperTask
	for _, task := range tasks {
		scraperTask, err := engine.UnmarshallTaskFromDatabase(task)
		if err != nil {
			r.logger.Error().Err(err).Stack().Dict("task", zerolog.Dict().
				Str("scraper_task_id", task.TaskID.String())).
				Msg("unable to unmarshall task from database")
			return nil, errors.New("unable to unmarshall task from database")
		}
		scraperTasks = append(scraperTasks, *scraperTask)
	}

	return scraperTasks, nil
}
