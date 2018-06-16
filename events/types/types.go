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
	Filter(f DurationIndexSearchFilter) ([]StateEventSummary, bool)
}

type Transition struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Event struct {
	Id        int         `form:"case_id" json:"case_id" binding:"required"`
	Timestamp time.Time   `form:"timestamp" json:"timestamp" binding:"required"`
	Assignee  string      `form:"assignee" json:"assignee" binding:"required"`
	Team      string      `form:"team" json:"team" binding:"required"`
	State     *Transition `form:"state" json:"state" binding:"required"`
}
