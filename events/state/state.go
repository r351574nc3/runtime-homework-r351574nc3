package state

import (
	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

const (
	CLOSED  = "closed"  // closed state
	OPEN    = "open"    // open state
	PENDING = "pending" // pending state
)

func New() types.Subscriber {
	return new(TransitionSubscriber).Init()
}
