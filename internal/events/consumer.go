package events

import (
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewConsumer(amqpURL, queue string) (*Consumer, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:    conn,
		channel: channel,
		queue:   queue,
	}, nil
}

func (c *Consumer) StartConsuming(handler func([]byte)) error {
	msgs, err := c.channel.Consume(
		c.queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		handler(msg.Body)
	}

	return nil
}

func (c *Consumer) Close() {
	if err := c.channel.Close(); err != nil {
		log.Printf("Error closing channel: %s", err)
	}
	if err := c.conn.Close(); err != nil {
		log.Printf("Error closing connection: %s", err)
	}
}
