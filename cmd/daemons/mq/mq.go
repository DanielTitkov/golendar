package mq

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/DanielTitkov/golendar/config"
	"github.com/streadway/amqp"
)

// Message with info about event
type Message struct {
	User     string
	Title    string
	Datetime time.Time
}

// Serialize message to json
func Serialize(msg Message) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}

// Deserialize json to message
func Deserialize(b []byte) (Message, error) {
	var msg Message
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}

// PrepareRabbitMQ return queue and channel
func PrepareRabbitMQ(
	c config.Config,
) (*amqp.Connection, *amqp.Queue, *amqp.Channel, error) {
	// connect to rabbit
	conn, err := amqp.Dial(c.MQURI)
	if err != nil {
		return nil, nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, nil, err
	}

	q, err := ch.QueueDeclare(
		"notifications", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusion
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		return nil, nil, nil, err
	}
	return conn, &q, ch, nil
}
