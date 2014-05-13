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
	chargePulser = 0.03
)

// The pulser fires indiscriminately.
func (p *pulser) detect() {
	p.node.Charge(chargePulser)
}

// Add a new pulser to a creature.
func (host *Creature) AddPulser() *pulser {

	// Link the pulser to the host's brain.
	node := brain.NewNode()
	host.brain.AddInputNode(node)

	input := inputStruct{
		putStruct: putStruct{host: host},
		node:      node,
		location:  locationmanager.CoordDelta{0, 0},
	}
	p := &pulser{input}

	// Add the pulser to the hosts' list of inputs.
	host.inputs = append(host.inputs, p)

	return p
}
