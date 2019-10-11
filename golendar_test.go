package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jmoiron/sqlx"
)

const apiURL string = "http://localhost:3000"
const eventEndpoint string = "/events"

func setup(addEvents bool) error {
	app, err := makeApp()
	if err != nil {
		return err
	}
	err = app.prepareDB()
	if err != nil {
		return err
	}
	err = clearTable(app)
	if err != nil {
		return err
	}
	if addEvents {
		err = app.MockEvents()
		if err != nil {
			return err
		}
	}
	return nil
}

func makeApp() (*App, error) {
	ctx := context.Context(context.Background())
	app := App{}
	err := app.Initialize(ctx, "./logger.json", "./config.yaml")
	if err != nil {
		return &App{}, err
	}
	return &app, nil
}

func clearTable(app *App) error {
	db, err := sqlx.Open("pgx", app.c.DBURI)
	if err != nil {
		return err
	}
	if _, err := db.Exec("DELETE FROM events"); err != nil {
		return err
	}
	if _, err := db.Exec("ALTER TABLE events AUTO_INCREMENT = 1"); err != nil {
		return err
	}
	db.Close()
	return nil
}

func TestGetEvents(t *testing.T) {
	err := setup(true)
	if err != nil {
		t.Error("Setup failed", err)
	}

	client := &http.Client{}
	url := apiURL + eventEndpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("Failed to make request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Failed to do request", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
