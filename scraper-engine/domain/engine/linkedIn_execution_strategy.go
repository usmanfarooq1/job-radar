package engine

type LinkedInExecutionStrategy struct {
	query string
}

func (ls LinkedInExecutionStrategy) JobExtractionInterface(t *ScraperTask) error {
	return nil
}
