package handler

import (
	"net/http"

	"github.com/enchik0reo/wildberriesL2/develop/dev11/repos"
)

type Handler struct {
	repo *repos.Store
}

func New(repository *repos.Store) *Handler {
	return &Handler{
		repo: repository,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.index)
	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/events_for_day", h.eventsForDay)
	mux.HandleFunc("/events_for_week", h.eventsForWeek)
	mux.HandleFunc("/events_for_month", h.eventsForMonth)

	return mux
}
