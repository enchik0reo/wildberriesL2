package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmpl/index.html")
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	_, date, msg, err := post(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.repo.CreateEvent(date, msg)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	id, date, msg, err := post(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.UpdateEvent(id, date, msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	id, _, _, err := post(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.DeleteEvent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	date, err := get(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.repo.EventsForDay(date, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	writeJSON(w, events)
}

func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	date, err := get(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.repo.EventsForDay(date, 7)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	writeJSON(w, events)
}

func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	date, err := get(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.repo.EventsForDay(date, 30)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	writeJSON(w, events)
}

func post(r *http.Request) (int, time.Time, string, error) {
	var id int
	var date time.Time
	var msg string

	idString := r.FormValue("id")
	if idString != "" {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			return 0, time.Time{}, "", errors.New("400: invalid int")
		}
		id = idInt
	}

	dateString := r.FormValue("date")
	if dateString != "" {
		dateString += "T00:00:00Z"
		dateTime, err := time.Parse(time.RFC3339, dateString)
		if err != nil {
			return 0, time.Time{}, "", errors.New("400: invalid date")
		}
		date = dateTime
	} else {
		return 0, time.Time{}, "", errors.New("400: invalid date")
	}

	msg = r.FormValue("msg")

	return id, date, msg, nil
}

func get(r *http.Request) (time.Time, error) {
	var date time.Time
	var err error

	dateString := r.FormValue("date")

	if dateString != "" {
		dateString += "T00:00:00Z"
		date, err = time.Parse(time.RFC3339, dateString)
		if err != nil {
			return time.Time{}, err
		}
	} else {
		return time.Time{}, errors.New("400: invalid date")
	}
	return date, nil
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	resultJSON := struct {
		Result interface{} `json:"result"`
	}{Result: v}

	js, err := json.Marshal(&resultJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
