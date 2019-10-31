package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
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
	err := app.Initialize(ctx, "./logger.test.json", "./config.yaml")
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
	if _, err := db.Exec("ALTER SEQUENCE events_id_seq RESTART WITH 1"); err != nil {
		return err
	}
	db.Close()
	return nil
}

func validateEvent(t *testing.T, resultData map[string]interface{}, expectedData map[string]interface{}) {
	// Check if all user fields are present and correct
	for k, v := range expectedData {
		got, ok := resultData[k]
		if !ok || got != v {
			t.Errorf("Expected %v of type %T, got %v of type %T, at %v", v, v, got, got, resultData)
		}
	}

	// Check if automated fields are present and valid - UUID
	if uid, ok := resultData["UUID"]; ok {
		if _, err := uuid.Parse(uid.(string)); err != nil {
			t.Errorf("Expected valid UUID, got %v in %v with parsing error %v", uid, resultData, err)
		}
	} else {
		t.Errorf("Expected UUID present, got %v", resultData)
	}

	// Check if automated fields are present and valid - Datetime
	if dt, ok := resultData["Datetime"]; ok {
		if _, err := time.Parse("2006-01-02T15:04:05Z", dt.(string)); err != nil {
			t.Errorf("Expected valid Datetime, got %v with error %v", dt, err)
		}
	} else {
		t.Errorf("Expected Datetime present, got %v", resultData)
	}
}

func TestGetEvents(t *testing.T) {
	err := setup(true)
	if err != nil {
		t.Error("Setup failed:", err)
	}

	client := &http.Client{}
	url := apiURL + eventEndpoint

	// Get events
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("Failed to make request:", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Failed to do request:", err)
	}
	defer resp.Body.Close()

	// Check if status code is correct
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var results []map[string]interface{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		t.Error("Failed to unmarshal json:", err)
	}

	expected := []map[string]interface{}{
		map[string]interface{}{
			"Title":  "Foo",
			"Desc":   "FOOBAR",
			"Notify": true,
		},
		map[string]interface{}{
			"Title":  "Spam",
			"Desc":   "BAZINGA!",
			"Notify": false,
		},
		map[string]interface{}{
			"Title":  "Vookah",
			"Desc":   "You gonna like it",
			"User":   "Mack",
			"Notify": true,
		},
	}

	// Check if all field are correct
	for i, m := range expected {
		validateEvent(t, results[i], m)
	}
}

func TestPostEventAndGetResponse(t *testing.T) {
	err := setup(false)
	if err != nil {
		t.Error("Setup failed:", err)
	}

	client := &http.Client{}
	url := apiURL + eventEndpoint

	eventData := map[string]interface{}{
		"Title":    "MOOOOOOOOOOOOOOOOOOOOOOSGGGGGGGGGGGGGGGHHHHHHHHHHH",
		"User":     "Mack",
		"Desc":     "YOU GONNA LIKE THIS",
		"Notify":   true,
		"Datetime": "2019-09-26T15:15:00Z",
	}

	// Post new event to api
	var reqBody bytes.Buffer
	json.NewEncoder(&reqBody).Encode(eventData)
	req, err := http.NewRequest("POST", url, &reqBody)
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Failed to make request:", err)
	}
	defer resp.Body.Close()

	// Check if status code is correct
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Error("Failed to unmarshal json:", err)
	}

	validateEvent(t, result, eventData)
}

func TestPostEventAndSaveToDB(t *testing.T) {
	err := setup(false)
	if err != nil {
		t.Error("Setup failed:", err)
	}

	client := &http.Client{}
	url := apiURL + eventEndpoint

	eventData := map[string]interface{}{
		"Title":    "MOOOOOOOOOOOOOOOOOOOOOOSGGGGGGGGGGGGGGGHHHHHHHHHHH",
		"User":     "Mack",
		"Desc":     "YOU GONNA LIKE THIS",
		"Notify":   true,
		"Datetime": "2019-09-26T15:15:00Z",
	}

	// Post new event to api
	var reqBody bytes.Buffer
	json.NewEncoder(&reqBody).Encode(eventData)
	postReq, err := http.NewRequest("POST", url, &reqBody)
	_, err = client.Do(postReq)
	if err != nil {
		t.Error("Failed to execute request:", err)
	}

	// Make get request to recieve posted event from DB
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Failed to do request:", err)
	}
	defer resp.Body.Close()

	// Check if status code is correct
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	var results []map[string]interface{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		t.Error("Failed to unmarshal json:", err)
	}

	// Check if results have the correct len
	if l := len(results); l != 1 {
		t.Errorf("Expected to get exactly 1 result, got %d results: %v", l, results)
	}

	// Check if all user fields are present and correct
	result := results[0]
	validateEvent(t, result, eventData)
}

func TestDeleteEvent(t *testing.T) {

}

func TestUpdateEvent(t *testing.T) {

}
