package pgstorage

import (
	"context"

	"github.com/DanielTitkov/golendar/internal/event"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/stdlib" // for postgres
	"github.com/jmoiron/sqlx"
)

// PGStorage is struct for postgres storage
type PGStorage struct {
	URI string
	DB  *sqlx.DB
	Ctx context.Context
}

// Init setups map in MapStorage
func (pgs *PGStorage) Init() error {
	db, err := sqlx.Open("pgx", pgs.URI)
	if err != nil {
		return err
	}
	pgs.DB = db
	return nil
}

// GetEvents gets all stored events
func (pgs *PGStorage) GetEvents() ([]event.Event, error) {
	sql := "select uuid, title, datetime, duration, description, userid, notify from events"
	rows, err := pgs.DB.QueryxContext(pgs.Ctx, sql)
	if err != nil {
		return []event.Event{}, err
	}
	defer rows.Close()

	events := []event.Event{}

	for rows.Next() {
		var e event.Event
		err := rows.StructScan(&e)
		if err != nil {
			return []event.Event{}, err
		}
		events = append(events, e)
	}
	if err := rows.Err(); err != nil {
		return []event.Event{}, err
	}
	return events, nil
}

// GetEvent gets event by UUID
func (pgs *PGStorage) GetEvent(eventUUID uuid.UUID) (event.Event, error) {
	sql := "select uuid, title, datetime, duration, description, userid, notify from events where uuid = $1"
	res := make(map[string]interface{})
	err := pgs.DB.QueryRowxContext(pgs.Ctx, sql, eventUUID.String()).MapScan(res)
	if err != nil {
		return event.Event{}, err
	}
	e := event.Event{
		UUID:     eventUUID,
		Title:    res["title"].(string),
		Datetime: res["datetime"].(string),
		Duration: res["duration"].(string),
		Desc:     res["description"].(string),
		User:     res["userid"].(string),
		Notify:   res["notify"].(string),
	}
	return e, nil
}

// GetUserEvents gets all events attributed to specific user
func (pgs *PGStorage) GetUserEvents(user string) ([]event.Event, error) {
	return []event.Event{}, nil
}

// CreateEvent generates uuid for event object and saves it to storage
func (pgs *PGStorage) CreateEvent(e event.Event) (event.Event, error) {
	sql := `insert into events(uuid, title, datetime, duration, description, userid, notify) 
		values(:uuid, :title, :datetime, :duration, :description, :userid, :notify)`
	eventUUID := uuid.New()
	_, err := pgs.DB.NamedExecContext(pgs.Ctx, sql, map[string]interface{}{
		"uuid":        eventUUID.String(),
		"title":       e.Title,
		"datetime":    e.Datetime,
		"duration":    e.Duration,
		"description": e.Desc,
		"userid":      e.User,
		"notify":      e.Notify,
	})
	if err != nil {
		return event.Event{}, err
	}
	e.UUID = eventUUID
	return e, nil
}

// UpdateEvent rewrites event with given UUID
func (pgs *PGStorage) UpdateEvent(eventUUID uuid.UUID, e event.Event) (event.Event, error) {
	sql := `update events set 
		title = :title, 
		datetime = :datetime, 
		duration = :duration, 
		description = :description, 
		userid = :userid, 
		notify = :notify
	where uuid = :uuid`
	_, err := pgs.DB.NamedExecContext(pgs.Ctx, sql, map[string]interface{}{
		"uuid":        eventUUID.String(),
		"title":       e.Title,
		"datetime":    e.Datetime,
		"duration":    e.Duration,
		"description": e.Desc,
		"userid":      e.User,
		"notify":      e.Notify,
	})
	if err != nil {
		return event.Event{}, err
	}
	return e, nil
}

// DeleteEvent deletes event with given UUID from storage
func (pgs *PGStorage) DeleteEvent(eventUUID uuid.UUID) error {
	sql := "delete from events where uuid = $1"
	_, err := pgs.DB.ExecContext(pgs.Ctx, sql, eventUUID.String())
	return err
}
