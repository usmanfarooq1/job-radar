package engine

type JobQueryBuilderStrategy interface {
	Construct(t *Task) (string, error)
}

func GenerateQueryBuilderStrategy(taskType TaskType) (JobQueryBuilderStrategy, error) {
	switch taskType {
	case LinkedIn:
		return LinkedInJobQueryBuilderStrategy{}, nil
	}
	return LinkedInJobQueryBuilderStrategy{}, ErrInvalidTaskType
}
