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

func (e *Event) ValidateToCreate() bool {
	if e.UserId == "" || e.Title == "" {
		return false
	}
	//var err error
	/*if _, err = time.Parse("2006-01-02", e.Date.String()); err != nil {
		return false
	}*/
	return true
}

func (e *Event) ValidateToUpdate() bool {
	if e.Uid == "" {
		return false
	}
	if !e.Date.IsZero() {
		var err error
		if _, err = time.Parse("2006-01-02", e.Date.String()); err != nil {
			return false
		}
	}
	return true
}

func (e *Event) ValidateToDelete() bool {
	if e.Uid == "" {
		return false
	}
	return true
}
