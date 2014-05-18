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

	// Number of colors detected by an antenna.
	colorNum = 3

	// Max value of colours and alpha.
	colorScale = 255

	// Offset of the antenna from the face of the creature.
	arc = (math.Pi / 6)

	// Length of the antenna.
	length = 40.0

	// Charge incremented per detected thing.
	chargeAntenna = 1.0
)

// Charge an input node for a colour.
func anCharge(n *brain.Node, colorVal, alpha uint8) {

	// Each value from color.RGBA is a 255-integer.  Scale this to a float in the range 0-1.
	n.Charge(chargeAntenna * (float64(colorVal) / colorScale) * (float64(alpha) / colorScale))
}

// The antenna twitches, and attempts to detect nearby entities.
func (an *antenna) detect() {
	// Charge nodes for each colour detected.
	for _, blip := range an.host.lm.GetCollisions(an.location, an.host) {
		c := blip.Color()
		anCharge(an.colorNodes[0], c.R, c.A)
		anCharge(an.colorNodes[1], c.G, c.A)
		anCharge(an.colorNodes[2], c.B, c.A)
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

	input := inputStruct{
		putStruct: putStruct{host: host},
		node:      nil,
		location:  locationmanager.CoordDelta{length, thisArc},
	}
	a := &antenna{input, make([]*brain.Node, 0)}

	// Creatures detect colours independently.  Add a node for each colour detected.
	for ii := 0; ii < colorNum; ii++ {
		newN := brain.NewNode()
		a.colorNodes = append(a.colorNodes, newN)
		host.brain.AddInputNode(newN)
	}

	// Add the antenna to host's list of inputs.
	host.inputs = append(host.inputs, a)

	return a
}
