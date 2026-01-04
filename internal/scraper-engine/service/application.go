package service

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/adapters"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app/command"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app/query"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

func NewApplication(ctx context.Context) app.Application {
	// Logger
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// RabbitMQ
	mqConn, err := amqp.Dial(os.Getenv("MQ_URL"))
	if err != nil {
		logger.Err(err).Msg("Unable to connect to rabbitmq")
	}
	channel, err := mqConn.Channel()
	if err != nil {
		logger.Err(err).Msg("Unable to create a channel to rabbitmq")
	}
	mq := adapters.NewMQPublisher(mqConn, channel, os.Getenv("MQ_JOB_LINK_QUEUE_NAME"), logger)

	// Database connection
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Err(err).Msg("Unable to connect to database")
		os.Exit(1)
	}
	engine := engine.Engine{}
	engine.StartEngine(&mq)

	defer func() {
		conn.Close(ctx)
		mqConn.Close()
		channel.Close()
	}()

	taskRepository := adapters.NewSQLScraperTaskRepository(conn, logger)
	logger.Info().Msg("A new application instance is established")
	return app.Application{
		Commands: app.Commands{
			AddScraperTask:    command.NewAddTaskHandler(engine, taskRepository),
			StopScraperTask:   command.NewStopTaskHandler(engine, taskRepository),
			RunScraperTask:    command.NewRunTaskHandler(engine, taskRepository),
			RemoveScraperTask: command.NewRemoveTaskHandler(engine, taskRepository),
			UpdateScraperTask: command.NewUpdateTaskHandler(engine, taskRepository),
		},
		Queries: app.Queries{
			GetTask:   query.NewGetTaskkHandler(engine, taskRepository),
			ListTasks: query.NewListTasksHandler(taskRepository),
		},
	}
}
