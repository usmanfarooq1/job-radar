package adapters

import (
	"fmt"
	"strings"

	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type LinkedInScraperTaskQueryBuilderStrategy struct {
}

func (l LinkedInScraperTaskQueryBuilderStrategy) Construct(task *engine.ScraperTask) (string, error) {

	query := fmt.Sprintf("https://www.linkedin.com/jobs/search?keywords=%s&location=%s&geoId=%s&distance=%d&f_TPR=r%d",
		strings.ReplaceAll(task.SearchKeyword(), " ", "%20"),
		task.SearchLocation(),
		task.LocationId(),
		task.DistanceRadius(),
		task.DelayInSeconds())

	return query, nil
}
