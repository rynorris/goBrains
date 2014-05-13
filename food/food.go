/*
 * Food behaviour.
 *
 * Methods for the food that creatures love best.
 */

package food

import (
	"math"

	"github.com/DiscoViking/goBrains/locationmanager"
)

const (
	decay_rate = 0.1 // Amount of food that decays each tick.
)

// Check the size of the food.  This is calculated from the amount of food represented by the instance.
func (f *Food) GetRadius() float64 {
	return math.Sqrt(f.content)
}

// Consumption of the food.  Returns the food content eaten.
func (f *Food) Consume() float64 {
	initFood := f.content

	// Each bite removes 1 from the food instance.  We cannot have a negative content, however.
	f.content--
	if f.content < 0 {
		f.content = 0
	}

	// Report the new radius to the collision detector.
	f.cm.ChangeRadius(f.GetRadius(), f)

	return (initFood - f.content)
}

// Attempt to tear down a food object.
// Call this at the end of each cycle, to remove it from the collision manager.
// Returns a boolean for whether the teardown occured.
func (f *Food) Check() bool {
	f.content -= decay_rate
	if f.content > 0 {
		return false
	}
	f.content = 0

	f.cm.RemoveEntity(f)
	return true
}

// Content checking.  This is provided for testing purposes.
func (f *Food) GetContent() float64 {
	return f.content
}

// Initialize a new food object.
func New(cm locationmanager.Detection, foodLevel float64) *Food {
	newF := &Food{
		content: foodLevel,
		cm:      cm,
	}

	// Add the new food to the collision detector.
	cm.AddEntity(newF)
	return newF
}
