package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/heroku/runtime-homework-r351574nc3/events/assign"
	"github.com/heroku/runtime-homework-r351574nc3/events/state"
	"github.com/heroku/runtime-homework-r351574nc3/events/store"
	"github.com/heroku/runtime-homework-r351574nc3/events/types"
)

const (
	EVENT_FILE_KEY = "EVENT_FILE"
)

var (
	// Store = new(store.Builder).WithRegistry().Build()
	Store = store.Setup(assign.New(), state.New())
)

// Read events as JSON from a file. This usually happens at startup.
func LoadEvents() []types.Event {
	raw, err := ioutil.ReadFile(os.Getenv(EVENT_FILE_KEY))
	if err != nil {
		fmt.Println(err.Error())
	}

	var c []types.Event
	json.Unmarshal(raw, &c)

	sort.Sort(ByTimestamp(c)) // sort everything by timestamp
	return c
}

func Replay(event_set []types.Event) {
	for _, event := range event_set {
		Store.Publish(event)
	}
}
