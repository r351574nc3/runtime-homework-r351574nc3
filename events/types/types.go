package types

import (
	"container/list"
	"time"
)

type SubscriptionFilter func(e Event) bool

type DurationIndexSearchFilter func(e StateEventSummary) bool

type Subscriber interface {
	Subscribe(SubscriptionFilter)
	Notify(*list.Element)
	Match(Event) bool
}

type SubscriptionRegistry interface {
	RegisterSubscriber(Subscriber)
	Publish(Event)
}

type StateEventSummary struct {
	Id       int
	Assignee string
	Team     string
	Status   string
	Duration time.Duration
}

type SearchableDurationIndex struct {
	Items []StateEventSummary
}

type Searchable interface {
	Filter(f DurationIndexSearchFilter) []StateEventSummary
}

type Transition struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Event struct {
	Id        int         `json:"case_id"`
	Timestamp time.Time   `json:"timestamp"`
	Assignee  string      `json:"assignee"`
	Team      string      `json:"team"`
	State     *Transition `json:"state"`
}
