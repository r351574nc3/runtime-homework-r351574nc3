package events

import (
	"os"
	"testing"

	"github.com/heroku/runtime-homework-r351574nc3/events/store"
)

func TestLoadEvents(t *testing.T) {
	os.Setenv("EVENT_FILE", "test_data/events.json")
	events := LoadEvents()

	expected := 24
	if len(events) != expected {
		t.Errorf("Expected %d events but got %d", expected, len(events))
	}
}

func TestReplay(t *testing.T) {
	events := LoadEvents()

	Replay(events)

	if len(store.DurationIndex.Items) != 21 {
		t.Errorf("Expected %d durations, but got %d", 34, len(store.DurationIndex.Items))
	}
}
