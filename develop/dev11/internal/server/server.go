package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server Структура Сервера
type Server struct {
	Config *Config
	Hl     *Handler
	router *http.ServeMux
}

// NewServer Создание Сервера
func NewServer(c *Config) *Server {
	return &Server{
		c,
		NewHandler(),
		new(http.ServeMux),
	}
}

// SetRoutes Установить маршрутизацию запросов
func (s *Server) SetRoutes() http.Handler {
	s.router.HandleFunc("/create_event", s.Hl.AddEvent)
	s.router.HandleFunc("/update_event", s.Hl.UpdateEvent)
	s.router.HandleFunc("/delete_event", s.Hl.DeleteEvent)
	s.router.HandleFunc("/events_for_day", s.Hl.GetEventForDay)
	s.router.HandleFunc("/events_for_week", s.Hl.GetEventForWeek)
	s.router.HandleFunc("/events_for_month", s.Hl.GetEventForMonth)
	return s.Logging(s.router)
}

// Logging Логгирование middleware
func (s *Server) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s  %s  %s  %s\n", r.Method, r.URL, timeStart.Format("2006-01-02T15:04:05"), time.Since(timeStart))
	})
}

// Run Запуск сервера
func (s *Server) Run() {

	serv := http.Server{
		Addr:    s.Config.GetAddress(),
		Handler: s.SetRoutes(),
	}
	c := make(chan os.Signal, 1)
	err := make(chan error)
	signal.Notify(c, os.Interrupt)

	go func() {
		err <- serv.ListenAndServe()
	}()
	log.Println("Server start")
	select {
	case <-c:
		log.Println("Shutdown server")
		serv.Shutdown(context.Background())

	case er := <-err:
		log.Printf("Shutdown server by error: %v\n", er)
		serv.Shutdown(context.Background())
	}

}
