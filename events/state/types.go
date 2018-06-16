package state

import "github.com/heroku/runtime-homework-r351574nc3/events/types"

type TransitionSubscriber struct {
	filter types.SubscriptionFilter
}
