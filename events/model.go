package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/heroku/runtime-homework-r351574nc3/events/store"
)

const (
	EVENT_FILE_KEY = "EVENT_FILE"
)

var (
	Store = store.New()
)

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

func Replay(event_set []Event) {
	for _, event := range event_set {
		Store.Events.PushBack(event)
	}
}
