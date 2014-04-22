/*
 * Testentity, an entity for testing!
 */

package entity

// Get the radius of the test entity.
func (te *TestEntity) GetRadius() float64 {
	return te.Radius
}

// Check the entity.  It returns as expected.
func (te *TestEntity) Check() bool {
	return false
}

// Test entities cannot be consumed.
func (te *TestEntity) Consume() float64 {
	return 0
}
