/*
 * The creature.
 *
 * The high-level behaviour of creatures.
 */

package creature

import (
	"image/color"

	"github.com/DiscoViking/goBrains/brain"
	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/genetics"
	"github.com/DiscoViking/goBrains/locationmanager"
)

// Fixed values.
const (
	SpeedDegredation = 0.9
)

// Creatures always report a radius of zero, as they cannot be detected.
func (c *Creature) Radius() float64 {
	return 10
}

// Get the colour of the creature.
func (c *Creature) Color() color.RGBA {
	return c.color
}

// Creatures cannot consume each other.
func (c *Creature) Consume() float64 {
	return 0
}

// Manage vitality.
func (c *Creature) manageVitality() bool {
	max := config.Global.Entity.MaxVitality
	if c.vitality <= 0 {
		c.lm.RemoveEntity(c)
		return true
	}

	// Decrement and cap vitality.
	c.vitality -= 0.07
	if c.vitality > max {
		c.vitality = max
	}

	return false
}

// Manage velocities.  Velocity must degrade, so that creatures can stop.
func (c *Creature) manageSpeed() {
	c.movement.move *= SpeedDegredation
	c.movement.rotate *= SpeedDegredation
}

// Do all updating and moving of a creature.
func (c *Creature) Work() {
	// Get all our inputs to charge appropriately.
	for _, in := range c.inputs {
		in.detect()
	}

	// Update the brain one cycle.
	c.brain.Work()
}

// Check the status of the creature and update LM appropriately.
// Returns a boolean for whether teardown occured.
func (c *Creature) Check() bool {
	death := c.manageVitality()
	if death {
		return true
	}

	// Update LM with the distance we are moving this check.
	c.lm.ChangeLocation(locationmanager.CoordDelta{c.movement.move,
		c.movement.rotate},
		c)
	c.manageSpeed()

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
		color:    color.RGBA{200, 50, 50, 255},
		vitality: config.Global.Entity.InitialVitality,
	}

	// Add the new creature to the location manager.
	lm.AddEntity(newC)
	return newC
}
