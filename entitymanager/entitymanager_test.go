package entitymanager

import "testing"

func TestReset(t *testing.T) {
	t.Logf("m.Reseting up EM.")
	m := New()
	m.Reset()
	if len(m.creatures) != initial_creatures {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures, len(m.creatures))
	}
}

func TestSpin(t *testing.T) {
	m := New()
	m.Reset()

	t.Log("m.Spinning one cycle. Number of creatures shouldn't change.")
	m.Spin()
	if len(m.creatures) != initial_creatures {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures, len(m.creatures))
	}

	t.Logf("Fast forwarding %v cycles.", breeding_rate)
	for i := 0; i < breeding_rate; i++ {
		m.Spin()
	}
	if len(m.creatures) != initial_creatures+1 {
		t.Errorf("Expected %v creatures, got %v.", initial_creatures+1, len(m.creatures))
	}
}
