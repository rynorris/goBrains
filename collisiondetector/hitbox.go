/*
 * Methods for hitboxes.
 *
 * These methods are the behaviours of hitboxes, and allow determination of whether they interact with a point or other entity.
 */

// Package collisiondetector provides all abilities to detect other entities in an environment.
package collisiondetector

import "math"

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
func (hb circleHitbox) update(move CoordDelta) {
	hb.centre.update(move)
}

// Update the radius of the hitbox.
func (hb circleHitbox) setRadius(radius float64) {
	hb.radius = radius
}
