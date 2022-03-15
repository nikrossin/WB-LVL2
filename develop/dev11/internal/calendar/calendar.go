package calendar

import (
	"github.com/google/uuid"
	"strings"
	"sync"
	"time"
)

// CalendarMem Кеш Календаря
type CalendarMem struct {
	*sync.RWMutex
	events map[string]*Event
}

// NewCalendarMem Создание кеша календаря
func NewCalendarMem() *CalendarMem {
	return &CalendarMem{
		&sync.RWMutex{},
		make(map[string]*Event),
	}
}

// CreateEvent Создание события для календаря
func (c *CalendarMem) CreateEvent(e *Event) {
	c.Lock()
	for {
		id := c.newID()
		if _, ok := c.events[id]; !ok {
			e.Uid = id
			c.events[id] = e
			break
		}
	}
	c.Unlock()
}

// UpdateEvent Обновление события в календаре
func (c *CalendarMem) UpdateEvent(e *Event, uid string) error {
	err, in := c.GetEvent(uid)
	if err != nil {
		return err
	}
	c.Lock()
	if e.Title != "" {
		in.Title = e.Title
	}
	if e.Description != "" {
		in.Description = e.Description
	}
	if e.UserId != "" {
		in.UserId = e.UserId
	}
	if !e.Date.IsZero() {
		in.Date = e.Date
	}
	c.Unlock()
	return nil
}

// DeleteEvent Удаление события календаря
func (c *CalendarMem) DeleteEvent(uid string) error {
	if err, _ := c.GetEvent(uid); err != nil {
		return err
	}
	c.Lock()
	delete(c.events, uid)
	c.Unlock()
	return nil
}

// GetEvent Получение События по id
func (c *CalendarMem) GetEvent(uid string) (error, *Event) {
	c.RLock()
	defer c.RUnlock()
	if e, ok := c.events[uid]; ok {
		return nil, e
	}
	return errNotFoundEvent, nil
}

// EventDay Получение событий пользователя на сутки по дате
func (c *CalendarMem) EventDay(user string, date time.Time) []*Event {
	var events []*Event
	c.RLock()
	for _, ev := range c.events {
		if ev.UserId == user && ev.Date.Equal(date) {
			events = append(events, ev)
		}
	}
	c.RUnlock()
	return events
}

// EventWeek Получение событий пользователя на неделю по дате
func (c *CalendarMem) EventWeek(user string, date time.Time) []*Event {
	var events []*Event
	c.RLock()
	for _, ev := range c.events {
		startTime := date
		endTime := date.AddDate(0, 0, 7)
		if ev.UserId == user && (ev.Date.After(startTime) && ev.Date.Before(endTime) || ev.Date.Equal(startTime) || ev.Date.Equal(endTime)) {
			events = append(events, ev)
		}
	}
	c.RUnlock()
	return events
}

// EventMonth Получение событий пользователя на месяц по дате
func (c *CalendarMem) EventMonth(user string, date time.Time) []*Event {
	var events []*Event
	c.RLock()
	for _, ev := range c.events {
		startTime := date
		endTime := date.AddDate(0, 1, 0)
		if ev.UserId == user && (ev.Date.After(startTime) && ev.Date.Before(endTime) || ev.Date.Equal(startTime) || ev.Date.Equal(endTime)) {
			events = append(events, ev)
		}
	}
	c.RUnlock()
	return events
}

// Генерация нового UID для события
func (c *CalendarMem) newID() string {
	uid := uuid.New()
	return strings.Replace(uid.String(), "-", "", -1)
}
