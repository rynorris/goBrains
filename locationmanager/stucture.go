/*
 * Structures for the locationmanager.
 *
 * These structures track properties of objects managed by the locationmanager.
 */

// Package locationmanager provides all abilities to detect other entities in an environment.
package locationmanager

import "github.com/DiscoViking/goBrains/entity"

// Coord structs hold a specific co-ordinate (point) within the environment.
type coord struct {
	locX, locY float64
}

// CircleHitbox represents a circular hitbox (handy, right?)
// They have two active values, representing the centre of the hitbox, and the radius.  It also references the entity it represents.
type circleHitbox struct {

	// Active values, holding state.
	centre      coord
	orientation float64
	radius      float64

	// External reference, to the entity that the hitbox represents.
	entity entity.Entity

	// List of spacial zones this hitbox exists in.
	mZones []*spacialZone
}

// CoordDelta structs represent a position relative to an entity.
type CoordDelta struct {
	Distance float64
	Rotation float64
}

// Combination represents a combination position of x, y and an orientation.
type Combination struct {
	X, Y, Orient float64
}

// LocationManager is an instance of a location manager.
// It holds all the state about entities in the environment.
type LocationManager struct {

	// Special spawning behaviour.
	spawnOrigin bool

	// The hitboxes know by LM.
	hitboxes map[entity.Entity]locatable

	// The maximum point that objects can exist at in LM.
	maxPoint coord

	// The 2-D array of buckets used for collisions.
	spacialZones [][]spacialZone
}
