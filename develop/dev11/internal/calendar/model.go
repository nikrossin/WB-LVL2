package calendar

import (
	"time"
)

type Event struct {
	Uid         string    `json:"uid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      string    `json:"user_id"`
	Date        time.Time `json:"date"`
}

func NewEvent() *Event {
	return new(Event)
}
