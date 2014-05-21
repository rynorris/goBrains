/*
 * Testentity, an entity for testing!
 */

package entity

import "image/color"

// Get the radius of the test entity.
func (te *TestEntity) Radius() float64 {
	return te.TeRadius
}

// Get the colour of this entity.
func (te *TestEntity) Color() color.RGBA {
	return color.RGBA{255, 255, 255, 255}
}

// Test entity does no work.
func (te *TestEntity) Work() {}

// Check the entity.  It returns as expected.
func (te *TestEntity) Check() bool {
	return false
}

// Test entities cannot be consumed.
func (te *TestEntity) Consume() float64 {
	return 0
}
