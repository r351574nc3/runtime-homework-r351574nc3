package state

import (
	"container/list"
	"time"

	"github.com/heroku/runtime-homework-r351574nc3/events/store"
	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

func (s *TransitionSubscriber) Init() types.Subscriber {
	s.Subscribe(func(event types.Event) bool {
		return event.State != nil
	})
	return s
}

func (s *TransitionSubscriber) Match(event types.Event) bool {
	return s.filter(event)
}

func (s *TransitionSubscriber) Subscribe(filter types.SubscriptionFilter) {
	s.filter = filter
}

func (s *TransitionSubscriber) Notify(e *list.Element) {
	var (
		duration       time.Duration
		previous_event *types.Event
	)
	event := e.Value.(types.Event)

	if previous, ok := store.GetPreviousEventFor(e); ok {
		duration = event.Timestamp.Sub(previous.Timestamp)
		previous_event = previous
	}

	// Update the duration index with state change
	store.DurationIndex.Items = append(store.DurationIndex.Items, types.StateEventSummary{
		Id:       previous_event.Id,
		Assignee: previous_event.Assignee,
		Team:     previous_event.Team,
		Status:   previous_event.State.To,
		Duration: duration,
	})
}
