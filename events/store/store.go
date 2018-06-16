package store

import (
	"container/list"

	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

var (
	store *EventStore
)

type EventStore struct {
	Events        *list.List
	subscriptions *EventSubscriptions
}

func Setup(subscribers ...types.Subscriber) *EventStore {
	store = store.New()

	for _, subscriber := range subscribers {
		Register(subscriber)
	}

	return store
}

func Register(subscriber types.Subscriber) error {
	return store.subscriptions.RegisterSubscriber(subscriber)
}

func (s *EventStore) New() *EventStore {
	if s == nil {
		return new(EventStore).Init()
	}
	return s
}

func (s *EventStore) Init() *EventStore {
	s.Events = list.New()
	subscriptions := new(EventSubscriptions)
	subscriptions.subscribers = list.New()

	s.subscriptions = subscriptions

	return s
}

func (s *EventStore) Publish(event types.Event) {
	s.Events.PushBack(event)
	s.subscriptions.Publish(event)
}

type EventSubscriptions struct {
	subscribers *list.List
}

func (s *EventSubscriptions) RegisterSubscriber(subscriber types.Subscriber) error {
	s.subscribers.PushBack(subscriber)
	return nil
}

func (s *EventSubscriptions) Publish(event types.Event) {
	for e := s.subscribers.Front(); e != nil; e = e.Next() {
		subscriber := e.Value.(types.Subscriber)

		if subscriber.Match(&event) {
			subscriber.Notify(&event)
		}
	}
}
