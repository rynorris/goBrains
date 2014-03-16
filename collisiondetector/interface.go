/*
 * Interfaces for the collision detector.
 *
 * These interfaces provide the mechanisms by which entities in an environment can detect other entities they are in contact with.
 */

// Package collisiondetector provides all abilities to detect other entities in an environment.
package collisiondetector

import "github.com/DiscoViking/goBrains/entity"

// Detection is the set of methods that entities use to access the collision detector.
type Detection interface {

	// Methods for querying collisions, returning the collided objects.
	// -- Determine which hitboxes occlude a location relative to an entity
	GetCollisions(loc CoordDelta, ent entity.Entity) []entity.Entity

	// Methods for updating state in the detector.
	// -- Inform the detector of a change in state
	AddEntity(ent entity.Entity)
	RemoveEntity(ent entity.Entity)
	ChangeLocation(move CoordDelta, ent entity.Entity)
	ChangeRadius(radius float64, ent entity.Entity)
}

// Locatable defines the ability to calculate if you can be located.
type locatable interface {

	// Get/Set whether the the interface in use or not.
	getActive() bool
	setActive(state bool)

	// Method to check if a coordinate lies within the entity.
	isInside(loc coord) bool

	// Methods to update the properties of the hitbox.
	update(move CoordDelta)
	setRadius(radius float64)

	// Methods to check the properties of the hitbox.
	getEntity() entity.Entity
	getRadius() float64
	getCoord() coord

	// Miscellaneous.
	printDebug()
}
