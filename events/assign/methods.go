package assign

import (
	"container/list"

	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

func (s *AssignmentSubscriber) Init() types.Subscriber {
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

func (s *AssignmentSubscriber) Notify(e *list.Element) {
}
