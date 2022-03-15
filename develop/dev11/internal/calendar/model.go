package calendar

import (
	"time"
)

// Event Модель события для хранения в кеше
type Event struct {
	Uid         string    `json:"uid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      string    `json:"user_id"`
	Date        time.Time `json:"date"`
}

// NewEvent Создание нового пустого события
func NewEvent() *Event {
	return new(Event)
}
