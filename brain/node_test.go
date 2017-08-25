package brain

import (
	"testing"

	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/testutils"
)

func TestNodeNew(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	n := NewNode()

	if n.firingThreshold != config.Global.Brain.NodeFiringThreshold {
		t.Errorf("Default firingThreshold was %v, expected %v.", n.firingThreshold, config.Global.Brain.NodeFiringThreshold)
	}

	if n.firingStrength != config.Global.Brain.NodeFiringStrength {
		t.Errorf("Default firingStrength was %v, expected %v.", n.firingStrength, config.Global.Brain.NodeFiringStrength)
	}

	if n.currentCharge != 0 {
		t.Errorf("Default currentCharge was %v, expected 0.", n.currentCharge)
	}
}

func TestNodeUpdate(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	n := NewNode()
	n.Work()
	if n.currentCharge > 0 {
		t.Errorf("ChargeCarrier should still be 0 after update. Got %v instead.", n.currentCharge)
	}

	n.Charge(0.5)
	n.Work()
	if !testutils.FloatsAreEqual(n.currentCharge, 0.48) {
		t.Errorf("ChargeCarrier should be 0.48 after update. Got %v instead.", n.currentCharge)
	}
}

func TestNodeFire(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	n := NewNode()
	m := NewNode()
	n.AddOutput(m)

	n.Fire()
	if !testutils.FloatsAreEqual(m.currentCharge, 0.8) {
		t.Errorf("m should have 0.8 ChargeCarrier after n fires. Got %v instead.", m.currentCharge)
	}
}

func TestNodeOutput(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	n := NewNode()
	m := NewNode()
	n.AddOutput(m)

	n.Charge(1.2)
	n.Work()
	if !testutils.FloatsAreEqual(m.currentCharge, 0.8) {
		t.Errorf("m should have 0.8 ChargeCarrier after n fires. Got %v instead.", m.currentCharge)
	}
}

func TestNegativeCharge(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	n := NewNode()
	m := NewNode()
	n.AddOutput(m)

	// Charge n negatively.
	n.Charge(-100)

	// Do work, n shouldn't charge m, and should have reset itself to 0.
	n.Work()

	if !testutils.FloatsAreEqual(m.currentCharge, 0) {
		t.Errorf("n had negative charge, but still charged m to %v", m.currentCharge)
	}

	if !testutils.FloatsAreEqual(n.currentCharge, 0) {
		t.Errorf("n should have capped its charge at 0, but it has %v", n.currentCharge)
	}
}
