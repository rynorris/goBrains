package entitymanager

import (
	"testing"
	"time"

	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/events"
)

var (
	initial_creatures   = 40
	initial_food        = 30
	breeding_rate       = 150
	food_replenish_rate = 200
	initial_entities    = initial_creatures + initial_food
)

func loadTestConfig() {
	config.Load("../config/test_config.gcfg")
	config.Global.Entity.InitialCreatures = initial_creatures
	config.Global.Entity.InitialFood = initial_food
	config.Global.Entity.BreedingRate = breeding_rate
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
	loadTestConfig()
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

	t.Logf("Forcing stats report cycle.")
	success_chan := make(chan struct{}, 1)
	events.Global.Register(events.POPULATION_STATE, func(ev events.Event) {
		success_chan <- struct{}{}
	})
	m.(*em).stats_timer = 100
	m.Spin()
	select {
	case <-time.After(100 * time.Millisecond):
		t.Errorf("Timed out waiting for POPULATION_STATE event.")
	case <-success_chan:
		t.Logf("Received POPULATION_STATE_EVENT")
	}
}
