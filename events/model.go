package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	EVENT_FILE_KEY = "EVENT_FILE"
)

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

// Read events as JSON from a file. This usually happens at startup.
func LoadEvents() []Event {
	raw, err := ioutil.ReadFile(os.Getenv(EVENT_FILE_KEY))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Event
	json.Unmarshal(raw, &c)
	return c
}
