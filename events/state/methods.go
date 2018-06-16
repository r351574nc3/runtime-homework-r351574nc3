package state

import (
	"fmt"

	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

func (s *TransitionSubscriber) Init() types.Subscriber {
	s.Subscribe(func(event *types.Event) bool {
		return event.State != nil
	})
	return s
}

func (s *TransitionSubscriber) Match(event *types.Event) bool {
	return s.filter(event)
}

func (s *TransitionSubscriber) Subscribe(filter types.SubscriptionFilter) {
	s.filter = filter
}

func (s *TransitionSubscriber) Notify(event *types.Event) {
	fmt.Printf("Transitioning state from: %s to: %s\n", event.State.From, event.State.To)
}
