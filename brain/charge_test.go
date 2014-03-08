package brain

import "testing"

func TestDefault(t *testing.T) {
    c := Charge{}

    if c.currentCharge != 0 {
        t.Errorf("Expected 0 default charge. Got %v", c.currentCharge)
    }
}

func TestCharge(t *testing.T) {
    c := Charge{}
    c.ChargeUp(0.5)

    if !ChargesAreEqual(c.currentCharge, 0.5) {
        t.Errorf("Charged by 0.5, but have %v charge.", c.currentCharge)
    }
}

func TestDecay(t *testing.T) {
    c := Charge{}
    c.ChargeUp(0.5)
    c.Decay()

    if !ChargesAreEqual(c.currentCharge, 0.48) {
        t.Errorf("Should have had 0.48 charge after Decay. Got %v", c.currentCharge)
    }
}
