package msgbus

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

const (
	QUEUE_BACKBURNER = "nienna_backburner"
)

type Msgbus struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewMsgbus(uri string, queues ...string) (*Msgbus, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	for _, queue := range queues {
		_, err := ch.QueueDeclare(queue, false, false, false, false, nil)
		if err != nil {
			return nil, err
		}
	}

	return &Msgbus{conn, ch}, nil
}

func (m *Msgbus) Publish(queue string, event *EventSerialization) error {
	// FIXME handle when rabbit is down

	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return m.ch.Publish("", queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        payload,
	})
}
