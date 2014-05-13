package entitymanager

import "testing"

func TestBreedRandom(t *testing.T) {
	m := New().(*em)
	m.Reset()

	t.Log("Breeding new creature.")
	m.breedRandom()
	if len(m.creatures) != initial_creatures+1 {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures+1, len(m.creatures))
	}

	t.Log("Breeding new creature.")
	m.breedRandom()
	if len(m.creatures) != initial_creatures+2 {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures+2, len(m.creatures))
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
