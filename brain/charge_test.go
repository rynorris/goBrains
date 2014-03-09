package brain

import "testing"

import "../testutils"

func TestDefault(t *testing.T) {
	c := ChargeCarrier{}

	if c.currentCharge != 0 {
		t.Errorf("Expected 0 default ChargeCarrier. Got %v", c.currentCharge)
	}
}

func TestCharge(t *testing.T) {
	c := ChargeCarrier{}
	c.Charge(0.5)

	if !testutils.FloatsAreEqual(c.currentCharge, 0.5) {
		t.Errorf("Charged by 0.5, but have %v ChargeCarrier.", c.currentCharge)
	}
}

func TestDecay(t *testing.T) {
	c := ChargeCarrier{}
	c.Charge(0.5)
	c.Decay()

	if !testutils.FloatsAreEqual(c.currentCharge, 0.48) {
		t.Errorf("Should have had 0.48 ChargeCarrier after Decay. Got %v", c.currentCharge)
	}
}
