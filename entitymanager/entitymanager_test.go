package entitymanager

import (
	"math"
	"testing"
)

var initial_entities = initial_creatures + initial_food

func TestReset(t *testing.T) {
	t.Logf("m.Reseting up EM.")
	m := New()
	m.Reset()
	if len(m.Entities()) != initial_entities {
		t.Errorf("Expected %v creatures, got %v.", initial_entities, len(m.Entities()))
	}
}

func TestSpin(t *testing.T) {
	m := New()
	m.Reset()

	t.Log("m.Spinning one cycle. Number of creatures shouldn't change.")
	m.Spin()
	if len(m.Entities()) != initial_entities {
		t.Errorf("Expected %v creatures, got %v.", initial_entities, len(m.Entities()))
	}

	t.Logf("Fast forwarding %v cycles.", breeding_rate)
	for i := 0.0; i < math.Min(breeding_rate, food_replenish_rate); i++ {
		m.Spin()
	}

	if len(m.Entities()) != initial_entities+1 {
		t.Errorf("Expected %v creatures, got %v.", initial_entities+1, len(m.Entities()))
	}
}
