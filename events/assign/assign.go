package assign

import (
	"fmt"

	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

type AssignmentSubscriber struct {
	filter types.SubscriptionFilter
}

func (s *AssignmentSubscriber) Init() types.Subscriber {
	s.Subscribe(func(event *types.Event) bool {
		return event.Team != "" && event.Assignee != ""
	})
	return s
}

func New() types.Subscriber {
	return new(AssignmentSubscriber).Init()
}

func (s *AssignmentSubscriber) Match(event *types.Event) bool {
	return s.filter(event)
}

func (s *AssignmentSubscriber) Subscribe(filter types.SubscriptionFilter) {
	s.filter = filter
}

func (s *AssignmentSubscriber) Notify(event *types.Event) {
	fmt.Printf("Reassigning assignee: %s, team: %s\n", event.Assignee, event.Team)
}
