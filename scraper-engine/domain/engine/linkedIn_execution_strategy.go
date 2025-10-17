package engine

import (
	"fmt"
	"time"
)

type LinkedInExecutionStrategy struct {
	query string
}

func (ls LinkedInExecutionStrategy) JobExtractor(task *ScraperTask) {
	ticker := time.NewTicker(time.Duration(task.delayInSeconds) * time.Second)
	select {
	case <-task.executionChannel:
		ticker.Stop()
		return
	case t := <-ticker.C:
		fmt.Printf("Executing the job search on: %s, at %s\n", ls.query, t)
	}
}
