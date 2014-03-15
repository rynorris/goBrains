/*
 * Methods for hitboxes.
 *
 * These methods are the behaviours of hitboxes, and allow determination of whether they interact with a point or other entity.
 */

// Package collisiondetector provides all abilities to detect other entities in an environment.
package collisiondetector

import (
	"fmt"
	"github.com/DiscoViking/goBrains/entity"
	"math"
)

// Calculation of whether a co-ordinate is within a circular hitbox.
func (hb circleHitbox) isInside(loc coord) bool {
	diffDist := (math.Pow((hb.centre.locX-loc.locX), 2) +
		math.Pow((hb.centre.locY-loc.locY), 2))

	if diffDist < (math.Pow(hb.radius, 2)) {
		return true
	}
	return false
}

// Update the location of a hitbox.
func (hb *circleHitbox) update(move CoordDelta) {
	hb.centre.update(move)
}

// Update the radius of the hitbox.
func (hb *circleHitbox) setRadius(radius float64) {
	hb.radius = radius
}

// Get the entity owned by the hitbox.
func (hb *circleHitbox) getEntity() entity.Entity {
	return hb.entity
}

// Get the radius of the entity.
func (hb *circleHitbox) getRadius() float64 {
	return hb.radius
}

// Get the central co-ordinates of the entity.
func (hb *circleHitbox) getCoord() coord {
	return hb.centre
}

// Print debug information.
func (hb *circleHitbox) printDebug() {
	fmt.Printf("    Centre: (%v, %v)\n", hb.centre.locX, hb.centre.locY)
	fmt.Printf("    Radius: %v\n", hb.radius)
	fmt.Printf("    Entity: %v\n", hb.entity)
}
