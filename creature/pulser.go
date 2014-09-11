/*
 * Pulsers.
 *
 * Pulsers fire consistently, providing the null-behaviour of the creature when no other inut is stimulated.
 */

package creature

import (
	"github.com/DiscoViking/goBrains/brain"
	"github.com/DiscoViking/goBrains/locationmanager"
)

// Fixed values.
const (
	// Charge incremented per detect.
	chargePulser = 0.1
)

// The pulser fires indiscriminately.
func (p *Pulser) detect() {
	p.node.Charge(chargePulser)
}

// Add a new pulser to a creature.
func (host *creature) AddPulser() *Pulser {

	// Link the pulser to the host's brain.
	node := brain.NewNode()
	host.brain.AddInputNode(node)

	input := inputStruct{
		putStruct: putStruct{host: host},
		node:      node,
		location:  locationmanager.CoordDelta{0, 0},
	}
	p := &Pulser{input}

	// Add the pulser to the hosts' list of inputs.
	host.inputs = append(host.inputs, p)

	return p
}
