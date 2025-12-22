package service

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/adapters"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app/command"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

func NewApplication(ctx context.Context) app.Application {
	engine := engine.Engine{}
	engine.StartEngine()

	// Logger
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Database connection
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Err(err).Msg("Unable to connect to database")
		os.Exit(1)
	}

	// RabbitMQ

	defer conn.Close(ctx)

	taskRepository := adapters.NewSQLScraperTaskRepository(conn, logger)
	logger.Info().Msg("A new application instance is established")
	return app.Application{
		Commands: app.Commands{
			AddScraperTask: command.NewAddTaskHandler(engine, taskRepository),
		},
	}
}
