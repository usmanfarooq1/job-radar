package engine

type JobQueryBuilderStrategy interface {
	Construct(t *ScraperTask) (string, error)
}

func GenerateQueryBuilderStrategy(taskType ScraperTaskType) (JobQueryBuilderStrategy, error) {
	switch taskType {
	case LinkedIn:
		return LinkedInJobQueryBuilderStrategy{}, nil
	}
	return LinkedInJobQueryBuilderStrategy{}, ErrInvalidTaskType
}
