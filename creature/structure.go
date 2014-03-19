/*
 * Creature structures.
 *
 * This covers all the structures owned by the creature object,
 */

package creature

import "github.com/DiscoViking/goBrains/locationmanager"

// The high-level creature struct.
type Creature struct {

	// The CollisionManager that this instance is managed by.
	cm locationmanager.Detection

	// The current vitality of the creature.
	// This decrements each update.  The creature dies when this reaches zero.
	vitality int
}

// An antenna belonging to a creature.
type antenna struct {

	// The host of this antenna.
	host *Creature

	// Location of the antenna on the host.
	// This is defined as length of the antenna and angle (in rads) from the host's direction.
	length float64
	arc    float64
}
