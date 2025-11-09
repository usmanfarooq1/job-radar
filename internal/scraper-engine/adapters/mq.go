package adapters

import (
	"context"
	"encoding/json"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type MQPublisher struct {
	conn      *amqp.Connection
	ch        *amqp.Channel
	logger    zerolog.Logger
	queueName string
}

type Message struct {
	Location   string `json:"location"`
	LocationId string `json:"location_id"`
	JobLink    string `json:"job_link"`
}

func CreateMessage(location string, locationId string, jobLink string) Message {
	return Message{Location: location, LocationId: locationId, JobLink: jobLink}
}

func NewMQPublisher(conn *amqp.Connection, ch *amqp.Channel, queue string) MQPublisher {

	return MQPublisher{conn: conn, ch: ch, queueName: queue}
}

func (p *MQPublisher) Publish(ctx context.Context, message Message) error {
	body, err := json.Marshal(message)
	if err != nil {
		p.logger.Error().Stack().Err(err).Dict("message", zerolog.Dict().
			Str("location", message.Location).
			Str("location_id", message.LocationId).
			Str("job_link", message.Location)).Msg("unable to marshall message to json")
		return errors.New("unable to marshall message to json")
	}
	p.ch.PublishWithContext(ctx,
		"",          // exchange
		p.queueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	return nil
}
