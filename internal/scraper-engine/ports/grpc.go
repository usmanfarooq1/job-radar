package ports

import (
	"context"

	"github.com/usmanfarooq1/job-radar/internal/common/genproto/task"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) AddTask(ctx context.Context, request *task.CreateTaskRequest) (*task.Task, error) {
	return nil, nil
}
func (g GrpcServer) StopTask(ctx context.Context, request *task.TaskIdRequest) (*task.TaskStatusResponse, error) {
	return nil, nil
}
func (g GrpcServer) RunTask(ctx context.Context, request *task.TaskIdRequest) (*task.TaskStatusResponse, error) {
	return nil, nil
}
func (g GrpcServer) RemoveTask(ctx context.Context, request *task.TaskIdRequest) (*task.RemovedTaskResponse, error) {
	return nil, nil
}
func (g GrpcServer) UpdateTask(ctx context.Context, request *task.UpdateTaskRequest) (*task.Task, error) {
	return nil, nil
}
func (g GrpcServer) GetTask(ctx context.Context, request *task.TaskIdRequest) (*task.Task, error) {
	return nil, nil
}
func (g GrpcServer) ListTasks(ctx context.Context, request *task.EmptyRequest) (*task.ListTasksResponse, error) {
	return nil, nil
}
