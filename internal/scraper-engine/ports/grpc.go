package ports

import (
	"context"

	"github.com/usmanfarooq1/job-radar/internal/common/genproto/task"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app/command"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) AddTask(ctx context.Context, request *task.CreateTaskRequest) (*emptypb.Empty, error) {
	if err := g.app.Commands.AddScraperTask.Handle(ctx, command.AddTask{
		DelayInSeconds: request.DelayInSeconds,
		SearchKeyword:  request.SearchKeyword,
		LocationId:     request.LocationId,
		TaskType:       request.TaskType,
		DistanceRadius: request.DistanceRadius,
		TaskLocation:   request.TaskLocation,
	}); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (g GrpcServer) StopTask(ctx context.Context, request *task.TaskIdRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (g GrpcServer) RunTask(ctx context.Context, request *task.TaskIdRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (g GrpcServer) RemoveTask(ctx context.Context, request *task.TaskIdRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (g GrpcServer) UpdateTask(ctx context.Context, request *task.UpdateTaskRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (g GrpcServer) GetTask(ctx context.Context, request *task.TaskIdRequest) (*task.Task, error) {
	return nil, nil
}
func (g GrpcServer) ListTasks(ctx context.Context, request *task.EmptyRequest) (*task.ListTasksResponse, error) {
	return nil, nil
}
