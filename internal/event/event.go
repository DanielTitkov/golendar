package event

import (
	"github.com/google/uuid"
)

// Event is an event model for go-calendar
type Event struct {
	UUID     uuid.UUID
	Title    string
	Datetime string
	Duration string
	Desc     string
	User     string
	Notify   string
}
