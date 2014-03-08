package brain

import "testing"

func TestNodeNew(t *testing.T) {
    n := NewNode()

    if n.firingThreshold != defaultFiringThreshold {
        t.Errorf("Default firingThreshold was %v, expected %v.", n.firingThreshold, defaultFiringThreshold)
    }

    if n.firingStrength != defaultFiringStrength {
        t.Errorf("Default firingStrength was %v, expected %v.", n.firingStrength, defaultFiringStrength)
    }

    if n.currentCharge != 0 {
        t.Errorf("Default currentCharge was %v, expected 0.", n.currentCharge)
    }
}

func TestNodeUpdate(t *testing.T) {
    n := NewNode()
    n.Work()
    if n.currentCharge > 0 {
        t.Errorf("Charge should still be 0 after update. Got %v instead.", n.currentCharge)
    }

    n.ChargeUp(0.5)
    n.Work()
    if !ChargesAreEqual(n.currentCharge, 0.48) {
        t.Errorf("Charge should be 0.48 after update. Got %v instead.", n.currentCharge)
    }
}

func TestNodeFire(t *testing.T) {
    n := NewNode()
    m := NewNode()
    n.AddOutput(m)

    n.Fire()
    if !ChargesAreEqual(m.currentCharge, 0.8) {
        t.Errorf("m should have 0.8 charge after n fires. Got %v instead.", m.currentCharge)
    }
}

func TestNodeOutput(t *testing.T) {
    n := NewNode()
    m := NewNode()
    n.AddOutput(m)

    n.ChargeUp(1.2)
    n.Work()
    if !ChargesAreEqual(m.currentCharge, 0.8) {
        t.Errorf("m should have 0.8 charge after n fires. Got %v instead.", m.currentCharge)
    }
}
