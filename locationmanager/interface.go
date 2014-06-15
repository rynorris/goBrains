/*
 * Interfaces for the location manager.
 *
 * These interfaces provide the mechanisms by which entities in an environment can detect other entities they are in contact with.
 */

// Package locationmanager provides all abilities to detect other entities in an environment.
package locationmanager

import "github.com/DiscoViking/goBrains/entity"

// Detection is the set of methods that entities use to access the collision detector.
type Detection interface {

	// Methods for querying collisions, returning the collided objects.
	// -- Determine which hitboxes occlude a location relative to an entity
	GetCollisions(loc CoordDelta, ent entity.Entity) []entity.Entity

	// Methods for updating state in the detector.
	// -- Inform the detector of a change in state
	AddEntity(ent entity.Entity)
	AddEntAtLocation(ent entity.Entity, comb Combination)
	RemoveEntity(ent entity.Entity)
	ChangeLocation(move CoordDelta, ent entity.Entity)
	ChangeRadius(radius float64, ent entity.Entity)
}

// Location is the set of methods that the graphics package uses to find the location of entities.
type Location interface {

	// Retrieve location of an entity.
	// This returns the x- and y-coordinates, as well as the orientation.
	// The first value is a boolean for whether the lookup was successful.
	// This will fail if the queried entity has not registered with the location manager.
	GetLocation(ent entity.Entity) (bool, Combination)
}

type LM interface {
	Detection
	Location
}

// Locatable defines the ability to calculate if you can be located.
type locatable interface {

	// Method to check if a coordinate lies within the entity.
	isInside(loc coord) bool

	// Methods to update the properties of the hitbox.
	update(move CoordDelta, max coord)
	setRadius(radius float64)

	// Methods to check the properties of the hitbox.
	getEntity() entity.Entity
	getCoord() coord
	getOrient() float64

	// Method to get the 4 corners of a box bounding this hitbox.
	boundingBox() []coord
	zones() []*spacialZone
	addZone(z *spacialZone)
	clearZones()
}
