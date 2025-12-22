package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/usmanfarooq1/job-radar/internal/common/genproto/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/ports"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/service"
)

func main() {
	ctx := context.Background()
	application := service.NewApplication(ctx)
	lis, err := net.Listen("tcp", fmt.Sprintf("engine:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	pb.RegisterScraperTaskRouteServer(grpcServer, ports.NewGrpcServer(application))
	fmt.Println("Starting engine service on 50051")
	grpcServer.Serve(lis)
}
