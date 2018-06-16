package assign

import (
	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

func New() types.Subscriber {
	return new(AssignmentSubscriber).Init()
}
