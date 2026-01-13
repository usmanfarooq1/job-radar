package adapters

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rs/zerolog"
	"github.com/usmanfarooq1/job-radar/internal/common/mq"
)

type MQPublisher struct {
	logger           zerolog.Logger
	connectionString string
	queueName        string
}

func NewMQPublisher(connString, queueName string, logger zerolog.Logger) (*MQPublisher, error) {

	// queue, err := ch.QueueDeclare(queueName, true, true, false, false, nil)
	// if err != nil {
	// 	return nil, errors.New("unable to marshall message to json")
	// }
	return &MQPublisher{connectionString: connString, queueName: queueName, logger: logger}, nil
}

func (p *MQPublisher) Publish(ctx context.Context, message mq.JobLinkMessagePayload) error {
	_, err := json.Marshal(message)
	if err != nil {
		p.logger.Error().Stack().Err(err).Dict("message", zerolog.Dict().
			Str("location", message.Location).
			Str("location_id", message.LocationId).
			Str("job_link", message.Location)).Msg("unable to marshall message to json")
		return errors.New("unable to marshall message to json")
	}
	// p.ch.PublishWithContext(ctx,
	// 	"",          // exchange
	// 	p.queueName, // routing key
	// 	false,       // mandatory
	// 	false,       // immediate
	// 	amqp.Publishing{
	// 		ContentType: "application/json",
	// 		Body:        []byte(body),
	// 	})
	return nil
}
