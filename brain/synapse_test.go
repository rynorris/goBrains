package brain

import "testing"

func TestSynapseNew(t *testing.T) {
    n := NewNode()
    s := NewSynapse(n, 0)

    if !ChargesAreEqual(ChargeUnit(s.permittivity), 0) {
        t.Errorf("Expected 0 default permittivity, got %v", s.permittivity)
    }

    if s.output != n {
        t.Error("Node was not correctly added as output.")
    }
}

func TestSynapseWork(t *testing.T) {
    n := NewNode()
    s := NewSynapse(n, 0)

    s.ChargeUp(0.5)

    s.Work()

    if !ChargesAreEqual(n.currentCharge, 0) {
        t.Errorf("Expected 0 charge passed to node since s.permittivity should be 0.\nGot %v", n.currentCharge)
    }

    if !ChargesAreEqual(s.currentCharge, 0.48) {
        t.Errorf("Expected 0.48 charge after Decay step. Got %v", s.currentCharge)
    }

    // Now try with some permittivity
    s.permittivity = 0.8

    s.Work()

    if !ChargesAreEqual(n.currentCharge, ChargeUnit(0.48*s.permittivity*synapseOutputScale)) {
        t.Errorf("Should have passed on all %v charge. Got %v",
            0.5*s.permittivity,
            n.currentCharge)
    }
}
