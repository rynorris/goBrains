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
	vitality int
}

// An antenna belonging to a creature.
type antenna struct {

	// The host of this antenna.
	host *Creature

	// Input node that this input charges in the brain.
	node *brain.Node

	// Location of the antenna on the host.
	// This is defined as length of the antenna and angle (in rads) from the host's direction.
	location locationmanager.CoordDelta
}
