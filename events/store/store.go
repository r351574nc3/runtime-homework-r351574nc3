package store

import (
	"container/list"
)

type EventStore struct {
	Events *list.List
}

func (s *EventStore) Init() *EventStore {
	s.Events = list.New()
	return s
}

func New() *EventStore {
	return new(EventStore).Init()
}
