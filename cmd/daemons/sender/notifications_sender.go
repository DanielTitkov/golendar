package main

import (
	"github.com/DanielTitkov/golendar/cmd/daemons/mq"
	"github.com/DanielTitkov/golendar/config"
	"github.com/DanielTitkov/golendar/logger"
	"go.uber.org/zap"
)

func notify(m mq.Message, l *zap.SugaredLogger) error {
	l.Warnf("Sending notification about upcoming event '%s' at %v to user '%s'", m.Title, m.Datetime, m.User)
	return nil
}

func main() {
	// setup logger
	l, err := logger.CreateLogger("./logger.json")
	if err != nil {
		panic(err)
	}

	// load config from file
	c, err := config.LoadYamlConfig("./config.yaml")
	if err != nil {
		l.Fatalf("config is not loaded: %v", err)
	}

	// setup rabbit
	conn, q, ch, err := mq.PrepareRabbitMQ(c)
	if err != nil {
		l.Fatal(err)
	}
	defer conn.Close()
	defer ch.Close()

	// register consumer
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		l.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			m, err := mq.Deserialize(d.Body)
			if err != nil {
				l.Errorf("Failed to deserialize, %v", err)
			}
			notify(m, l) // mock user notification
		}
	}()

	l.Info("Golendar Notification Sender up and running")
	l.Info("Waiting for messages")
	<-forever
}
