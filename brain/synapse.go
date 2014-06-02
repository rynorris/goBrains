package brain

import "github.com/DiscoViking/goBrains/config"

// A Synapse is a channel for transferring ChargeCarrier between brain elements.
// It has exactly one output, and should be pointed to by exactly one input.
// Unlike Nodes, Synapses do not distribute ChargeCarrier in bursts, they slowly
// convey ChargeCarrier to their output over time.
type Synapse struct {
	ChargeCarrier

	// What fraction of ChargeCarrier a Synapse will pass through from it's
	// inputs to it's outputs.
	// Generally between -1.0 and 1.0.
	permittivity float64

	// Where this Synapse will pass ChargeCarrier on to.
	// Synapses can have only one output.
	output Chargeable
}

// Work should be called once per tick.
// When a Synapse does work, it convey's ChargeCarrier to it's output
// at a rate relative to the Synapse's currentCharge.
func (s *Synapse) Work() {
	if s.currentCharge > config.Global.Brain.SynapseMaxCharge {
		s.currentCharge = config.Global.Brain.SynapseMaxCharge
	}

	if s.currentCharge != 0 {
		s.output.Charge(s.currentCharge *
			s.permittivity *
			config.Global.Brain.SynapseOutputScale)
	}
	s.Decay()
}

// Creates a new Synapse which points at the given output.
func NewSynapse(output Chargeable, permittivity float64) *Synapse {
	return &Synapse{output: output, permittivity: permittivity}
}
