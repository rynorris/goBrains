package brain

import "testing"

import "../testutils"

func TestSynapseNew(t *testing.T) {
	n := NewNode()
	s := NewSynapse(n, 0)

	if !testutils.FloatsAreEqual(s.permittivity, 0) {
		t.Errorf("Expected 0 default permittivity, got %v", s.permittivity)
	}

	if s.output != n {
		t.Error("Node was not correctly added as output.")
	}
}

func TestSynapseWork(t *testing.T) {
	n := NewNode()
	s := NewSynapse(n, 0)

	s.Charge(0.5)

	s.Work()

	if !testutils.FloatsAreEqual(n.currentCharge, 0) {
		t.Errorf("Expected 0 ChargeCarrier passed to node since s.permittivity should be 0.\nGot %v", n.currentCharge)
	}

	if !testutils.FloatsAreEqual(s.currentCharge, 0.48) {
		t.Errorf("Expected 0.48 ChargeCarrier after Decay step. Got %v", s.currentCharge)
	}

	// Now try with some permittivity
	s.permittivity = 0.8

	s.Work()

	if !testutils.FloatsAreEqual(n.currentCharge, 0.48*s.permittivity*synapseOutputScale) {
		t.Errorf("Should have passed on all %v ChargeCarrier. Got %v",
			0.5*s.permittivity,
			n.currentCharge)
	}
}
