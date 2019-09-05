package event

import (
	"github.com/google/uuid"
)

// Event is an event model for go-calendar
type Event struct {
	UUID     uuid.UUID `db:"uuid"`
	Title    string
	Datetime string
	Duration string
	Desc     string `db:"description"`
	User     string `db:"userid"`
	Notify   string
}
