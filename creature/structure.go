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
	brain brain.BrainExternal

	// Brain inputs!
	inputs []input

	// The current vitality of the creature.
	// This decrements each update.  The creature dies when this reaches zero.
	vitality float64

	// Movement information.
	movement velocity
}

// Veloocity information stored by a creature.
type velocity struct {

	// Current linear velocity of the creature.
	move float64

	// Current angular velocity of the creature.
	rotate float64
}

// A generic input/output structure for an input or output belonging to a creature.
type putStruct struct {

	// The host of this antenna.
	host *Creature
}

// Generic input structure.
type inputStruct struct {

	// Generic input/output structure.
	putStruct

	// Input node that this input charges in the brain.
	node *brain.Node

	// Location of the input on the host.
	location locationmanager.CoordDelta
}

// Generic output structure.
type outputStruct struct {

	// Generic input/output structure.
	putStruct

	// Current charge.
	charge float64
}

type antenna struct {
	inputStruct
}

type mouth struct {
	inputStruct
}

type booster struct {
	putStruct

	// Charge held by the booster.
	charge float64

	// Type of booster.
	btype int
}
