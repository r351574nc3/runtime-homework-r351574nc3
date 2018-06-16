package store

import (
	"container/list"

	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

var (
	store         *EventStore
	DurationIndex *types.SearchableDurationIndex
)

func Setup(subscribers ...types.Subscriber) *EventStore {
	DurationIndex = DurationIndex.New()

	store = store.New()

	for _, subscriber := range subscribers {
		Register(subscriber)
	}

	return store
}

func Register(subscriber types.Subscriber) error {
	return store.subscriptions.RegisterSubscriber(subscriber)
}

func GetPreviousEventFor(e *list.Element) (*types.Event, bool) {
	if current, ok := e.Value.(types.Event); ok {
		retval := new(types.Event)
		retval.Id = current.Id
		retval.State = new(types.Transition)

		is_ts_set := false

		for ; e != nil; e = e.Prev() {
			if previous, ok := e.Value.(types.Event); ok && previous.Id == current.Id {
				if !is_ts_set {
					retval.Timestamp = previous.Timestamp
				}
				if ok := previous.AssignTo(retval); ok {
					return retval, true
				}
			}
		}
	}

	return nil, false
}
