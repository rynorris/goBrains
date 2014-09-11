/*
 * Creature interfaces.
 *
 * These interfaces are those internally and externally applicable to a creature.
 */

package creature

import (
	"github.com/DiscoViking/goBrains/entity"
)

// Interface exposed by input objects.
type input interface {

	// Activate an input, supplying environmental information to the brain.
	detect()
}

// Creature exposed methods.
type Creature interface {
	entity.Entity
	Breed(other Creature) (Creature, error)
	Clone() Creature
}
