package engine

import "fmt"

type LinkedInJobQueryBuilderStrategy struct {
}

func (l LinkedInJobQueryBuilderStrategy) Construct(task *ScraperTask) (string, error) {

	query := fmt.Sprintf("https://www.linkedin.com/jobs/search?keywords=%s&geoId=%s&distance=%d&f_TPR=r%d",
		task.searchKeyword,
		task.taskLocationId,
		task.distanceRadius,
		task.delayInSeconds)

	return query, nil
}
