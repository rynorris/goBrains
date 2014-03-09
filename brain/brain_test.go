package brain

import "testing"
import "math"

import "../testutils"

type testOutput struct {
	ChargeCarrier
}

func (o *testOutput) Work() {}

func TestBrainNew(t *testing.T) {
	b := NewBrain(5)

	if len(b.centralNodes) != 5 {
		t.Errorf("Created a brain with 5 central nodes. It only had %v", len(b.centralNodes))
	}
}

func TestBrainPropogation(t *testing.T) {
	b := NewBrain(4)
	in := NewNode()
	b.AddInputNode(in)
	out := testOutput{}
	b.AddOutput(&out)

	chargePerFire := float64(defaultFiringStrength - chargeDecayRate)
	expectedCharge := float64(0)

	// Give synapses some permittivity to allow propogation
	for _, synapse := range b.inSynapses {
		synapse.permittivity = 1.0
	}
	for _, synapse := range b.outSynapses {
		synapse.permittivity = 1.0
	}

	// ChargeCarrier the input, causing it to fire.
	in.Charge(defaultFiringThreshold)

	// Cause the brain to update, propogating ChargeCarrier.
	b.Update()

	// Check propogation to synapses
	// They should have been charged defaultFiringStrength by the input, and then
	// decayed chargeDecayRate. For a final total of chargePerFire.
	for _, synapse := range b.inSynapses {
		if !testutils.FloatsAreEqual(synapse.currentCharge, chargePerFire) {
			t.Errorf("In Synapse should have been charged to %v. Got %v", chargePerFire, synapse.currentCharge)
		}
	}

	// Check propogation to central nodes
	// They will have been charged, and then decayed by this point.
	expectedCharge = defaultFiringStrength*0.1 - chargeDecayRate
	for _, node := range b.centralNodes {
		if !testutils.FloatsAreEqual(node.currentCharge, expectedCharge) {
			t.Errorf("Node should have %v ChargeCarrier. Got %v", expectedCharge, node.currentCharge)
		}
	}

	// Repeatedly cause the input to fire.
	// This will keep the synapses permenantly at synapseMaxCharge ChargeCarrier
	// They will ChargeCarrier the central nodes by 0.08 per loop (including decay).
	// We want to get the central nodes to fire, so we need to ChargeCarrier them
	// synapseMaxCharge / 0.06 ~ 20 times
	loopsToFire := int(math.Ceil(1.0/(float64(synapseMaxCharge)*0.1-chargeDecayRate))) - 1
	for i := 0; i < loopsToFire; i++ {
		in.Charge(defaultFiringThreshold)
		b.Update()
	}

	// Check propogation to synapses
	// They should have been charged defaultFiringStrength by the central nodes, and then
	// decayed chargeDecayRate. For a final total of chargePerFire.
	expectedCharge = defaultFiringStrength - chargeDecayRate
	for _, synapse := range b.outSynapses {
		if !testutils.FloatsAreEqual(synapse.currentCharge, expectedCharge) {
			t.Errorf("Out Synapse should have been charged to %v. Got %v", expectedCharge, synapse.currentCharge)
		}
	}

	// Check propogation to output
	// It should have been charged defaultFiringStrength * 0.1 by each synapse. So in total gained 3.2 ChargeCarrier.
	expectedCharge = defaultFiringStrength * synapseOutputScale * 4
	if !testutils.FloatsAreEqual(out.currentCharge, expectedCharge) {
		t.Errorf("Output should have %v total ChargeCarrier. Got %v", expectedCharge, out.currentCharge)
	}
}
