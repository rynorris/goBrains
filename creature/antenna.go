/*
 * Antennae.
 *
 * The basic detection method for a creature.
 */

package creature

import (
	"github.com/DiscoViking/goBrains/brain"
	"github.com/DiscoViking/goBrains/locationmanager"
	"math"
)

// Fixed values.
const (
	// Antenna types.
	AntennaLeft  = 1
	AntennaRight = 2

	// Offset of the antenna from the face of the creature.
	arc = (math.Pi / 6)

	// Length of the antenna.
	length = 10.0

	// Charge incremented per detected thing.
	chargeAntenna = 1.0
)

// The antenna twitches, and attempts to detect nearby food.
func (an *antenna) detect() {
	blips := an.host.lm.GetCollisions(an.location, an.host)

	// The antenna detects all objects collided with.
	// This does not actually do any checking on what the entity is - we just detect it!
	for ii := 0; ii < len(blips); ii++ {
		an.node.Charge(chargeAntenna)
	}
}

// Add a new antenna object to a creature.
func (host *Creature) AddAntenna(antType int) *antenna {
	thisArc := 0.0
	if antType == AntennaLeft {
		thisArc = arc
	} else if antType == AntennaRight {
		thisArc = -1 * arc
	}

	// Link the antenna to the hosts' brain.
	node := brain.NewNode()
	host.brain.AddInputNode(node)

	input := inputStruct{
		putStruct: putStruct{host: host},
		node:      node,
		location:  locationmanager.CoordDelta{length, thisArc},
	}
	a := &antenna{input}

	// Add the antenna to host's list of inputs.
	host.inputs = append(host.inputs, a)

	return a
}
