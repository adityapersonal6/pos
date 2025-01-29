package events

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type Producer struct {
	Channel *amqp.Channel
	Queue   string
}

func NewProducer(channel *amqp.Channel, queue string) *Producer {
	return &Producer{
		Channel: channel,
		Queue:   queue,
	}
}

func (p *Producer) Publish(message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.Channel.Publish(
		"",      // exchange
		p.Queue, // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %s", err)
		return err
	}

	return nil
}
