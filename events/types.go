package events

type Transition struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Event struct {
	Id        int        `json:"case_id"`
	Timestamp string     `json:"timestamp"`
	Assignee  string     `json:"assignee"`
	Team      string     `json:"team"`
	State     Transition `json:"state"`
}

type EventHandlerRegistry interface {
	Notify(*Event)
	Register(*EventHandler)
}

type EventHandler interface {
	Notify(*Event)
}
