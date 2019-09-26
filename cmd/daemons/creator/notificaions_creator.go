package main

import (
	"context"
	"time"

	"github.com/DanielTitkov/golendar/cmd/daemons/mq"
	"github.com/DanielTitkov/golendar/config"
	"github.com/DanielTitkov/golendar/internal/event"
	"github.com/DanielTitkov/golendar/internal/storage"
	"github.com/DanielTitkov/golendar/logger"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func publishMessages(
	events []event.Event,
	l *zap.SugaredLogger,
	s storage.Storage,
	q *amqp.Queue,
	ch *amqp.Channel,
) {
	for _, e := range events {
		msg := mq.Message{
			User:     e.User,
			Title:    e.Title,
			Datetime: e.Datetime,
		}

		JSONMsg, err := mq.Serialize(msg)
		if err != nil {
			l.Error(err, msg)
		}

		err = ch.Publish(
			"", // exchange
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(JSONMsg),
			},
		)
		l.Infof("Sent %s", msg)
		if err != nil {
			l.Error(err)
		} else {
			e.Notify = false
			_, err := s.UpdateEvent(e.UUID, e)
			if err != nil {
				l.Error(err)
			}
		}
	}
}

func main() {
	ctx := context.Context(context.Background())

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

	// setup storage
	l.Infof("Using '%s' storage", c.Storage)
	s, err := storage.PrepareStorage(ctx, c)
	if err != nil {
		l.Fatal(err)
	}

	// connect to mq
	conn, q, ch, err := mq.PrepareRabbitMQ(c)
	if err != nil {
		l.Fatal(err)
	}
	defer conn.Close()
	defer ch.Close()

	// get events

	for range time.NewTicker(5 * time.Second).C {
		events, err := s.GetUpcomingEvents(c.Notify)
		l.Infof("Fetching events, got %d", len(events))
		if err != nil {
			l.Errorf("Events fetching error, %v", err)
		}
		publishMessages(events, l, s, q, ch)
	}
}
