package main

import (
	"context"
	"time"

	"github.com/DanielTitkov/golendar/config"
	"github.com/DanielTitkov/golendar/internal/event"
	"github.com/DanielTitkov/golendar/internal/storage"
	"github.com/DanielTitkov/golendar/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// App keeps app settings
type App struct {
	l *zap.SugaredLogger
	c config.Config
	s storage.Storage
}

// Initialize setups app
func (app *App) Initialize(
	ctx context.Context,
	loggerFile string,
	configFile string,
) error {
	// setup logger
	l, err := logger.CreateLogger(loggerFile)
	if err != nil {
		return err
	}
	app.l = l

	// load config from file
	c, err := config.LoadYamlConfig(configFile)
	if err != nil {
		return err
	}
	app.c = c

	// setup storage
	l.Infof("Using '%s' storage", c.Storage)
	s, err := storage.PrepareStorage(ctx, c)
	if err != nil {
		return err
	}
	app.s = s

	return nil
}

// MockEvents creates some events in database
func (app *App) MockEvents() error {
	// create db table if databese is used as storage
	if app.c.Storage == "Postgres" {
		err := app.prepareDB()
		if err != nil {
			return err
		}
	}

	events := []event.Event{
		{Title: "Foo", Desc: "FOOBAR", Notify: true, Datetime: time.Now()},
		{Title: "Spam", Desc: "BAZINGA!", Notify: false, Datetime: time.Now()},
		{Title: "Vookah", User: "Mack", Desc: "You gonna like it", Notify: true, Datetime: time.Now()},
	}
	for _, e := range events {
		_, err := app.s.CreateEvent(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (app *App) prepareDB() error {
	db, err := sqlx.Open("pgx", app.c.DBURI)
	if err != nil {
		return err
	}
	initTableQuery := `
		create table if not exists events (
			id serial primary key,
			UUID text not null,
			title text,
			datetime timestamp,
			duration text,
			description text,
			userid text,
			notify boolean);
	`
	if _, err := db.Exec(initTableQuery); err != nil {
		return err
	}
	db.Close()
	return nil
}
