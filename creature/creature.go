/*
 * The creature.
 *
 * The high-level behaviour of creatures.
 */

package creature

import "github.com/DiscoViking/goBrains/locationmanager"

// Creatures always report a radius of zero, as they cannot be detected.
func (c *Creature) GetRadius() float64 {
	return 0
}

// Attempt to tear down a creature.
// Call this at the end of each cycle, to remove it from the collision manager.
// Returns a boolean for whether the teardown occured.
func (c *Creature) Check() bool {
	if c.vitality > 0 {
		return false
	}

	c.cm.RemoveEntity(c)
	return true
}

// Initialize a new creature object.
func New(cm locationmanager.Detection) *Creature {
	newC := &Creature{
		cm:       cm,
		vitality: 10,
	}

	// Add the new creature to the location manager.
	cm.AddEntity(newC)
	return newC
}
