package store

import "container/list"

type EventStore struct {
	Events        *list.List
	subscriptions *EventSubscriptions
}

type EventSubscriptions struct {
	subscribers *list.List
}
