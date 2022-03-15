package server

import (
	"lvl2/develop/dev11/internal/calendar"
	"time"
)

// Модель  для парсинга События из json
type eventParse struct {
	Uid         string `json:"uid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      string `json:"user_id"`
	Date        string `json:"date"`
}

// ConvToModel Конвертировать модель после парсинга в модель для хранения в кеше
func (ev *eventParse) ConvToModel() *calendar.Event {
	event := calendar.NewEvent()
	event.Uid = ev.Uid
	event.Title = ev.Title
	event.Description = ev.Description
	event.UserId = ev.UserId
	event.Date, _ = time.Parse("2006-01-02", ev.Date)
	return event
}

// ValidateToCreate Валидация данных при запросе на создание события
func (ev *eventParse) ValidateToCreate() bool {
	if ev.UserId == "" || ev.Title == "" {
		return false
	}
	var err error
	if _, err = time.Parse("2006-01-02", ev.Date); err != nil {
		return false
	}
	return true
}

// ValidateToUpdate Валидация данных при запросе на обновление события
func (ev *eventParse) ValidateToUpdate() bool {
	if ev.Uid == "" {
		return false
	}
	if ev.Date != "" {
		var err error
		if _, err = time.Parse("2006-01-02", ev.Date); err != nil {
			return false
		}
	}
	return true
}

// ValidateToDelete Валидация данных при запросе на удаление события
func (ev *eventParse) ValidateToDelete() bool {
	if ev.Uid == "" {
		return false
	}
	return true
}
