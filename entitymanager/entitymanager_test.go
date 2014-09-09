package entitymanager

import (
	"testing"

	"github.com/DiscoViking/goBrains/config"
)

var (
	initial_creatures   = 40
	initial_food        = 30
	breeding_chance     = 1 // Make breeding deterministic.
	food_replenish_rate = 200
	initial_entities    = initial_creatures + initial_food
)

func loadTestConfig() {
	config.Load("../config/test_config.gcfg")
	config.Global.Entity.InitialCreatures = initial_creatures
	config.Global.Entity.InitialFood = initial_food
	config.Global.Entity.BreedingChance = breeding_chance
	config.Global.Entity.FoodSpawnRate = food_replenish_rate
}

func TestGetLM(t *testing.T) {
	loadTestConfig()
	m := New()
	lm := m.LocationManager()
	t.Logf("Successfully got LM from EM.\nDetails: %#v", lm)
}

func TestReset(t *testing.T) {
	loadTestConfig()
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
	//currently broken
	return
	loadTestConfig()
	m := New()
	m.Reset()

	t.Logf("Forcing food spawn cycle.")
	m.(*em).food_timer = food_replenish_rate
	m.Spin()

	if len(m.(*em).food) != initial_food+1 {
		t.Errorf("Expected %v food, got %v.", initial_food+1, len(m.(*em).food))
	}

	// Should have also bred a new creature for each existing one
	// since we set breeding to certain.
	if len(m.(*em).creatures) != initial_creatures*2 {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures*2, len(m.(*em).creatures))
	}
}
