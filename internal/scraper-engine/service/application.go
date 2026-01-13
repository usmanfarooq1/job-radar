package service

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/adapters"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app/command"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app/query"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

func NewApplication(ctx context.Context, logger zerolog.Logger, db *pgx.Conn, engine engine.Engine) app.Application {

	taskRepository := adapters.NewSQLScraperTaskRepository(db, logger)
	logger.Info().Msg("A new application instance is established")
	return app.Application{
		Commands: app.Commands{
			AddScraperTask:    command.NewAddTaskHandler(engine, logger, taskRepository),
			StopScraperTask:   command.NewStopTaskHandler(engine, logger, taskRepository),
			RunScraperTask:    command.NewRunTaskHandler(engine, logger, taskRepository),
			RemoveScraperTask: command.NewRemoveTaskHandler(engine, logger, taskRepository),
			UpdateScraperTask: command.NewUpdateTaskHandler(engine, logger, taskRepository),
		},
		Queries: app.Queries{
			GetTask:   query.NewGetTaskkHandler(engine, logger, taskRepository),
			ListTasks: query.NewListTasksHandler(taskRepository, logger),
		},
	}
}
