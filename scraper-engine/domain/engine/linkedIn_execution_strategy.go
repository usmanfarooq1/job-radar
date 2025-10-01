package engine

type LinkedInExecutionStrategy struct {
	query string
}

func (ls LinkedInExecutionStrategy) JobExtractionInterface(t *Task) error {
	return nil
}
