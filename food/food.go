/*
 * Food behaviour.
 *
 * Methods for the food that creatures love best.
 */

package food

import "math"

// Check the size of the food.  This is calculated from the amount of food represented by the instance.
func (f *Food) GetRadius() float64 {
	return math.Sqrt(f.content)
}

// Initialize a new food object.
func NewFood(foodLevel float64) *Food {
	return &Food{foodLevel}
}
