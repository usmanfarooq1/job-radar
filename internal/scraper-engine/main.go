package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/usmanfarooq1/job-radar/internal/common/genproto/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/ports"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/service"
)

func main() {
	ctx := context.Background()
	application := service.NewApplication(ctx)
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
	fmt.Printf("Starting engine service on %s\n", os.Getenv("SERVICE_PORT"))
	grpcServer.Serve(lis)
}
