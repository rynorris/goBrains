package entitymanager

import "testing"

var initial_entities = initial_creatures + initial_food

func TestGetLM(t *testing.T) {
	m := New()
	lm := m.LocationManager()
	t.Logf("Successfully got LM from EM.\nDetails: %#v", lm)
}

func TestReset(t *testing.T) {
	m := New()
	m.Reset()
	if len(m.Entities()) != initial_entities {
		t.Errorf("Expected %v creatures, got %v.", initial_entities, len(m.Entities()))
	}

	// Mess up the state of EM.
	em := m.(*em)
	em.breedRandom()
	em.spawnFood()

	// And reset.
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

	t.Logf("Forcing breeding cycle.")
	m.(*em).breeding_timer = breeding_rate
	m.Spin()

	if len(m.(*em).creatures) != initial_creatures+1 {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures+1, len(m.(*em).creatures))
	}

	t.Logf("Forcing food spawn cycle.")
	m.(*em).food_timer = food_replenish_rate
	m.Spin()

	if len(m.(*em).food) != initial_food+1 {
		t.Errorf("Expected %v food, got %v.", initial_food+1, len(m.(*em).food))
	}
}
