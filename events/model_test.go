package events

import (
	"os"
	"testing"
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
}
