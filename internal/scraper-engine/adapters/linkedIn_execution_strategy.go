package adapters

import (
	"fmt"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/usmanfarooq1/job-radar/internal/common/mq"
	"github.com/usmanfarooq1/job-radar/internal/scraper-engine/domain/engine"
)

type LinkedInExecutionStrategy struct {
	query    string
	pBrowser playwright.Browser
}

func (ls LinkedInExecutionStrategy) JobExtractor(task *engine.ScraperTask) {
	ticker := time.NewTicker(time.Duration(task.DelayInSeconds()) * time.Second)
	for {
		select {
		case <-task.ExecutionChannel():
			ticker.Stop()
		case t := <-ticker.C:
			fmt.Printf("Executing the job search on: %s, at %s\n", ls.query, t)
			ls.scrapeJobs(ls.pBrowser, *task)
		}
	}
}

func (ls LinkedInExecutionStrategy) scrapeJobs(browser playwright.Browser, task engine.ScraperTask) {

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("can't create a page : %v", err)
	}
	if _, err := page.Goto(ls.query); err != nil {
		log.Fatalf("can't create a page : %v", err)
	}
	entries, err := page.Locator(".jobs-search__results-list > li").All()
	if err != nil {
		log.Fatalf("can't find the job search : %v", err)
	}

	for _, entry := range entries {
		jobLink, err := entry.Locator("a").First().GetAttribute("href")
		if err != nil {
			log.Println(err)
		}
		task.ResultChannel() <- mq.CreateJobLinkMessagePayload(task.TaskLocation(), task.LocationId(), jobLink)
	}
}

func GenerateExecutionStrategy(task *engine.ScraperTask) (engine.ExecutionStrategy, error) {
	queryBuilder, err := GenerateQueryBuilderStrategy(task.TaskType())
	if err != nil {
		return nil, err
	}
	query, err := queryBuilder.Construct(task)
	if err != nil {
		return nil, err
	}

	switch task.TaskType() {
	case engine.LinkedIn:
		return LinkedInExecutionStrategy{query: query}, nil
	}
	return nil, engine.ErrInvalidTaskType
}
