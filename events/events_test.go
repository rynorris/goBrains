package events

import (
	"testing"
)

// Confirm all event types report correctly.
func TestEventTypes(t *testing.T) {
	var e Event

	// BasicEvent
	e = BasicEvent{TERMINATE}
	if e.GetType() != TERMINATE {
		t.Errorf("BasicEvent reports type %v instead of TERMINATE.\n", e.GetType())
	}

	// EntityEvent
	e = EntityEvent{ENTITY_CREATE, nil}
	if e.GetType() != ENTITY_CREATE {
		t.Errorf("BasicEvent reports type %v instead of ENTITY_CREATE.\n", e.GetType())
	}

	// PopulationEvent
	e = PopulationEvent{0, 10, 100}
	if e.GetType() != POPULATION_STATE {
		t.Errorf("BasicEvent reports type %v instead of POPULATION_STATE.\n", e.GetType())
	}
}
