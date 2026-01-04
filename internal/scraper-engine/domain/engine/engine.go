package engine

type Engine struct {
	/*
		The Engine contains a manager object and all the behaviour for restarting the engine, stoping the engine,
		or halt the current processing of the engine.
	*/
	manager Manager
}

func (e *Engine) StartEngine(mq ScraperTaskPublishRepository) {
	e.manager = MakeManager(mq)
}

func (e *Engine) Manager() Manager {
	return e.manager
}
