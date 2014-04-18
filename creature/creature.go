/*
 * The creature.
 *
 * The high-level behaviour of creatures.
 */

package creature

import "github.com/DiscoViking/goBrains/locationmanager"

// Fixed values.
const (
	// Maximum velocities.
	MaxLinearVel  = 5.0
	MaxAngularVel = 0.1
)

// Creatures always report a radius of zero, as they cannot be detected.
func (c *Creature) GetRadius() float64 {
	return 0
}

// Creatures cannot consume each other.
func (c *Creature) Consume() float64 {
	return 0
}

// Check the status of the creature and update LM appropriately.
// Returns a boolean for whether teardown occured.
func (c *Creature) Check() bool {
	if c.vitality > 0 {
		return false
	}

	c.lm.RemoveEntity(c)
	return true
}

// Initialize a new creature object.
func NewCreature(lm locationmanager.Detection) *Creature {
	newC := &Creature{
		lm:       lm,
		brain:    nil,
		inputs:   make([]input, 0),
		vitality: 10,
	}

	// Add the new creature to the location manager.
	lm.AddEntity(newC)
	return newC
}
