package engine

import "github.com/rs/zerolog"

type Engine struct {
	/*
		The Engine contains a manager object and all the behaviour for restarting the engine, stoping the engine,
		or halt the current processing of the engine.
	*/
	manager Manager
}

func MakeEngine(mq ScraperTaskPublishRepository, logger zerolog.Logger) Engine {
	engine := Engine{}
	engine.startEngine(mq)
	return engine
}

func (e *Engine) startEngine(mq ScraperTaskPublishRepository) {
	e.manager = MakeManager(mq, e.manager.logger)
}

func (e *Engine) Manager() Manager {
	return e.manager
}
