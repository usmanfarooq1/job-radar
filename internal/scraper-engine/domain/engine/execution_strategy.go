package engine

type ExecutionStrategy interface {
	JobExtractor(t *ScraperTask)
}

func GenerateExecutionStrategy(task *ScraperTask) (ExecutionStrategy, error) {
	queryBuilder, err := GenerateQueryBuilderStrategy(task.taskType)
	if err != nil {
		return nil, err
	}
	query, err := queryBuilder.Construct(task)
	if err != nil {
		return nil, err
	}

	switch task.taskType {
	case LinkedIn:
		return LinkedInExecutionStrategy{query: query}, nil
	}
	return nil, ErrInvalidTaskType
}
