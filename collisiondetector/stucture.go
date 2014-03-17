/*
 * Structures for the collisiondetector.
 *
 * These structures track properties of objects managed by the collisiondetector.
 */

// Package collisiondetector provides all abilities to detect other entities in an environment.
package collisiondetector

import "github.com/DiscoViking/goBrains/entity"

// Coord structs hold a specific co-ordinate (point) within the environment.
type coord struct {
	locX, locY float64
}

// CircleHitbox represents a circular hitbox (handy, right?)
// They have two active values, representing the centre of the hitbox, and the radius.  It also references the entity it represents.
type circleHitbox struct {

	// Active values, holding state.
	centre coord
	radius float64

	// External reference, to the entity that the hitbox represents.
	entity entity.Entity
}

// CoordDelta structs represent a position relative to an entity.
type CoordDelta struct {
	deltaX, deltaY float64
}

// CollisionManager is an instance of a collision manager.
// It holds all the state about entities in the environment.
type CollisionManager struct {
	hitboxes []locatable
}
