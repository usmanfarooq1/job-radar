package engine

import (
	"fmt"
	"strings"
)

type LinkedInScraperTaskQueryBuilderStrategy struct {
}

func (l LinkedInScraperTaskQueryBuilderStrategy) Construct(task *ScraperTask) (string, error) {

	query := fmt.Sprintf("https://www.linkedin.com/jobs/search?keywords=%s&location=%s&geoId=%s&distance=%d&f_TPR=r%d",
		strings.ReplaceAll(task.searchKeyword, " ", "%20"),
		task.taskLocation,
		task.taskLocationId,
		task.distanceRadius,
		task.delayInSeconds)

	return query, nil
}
