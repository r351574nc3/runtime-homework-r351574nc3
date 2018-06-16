package store

import (
	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

var (
	store *EventStore
)

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
