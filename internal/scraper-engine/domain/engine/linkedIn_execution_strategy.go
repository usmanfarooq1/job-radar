package engine

import (
	"fmt"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/usmanfarooq1/job-radar/internal/common/mq"
)

type LinkedInExecutionStrategy struct {
	query string
}

func (ls LinkedInExecutionStrategy) JobExtractor(task *ScraperTask) <-chan mq.JobLinkMessagePayload {
	ticker := time.NewTicker(time.Duration(task.delayInSeconds) * time.Second)
	select {
	case <-task.executionChannel:
		ticker.Stop()
		return nil
	case t := <-ticker.C:
		fmt.Printf("Executing the job search on: %s, at %s\n", ls.query, t)
		ls.scrapeJobs(*task.pBrowser)
		return task.resultChannel
	}
}

func (ls LinkedInExecutionStrategy) scrapeJobs(browser playwright.Browser) {

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

	for i, entry := range entries {
		text, err := entry.Locator("a").First().GetAttribute("href")
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%d: %s\n", i+1, text)

	}
	if err := browser.Close(); err != nil {
		log.Fatalf("can't stop browser : %v", err)
	}

}
