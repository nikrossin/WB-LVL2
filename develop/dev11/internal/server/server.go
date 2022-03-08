package server

import (
	"lvl2/develop/dev11/internal/calendar"
	"net/http"
)

type Server struct {
	Config   *Config
	Calendar calendar.CalendarMem
	router   *http.ServeMux
}
