package server

import (
	"encoding/json"
	"fmt"
	"lvl2/develop/dev11/internal/calendar"
	"net/http"
	"time"
)

type Handler struct {
	Calendar *calendar.CalendarMem
}

func NewHandler() *Handler {
	return &Handler{
		calendar.NewCalendarMem(),
	}
}

func (h *Handler) AddEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responseJSON(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("method %v not allowed", r.Method))
		return
	}

	ev := &eventParse{}
	if err := json.NewDecoder(r.Body).Decode(ev); err != nil {
		responseJSON(true, w, http.StatusBadRequest, err.Error())
		return
	}
	if !ev.ValidateToCreate() {
		responseJSON(true, w, http.StatusBadRequest, "User_id is empty or Time is not formated")
		return
	}
	event := ev.ConvToModel()

	h.Calendar.CreateEvent(event)
	responseJSON(false, w, http.StatusOK, "Event create")
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responseJSON(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("method %v not allowed", r.Method))
		return
	}

	ev := &eventParse{}
	if err := json.NewDecoder(r.Body).Decode(ev); err != nil {
		responseJSON(true, w, http.StatusBadRequest, err.Error())
		return
	}
	if !ev.ValidateToUpdate() {
		responseJSON(true, w, http.StatusBadRequest, "Time is not formatted")
		return
	}

	event := ev.ConvToModel()
	if err := h.Calendar.UpdateEvent(event, event.Uid); err != nil {
		responseJSON(true, w, http.StatusBadRequest, "With current UID is not found Event")
		return
	}
	responseJSON(false, w, http.StatusOK, "Event update")
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responseJSON(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("method %v not allowed", r.Method))
		return
	}

	ev := &eventParse{}
	if err := json.NewDecoder(r.Body).Decode(ev); err != nil {
		responseJSON(true, w, http.StatusBadRequest, err.Error())
		return
	}
	if !ev.ValidateToDelete() {
		responseJSON(true, w, http.StatusBadRequest, "UID is empty")
		return
	}

	if err := h.Calendar.DeleteEvent(ev.Uid); err != nil {
		responseJSON(true, w, http.StatusBadRequest, err.Error())
		return
	}
	responseJSON(false, w, http.StatusOK, "Event delete")
}

func (h *Handler) GetEventForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		responseJSON(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("method %v not allowed", r.Method))
		return
	}
	if !r.URL.Query().Has("user_id") || !r.URL.Query().Has("date") {
		responseJSON(true, w, http.StatusBadRequest, "Not enough parameters")
		return
	}
	userIdParam := r.URL.Query().Get("user_id")
	dateParam := r.URL.Query().Get("date")

	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		responseJSON(true, w, http.StatusBadRequest, err.Error())
		return
	}
	events := h.Calendar.EventDay(userIdParam, date)
	responseJSON(false, w, http.StatusOK, events)
}

func (h *Handler) GetEventForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		responseJSON(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("method %v not allowed", r.Method))
		return
	}
	if !r.URL.Query().Has("user_id") || !r.URL.Query().Has("date") {
		responseJSON(true, w, http.StatusBadRequest, "Not enough parameters")
		return
	}
	userIdParam := r.URL.Query().Get("user_id")
	dateParam := r.URL.Query().Get("date")

	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		responseJSON(true, w, http.StatusBadRequest, err.Error())
		return
	}
	events := h.Calendar.EventWeek(userIdParam, date)
	responseJSON(false, w, http.StatusOK, events)
}

func (h *Handler) GetEventForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		responseJSON(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("method %v not allowed", r.Method))
		return
	}
	if !r.URL.Query().Has("user_id") || !r.URL.Query().Has("date") {
		responseJSON(true, w, http.StatusBadRequest, "Not enough parameters")
		return
	}
	userIdParam := r.URL.Query().Get("user_id")
	dateParam := r.URL.Query().Get("date")

	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		responseJSON(true, w, http.StatusBadRequest, err.Error())
		return
	}
	events := h.Calendar.EventMonth(userIdParam, date)
	responseJSON(false, w, http.StatusOK, events)
}
