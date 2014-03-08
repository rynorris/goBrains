package brain

// The proportion of a Synapse's charge which it will attempt to convey
// to it's output each time it does work.
// Note: the actual conveyed charge will be further modified by permittivity.
const synapseOutputScale = 0.1

// The maximum charge a synapse can hold. If it is charged beyond this limit
// any extra charge is lost.
const synapseMaxCharge = 1.0

// A Synapse is a channel for transferring Charge between brain elements.
// It has exactly one output, and should be pointed to by exactly one input.
// Unlike Nodes, Synapses do not distribute charge in bursts, they slowly
// convey charge to their output over time.
type Synapse struct {
    Charge

    // What fraction of charge a Synapse will pass through from it's
    // inputs to it's outputs.
    // Generally between -1.0 and 1.0.
    permittivity float64

    // Where this Synapse will pass charge on to.
    // Synapses can have only one output.
    output Chargeable
}

// Work should be called once per tick.
// When a Synapse does work, it convey's charge to it's output
// at a rate relative to the Synapse's currentCharge.
func (s *Synapse) Work() {
    if s.currentCharge > synapseMaxCharge {
        s.currentCharge = synapseMaxCharge
    }

    s.output.ChargeUp(s.currentCharge * ChargeUnit(s.permittivity) * synapseOutputScale)
    s.Decay()
}

// Creates a new Synapse which points at the given output.
func NewSynapse(output Chargeable, permittivity float64) *Synapse {
    return &Synapse{output: output, permittivity: permittivity}
}
