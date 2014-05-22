package entitymanager

import "testing"

func TestBreedRandom(t *testing.T) {
	m := New().(*em)

	// Test with no creatures.
	m.breedRandom()
	if len(m.creatures) != 0 {
		t.Errorf("Something happened when breeding with 0 creatures in pool.\nCreatures: %#v", m.creatures)
	}
	m.Reset()

	for i := 1; i < 20; i++ {
		t.Log("Breeding new creature.")
		m.breedRandom()
		if len(m.creatures) != initial_creatures+i {
			t.Errorf("Expected %v creatures, got %v.", initial_creatures+i, len(m.creatures))
		}
	}
}

func TestSpawnFood(t *testing.T) {
	m := New().(*em)
	m.Reset()

	t.Log("Spawning food.")
	m.spawnFood()
	if len(m.food) != initial_food+1 {
		t.Errorf("Expected %v food, got %v.", initial_food+1, len(m.creatures))
	}

	t.Log("Spawning food.")
	m.spawnFood()
	if len(m.food) != initial_food+2 {
		t.Errorf("Expected %v food, got %v.", initial_food+2, len(m.creatures))
	}
}
