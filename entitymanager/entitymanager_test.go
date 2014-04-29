package entitymanager

import "testing"

func TestStart(t *testing.T) {
	t.Logf("Starting up EM.")
	Start()
	if len(Creatures) != initial_creatures {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures, len(Creatures))
	}
}

func TestBreedRandom(t *testing.T) {
	Start()

	t.Log("Breeding new creature.")
	breedRandom()
	if len(Creatures) != initial_creatures+1 {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures+1, len(Creatures))
	}

	t.Log("Breeding new creature.")
	breedRandom()
	if len(Creatures) != initial_creatures+2 {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures+2, len(Creatures))
	}
}

func TestSpin(t *testing.T) {
	Start()

	t.Log("Spinning one cycle. Number of creatures shouldn't change.")
	Spin()
	if len(Creatures) != initial_creatures {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures, len(Creatures))
	}

	t.Logf("Fast forwarding %v cycles.", breeding_rate)
	for i := 0; i < breeding_rate; i++ {
		Spin()
	}
	if len(Creatures) != initial_creatures+1 {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures+1, len(Creatures))
	}
}
