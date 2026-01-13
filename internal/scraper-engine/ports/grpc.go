package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/usmanfarooq1/job-radar/internal/common/genproto/task"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app/command"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app/query"
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
	id, err := uuid.Parse(request.TaskId)
	if err != nil {
		return nil, err
	}

	if err := g.app.Commands.StopScraperTask.Handle(ctx, command.StopTask{TaskId: id}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
func (g GrpcServer) RunTask(ctx context.Context, request *task.TaskIdRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(request.TaskId)
	if err != nil {
		return nil, err
	}

	if err := g.app.Commands.RunScraperTask.Handle(ctx, command.RunTask{TaskId: id}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
func (g GrpcServer) RemoveTask(ctx context.Context, request *task.TaskIdRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(request.TaskId)
	if err != nil {
		return nil, err
	}

	if err := g.app.Commands.RemoveScraperTask.Handle(ctx, command.RemoveTask{TaskId: id}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
func (g GrpcServer) UpdateTask(ctx context.Context, request *task.UpdateTaskRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(request.TaskId.GetTaskId())
	if err != nil {
		return nil, err
	}

	if err := g.app.Commands.UpdateScraperTask.Handle(ctx, command.UpdateTask{
		TaskId:         id,
		DelayInSeconds: request.DelayInSeconds,
		SearchKeyword:  request.SearchKeyword,
		LocationId:     request.LocationId,
		DistanceRadius: request.DistanceRadius,
		TaskLocation:   request.TaskLocation,
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
func (g GrpcServer) GetTask(ctx context.Context, request *task.TaskIdRequest) (*task.Task, error) {
	id, err := uuid.Parse(request.TaskId)
	if err != nil {
		return nil, err
	}

	query, err := g.app.Queries.GetTask.Handle(ctx, query.TaskQuery{
		TaskId: id,
	})

	if err != nil {
		return nil, err
	}

	return &task.Task{
		TaskId:         query.TaskId,
		TaskType:       query.TaskType,
		DelayInSeconds: query.DelayInSeconds,
		SearchKeyword:  query.SearchKeyword,
		LocationId:     query.LocationId,
		DistanceRadius: query.DistanceRadius,
		TaskLocation:   query.TaskLocation,
		CreatedAt:      query.CreatedAt,
		UpdatedAt:      query.UpdatedAt,
	}, nil
}
func (g GrpcServer) ListTasks(ctx context.Context, request *task.EmptyRequest) (*task.ListTasksResponse, error) {
	query, err := g.app.Queries.ListTasks.Handle(ctx, query.ListTasksQuery{})
	if err != nil {
		return nil, err
	}
	tasksList := make([]*task.Task, len(query))
	for _, ele := range query {
		tasksList = append(tasksList, &task.Task{
			TaskId:         ele.TaskId,
			TaskType:       ele.TaskType,
			DelayInSeconds: ele.DelayInSeconds,
			SearchKeyword:  ele.SearchKeyword,
			LocationId:     ele.LocationId,
			DistanceRadius: ele.DistanceRadius,
			TaskLocation:   ele.TaskLocation,
			CreatedAt:      ele.CreatedAt,
			UpdatedAt:      ele.UpdatedAt,
		})
	}

	return &task.ListTasksResponse{Tasks: tasksList}, nil
}
