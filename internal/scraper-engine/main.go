package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	pb "github.com/usmanfarooq1/job-radar/internal/common/genproto/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/adapters"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/ports"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/service"
)

func main() {
	ctx := context.Background()
	// Logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// // RabbitMQ
	// mqConn, err := amqp.Dial()
	// if err != nil {
	// 	logger.Err(err).Msg("Unable to connect to rabbitmq")
	// }
	// channel, err := mqConn.Channel()
	// if err != nil {
	// 	logger.Err(err).Msg("Unable to create a channel to rabbitmq")
	// }
	mq, err := adapters.NewMQPublisher(os.Getenv("MQ_URL"), os.Getenv("MQ_JOB_LINK_QUEUE_NAME"), logger)
	if err != nil {
		logger.Err(err).Msg("Unable to connect to RabbitMQ")
		os.Exit(1)
	}
	// Database connection
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Err(err).Msg("Unable to connect to database")
		os.Exit(1)
	}
	// Engine
	engine := engine.MakeEngine(mq, logger)

	defer func() {
		conn.Close(ctx)
		// mqConn.Close()
		// channel.Close()
	}()
	application := service.NewApplication(ctx, logger, conn, engine)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s",
		os.Getenv("SERVICE_HOSTNAME"),
		os.Getenv("SERVICE_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	pb.RegisterScraperTaskRouteServer(grpcServer, ports.NewGrpcServer(application))
	logger.Info().Msg(fmt.Sprintf("Starting engine service on %s\n", os.Getenv("SERVICE_PORT")))
	grpcServer.Serve(lis)
}
