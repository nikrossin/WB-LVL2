package calendar

import (
	"github.com/google/uuid"
	"strings"
	"sync"
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
			c.events[id] = e
			break
		}
	}
	c.Unlock()
}

func (c *CalendarMem) UpdateEvent(e *Event, uid string) error {
	if err, _ := c.GetEvent(uid); err != nil {
		return err
	}
	c.Lock()
	c.events[uid] = e
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
func (c *CalendarMem) newID() string {
	uid := uuid.New()
	return strings.Replace(uid.String(), "-", "", -1)
}
