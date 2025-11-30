package engine

import (
	"fmt"
	"time"

	"github.com/usmanfarooq1/job-radar/internal/common/mq"
)

type LinkedInExecutionStrategy struct {
	query string
}

func (ls LinkedInExecutionStrategy) JobExtractor(task *ScraperTask) <-chan mq.JobLinkMessagePayload {
	ticker := time.NewTicker(time.Duration(task.delayInSeconds) * time.Second)

	select {
	case <-task.executionChannel:
		ticker.Stop()
		return nil
	case t := <-ticker.C:
		fmt.Printf("Executing the job search on: %s, at %s\n", ls.query, t)
		return task.resultChannel
	}
}
