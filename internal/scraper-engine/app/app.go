package app

import (
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app/command"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}
type Commands struct {
	AddScraperTask  command.AddTaskHandler
	StopScraperTask command.StopTaskHandler
	RunScraperTask  command.RunTaskHandler
}

type Queries struct {
	GetTask   query.GetTaskHandler
	ListTasks query.ListTasksHandler
}
