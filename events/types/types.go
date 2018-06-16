package types

type SubscriptionFilter func(e *Event) bool

type Subscriber interface {
	Subscribe(SubscriptionFilter)
	Notify(*Event)
	Match(*Event) bool
}

type SubscriptionRegistry interface {
	RegisterSubscriber(Subscriber)
	Publish(*Event)
}

type Transition struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Event struct {
	Id        int         `json:"case_id"`
	Timestamp string      `json:"timestamp"`
	Assignee  string      `json:"assignee"`
	Team      string      `json:"team"`
	State     *Transition `json:"state"`
}
