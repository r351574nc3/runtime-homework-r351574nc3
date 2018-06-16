package assign

import (
	"container/list"

	"github.com/heroku/runtime-homework-r351574nc3/events/store"
	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

func (s *AssignmentSubscriber) Init() types.Subscriber {
	// Only receive notifications if team or assignee was changed
	s.Subscribe(func(event types.Event) bool {
		return event.Team != "" && event.Assignee != ""
	})
	return s
}

func (s *AssignmentSubscriber) Match(event types.Event) bool {
	return s.filter(event)
}

func (s *AssignmentSubscriber) Subscribe(filter types.SubscriptionFilter) {
	s.filter = filter
}

// Received a notification that the Support Case state has changed. Either an Assignee or Team has been
// modified. Record the length of the previous state by comparing timestamps and index the change.
func (s *AssignmentSubscriber) Notify(e *list.Element) {
	event := e.Value.(types.Event)

	if previous, ok := store.GetPreviousEventFor(e); ok {
		duration := event.Timestamp.Sub(previous.Timestamp)

		// Update the duration index with state change
		store.DurationIndex.Items = append(store.DurationIndex.Items, types.StateEventSummary{
			Id:       previous.Id,
			Assignee: previous.Assignee,
			Team:     previous.Team,
			Status:   previous.State.To,
			Duration: duration,
		})
	}
}
