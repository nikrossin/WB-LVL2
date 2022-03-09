package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Config *Config
	Hl     *Handler
	router *http.ServeMux
}

func NewServer(c *Config) *Server {
	return &Server{
		c,
		NewHandler(),
		new(http.ServeMux),
	}
}

func (s *Server) SetRoutes() http.Handler {
	s.router.HandleFunc("/create_event", s.Hl.AddEvent)
	s.router.HandleFunc("/update_event", s.Hl.UpdateEvent)
	s.router.HandleFunc("/delete_event", s.Hl.DeleteEvent)
	s.router.HandleFunc("/events_for_day", s.Hl.GetEventForDay)
	s.router.HandleFunc("/events_for_week", s.Hl.GetEventForWeek)
	s.router.HandleFunc("/events_for_month", s.Hl.GetEventForMonth)
	return s.Logging(s.router)
}

func (s *Server) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s\t%s\t%s\t%s\n", r.Method, r.URL, timeStart, time.Since(timeStart))
	})
}

func (s *Server) Run() {
	s.SetRoutes()
	serv := http.Server{
		Addr:    s.Config.GetAddress(),
		Handler: s.router,
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
