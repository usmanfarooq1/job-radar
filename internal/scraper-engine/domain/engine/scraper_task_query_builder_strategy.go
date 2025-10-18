package engine

type ScraperTaskQueryBuilderStrategy interface {
	Construct(t *ScraperTask) (string, error)
}

func GenerateQueryBuilderStrategy(taskType ScraperTaskType) (ScraperTaskQueryBuilderStrategy, error) {
	switch taskType {
	case LinkedIn:
		return LinkedInScraperTaskQueryBuilderStrategy{}, nil
	}
	return LinkedInScraperTaskQueryBuilderStrategy{}, ErrInvalidTaskType
}
