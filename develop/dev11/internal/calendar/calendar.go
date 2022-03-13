package calendar

import (
	"github.com/google/uuid"
	"strings"
	"sync"
	"time"
)

type CalendarMem struct {
	*sync.RWMutex
	events map[string]*Event
}

func NewCalendarMem() *CalendarMem {
	return &CalendarMem{
		&sync.RWMutex{},
		make(map[string]*Event),
	}
}
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

func (c *CalendarMem) UpdateEvent(e *Event, uid string) error {
	if err, in := c.GetEvent(uid); err != nil {
		return err
	} else {
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
	}
	c.Unlock()
	return nil
}

func (c *CalendarMem) DeleteEvent(uid string) error {
	if err, _ := c.GetEvent(uid); err != nil {
		return err
	}
	c.Lock()
	delete(c.events, uid)
	c.Unlock()
	return nil
}
func (c *CalendarMem) GetEvent(uid string) (error, *Event) {
	c.RLock()
	defer c.RUnlock()
	if e, ok := c.events[uid]; ok {
		return nil, e
	}
	return errNotFoundEvent, nil
}

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

func (c *CalendarMem) newID() string {
	uid := uuid.New()
	return strings.Replace(uid.String(), "-", "", -1)
}
