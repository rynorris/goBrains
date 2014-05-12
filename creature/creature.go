/*
 * The creature.
 *
 * The high-level behaviour of creatures.
 */

package creature

import (
	"github.com/DiscoViking/goBrains/brain"
	"github.com/DiscoViking/goBrains/genetics"
	"github.com/DiscoViking/goBrains/locationmanager"
)

// Fixed values.
const (
	// Maximum velocities.
	MaxLinearVel  = 5.0
	MaxAngularVel = 0.2
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
	if c.vitality <= 0 {
		c.lm.RemoveEntity(c)
		return true
	}

	// Get all our inputs to charge appropriately.
	for _, in := range c.inputs {
		in.detect()
	}

	// Update the brain one cycle.
	c.brain.Work()

	// Update LM with the distance we are moving this check.
	c.lm.ChangeLocation(locationmanager.CoordDelta{c.movement.move,
		c.movement.rotate},
		c)

	// HACK DO THIS PROPERLY LATER
	c.movement.move -= 0.01

	if c.movement.move > 1 {
		c.movement.move = 1
	} else if c.movement.move < 0 {
		c.movement.move = 0
	}
	if c.movement.rotate > 1 {
		c.movement.rotate = 1
	} else if c.movement.rotate < -1 {
		c.movement.rotate = -1
	}
	c.movement.rotate = 0
	c.vitality -= 0.1
	if c.vitality > 300 {
		c.vitality = 300
	}

	return false
}

// Breed a new creature from two existing ones.
func (cx *Creature) Breed(cy *Creature) *Creature {
	newC := NewSimple(cx.lm)
	newC.dna = cx.dna.Breed(cy.dna)
	newC.brain.Restore(newC.dna)
	return newC
}

// Clone an existing creature.
func (c *Creature) Clone() *Creature {
	newC := NewSimple(c.lm)
	newC.dna = c.dna.Clone()
	newC.brain.Restore(newC.dna)
	return newC
}

// Generates a new random DNA string for a creature and injects it into the brain.
// Must be called AFTER all outputs and inputs have been added.
func (c *Creature) Prepare() {
	n := c.brain.GenesNeeded()
	c.dna = genetics.NewDna()
	for i := 0; i < n; i++ {
		c.dna.AddGene(genetics.NewRandomGene())
	}
	c.brain.Restore(c.dna)
}

// Generate a basic creature.
func NewSimple(lm locationmanager.Detection) *Creature {
	c := New(lm)
	c.AddPulser()
	c.AddMouth()
	c.AddAntenna(AntennaLeft)
	c.AddAntenna(AntennaRight)
	c.AddBoosters()
	c.Prepare()
	return c
}

// Initialize a new creature object.
func New(lm locationmanager.Detection) *Creature {
	newC := &Creature{
		lm:       lm,
		dna:      genetics.NewDna(),
		brain:    brain.NewBrain(4),
		inputs:   make([]input, 0),
		vitality: 300,
	}

	// Add the new creature to the location manager.
	lm.AddEntity(newC)
	return newC
}
