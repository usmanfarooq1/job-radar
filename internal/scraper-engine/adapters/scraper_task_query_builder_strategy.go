package adapters

import "github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"

type ScraperTaskQueryBuilderStrategy interface {
	Construct(t *engine.ScraperTask) (string, error)
}

func GenerateQueryBuilderStrategy(taskType engine.ScraperTaskType) (ScraperTaskQueryBuilderStrategy, error) {
	switch taskType {
	case engine.LinkedIn:
		return LinkedInScraperTaskQueryBuilderStrategy{}, nil
	}
	return LinkedInScraperTaskQueryBuilderStrategy{}, engine.ErrInvalidTaskType
}
