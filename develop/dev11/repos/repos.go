package repos

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

func NewStore(mu *sync.Mutex, mp map[int]Event) *Store {
	return &Store{mu: mu, mp: mp}
}

func (s *Store) CreateEvent(date time.Time, msg string) {
	id := len(s.mp)

	event := Event{date, id, msg}

	s.mu.Lock()
	defer s.mu.Unlock()

	if reflect.DeepEqual(s.mp[id], Event{}) {
		s.mp[id] = event
		id++
	}
}

func (s *Store) UpdateEvent(id int, date time.Time, msg string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if reflect.DeepEqual(s.mp[id], Event{}) {
		return errors.New("503: invalid event")
	}

	event := Event{date, id, msg}

	s.mp[id] = event

	return nil
}

func (s *Store) DeleteEvent(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if reflect.DeepEqual(s.mp[id], Event{}) {
		return errors.New("503: No event for delete")
	}

	delete(s.mp, id)

	return nil
}

func (s *Store) EventsForDay(date time.Time, days int) ([]Event, error) {
	var result []Event
	tt := 24 * days

	for _, event := range s.mp {
		t := int(time.Duration.Hours(event.Date.Sub(date)))
		if (t >= 0) && (t <= tt) {
			result = append(result, event)
		}
	}
	if len(result) == 0 {
		return []Event{}, errors.New("503: No event for date")
	}

	return result, nil
}
