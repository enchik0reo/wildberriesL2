package repos

import (
	"sync"
	"time"
)

type Event struct {
	Date time.Time
	Id   int
	Msg  string
}

type Store struct {
	mu *sync.Mutex
	mp map[int]Event
}
