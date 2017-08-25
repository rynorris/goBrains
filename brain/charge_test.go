package brain

import (
	"testing"

	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/testutils"
)

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
	config.Load("../config/test_config.gcfg")
	c := ChargeCarrier{}
	c.Charge(0.5)
	c.Decay()

	expected := 0.5 - config.Global.Brain.ChargeDecayRate
	if !testutils.FloatsAreEqual(c.currentCharge, expected) {
		t.Errorf("Should have had %v ChargeCarrier after Decay. Got %v", expected, c.currentCharge)
	}
}
