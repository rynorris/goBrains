/*
 * Creature structures.
 *
 * This covers all the structures owned by the creature object,
 */

package creature

import (
	"github.com/DiscoViking/goBrains/brain"
	"github.com/DiscoViking/goBrains/locationmanager"
)

// The high-level creature struct.
type Creature struct {

	// The CollisionManager that this instance is managed by.
	lm locationmanager.Detection

	// The nervous system of this creature.  Hopefully doing something intelligent.
	brain brain.AcceptInput

	// Brain inputs!
	inputs []input

	// The current vitality of the creature.
	// This decrements each update.  The creature dies when this reaches zero.
	vitality float64
}

// A generic input structure for an input belonging to a creature.
type inputStruct struct {

	// The host of this antenna.
	host *Creature

	// Input node that this input charges in the brain.
	node *brain.Node

	// Location of the input on the host.
	location locationmanager.CoordDelta
}

// Basic creature detection.
type antenna struct {
	inputStruct
}

// Basic creature food consumption.
type mouth struct {
	inputStruct
}
