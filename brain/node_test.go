package brain

import "testing"

const delta = 0.0001

func TestNew(t *testing.T) {
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

func TestCharge(t *testing.T) {
    n := NewNode()
    n.Charge(0.5)

    if n.currentCharge-0.5 > delta {
        t.Errorf("After charging by 0.5, charge was %v instead of 0.5!", n.currentCharge)
    }

    n.Charge(0.7)
    if n.currentCharge-1.2 > delta {
        t.Errorf("Charge should now be 1.2. Got %v instead.", n.currentCharge)
    }
}

func TestUpdate(t *testing.T) {
    n := NewNode()
    n.Work()
    if n.currentCharge > 0 {
        t.Errorf("Charge should still be 0 after update. Got %v instead.", n.currentCharge)
    }

    n.Charge(0.5)
    n.Work()
    if n.currentCharge-0.48 > delta {
        t.Errorf("Charge should be 0.48 after update. Got %v instead.", n.currentCharge)
    }
}

func TestFire(t *testing.T) {
    n := NewNode()
    m := NewNode()
    n.AddOutput(m)

    n.Fire()
    if m.currentCharge-0.8 > delta {
        t.Errorf("m should have 0.8 charge after n fires. Got %v instead.", m.currentCharge)
    }
}

func TestOutput(t *testing.T) {
    n := NewNode()
    m := NewNode()
    n.AddOutput(m)

    n.Charge(1.2)
    n.Work()
    if m.currentCharge-0.8 > delta {
        t.Errorf("m should have 0.8 charge after n fires. Got %v instead.", m.currentCharge)
    }
}
