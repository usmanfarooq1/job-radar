package engine

import (
	"context"

	"github.com/usmanfarooq1/job-radar/internal/common/mq"
)

type ScraperTaskPublishRepository interface {
	Publish(ctx context.Context, message mq.JobLinkMessagePayload) error
}
