package msgbus

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const (
	QUEUE_BACKBURNER = "nienna_backburner"
)

type Msgbus struct {
	uri    string
	queues []string
	conn   *amqp.Connection
	ch     *amqp.Channel
}

func NewMsgbus(uri string, queues ...string) (*Msgbus, error) {
	msgBus := Msgbus{uri, queues, nil, nil}
	if err := msgBus.connect(); err != nil {
		return nil, err
	}
	if err := msgBus.createQueues(); err != nil {
		return nil, err
	}
	return &msgBus, nil
}

func (m *Msgbus) Publish(queue string, event *EventSerialization) error {
	if err := m.tryReconnect(); err != nil {
		return err
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return m.ch.Publish("", queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        payload,
	})
}

func (m *Msgbus) connect() error {
	var err error
	if m.conn, err = amqp.Dial(m.uri); err != nil {
		return err
	}
	if m.ch, err = m.conn.Channel(); err != nil {
		return err
	}
	return err
}

func (m *Msgbus) createQueues() error {
	for _, queue := range m.queues {
		if _, err := m.ch.QueueDeclare(queue, false, false, false, false, nil); err != nil {
			return err
		}
	}
	return nil
}

func (m *Msgbus) tryReconnect() error {
	var try uint
	for m.conn.IsClosed() && try < 10 {
		log.Debug("Msgbug - is closed, attempting to reconnect")
		if err := m.connect(); err != nil {
			log.Error("Reconnection error ", err)
		}
		if err := m.createQueues(); err != nil {
			log.Error("Reconnection error ", err)
		}
		time.Sleep(5 * time.Second)
		try++
	}
	if m.conn.IsClosed() && try >= 10 {
		return fmt.Errorf("fail to reconnect to amqp server")
	}
	return nil
}
