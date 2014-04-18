/*
 * Methods for hitboxes.
 *
 * These methods are the behaviours of hitboxes, and allow determination of whether they interact with a point or other entity.
 */

// Package locationmanager provides all abilities to detect other entities in an environment.
package locationmanager

import (
	"fmt"
	"github.com/DiscoViking/goBrains/entity"
	"math"
)

// Get whether the hitbox is active or not.
func (hb *circleHitbox) getActive() bool {
	return hb.active
}

// Set whether the hitbox is active or not.
func (hb *circleHitbox) setActive(state bool) {
	hb.active = state
}

// Calculation of whether a co-ordinate is within a circular hitbox.
func (hb *circleHitbox) isInside(loc coord) bool {

	// A radius of zero means that a hitbox is unhittable.
	if hb.radius == 0 {
		return false
	}

	diffDist := (math.Pow((hb.centre.locX-loc.locX), 2) +
		math.Pow((hb.centre.locY-loc.locY), 2))

	if diffDist < (math.Pow(hb.radius, 2)) {
		return true
	}
	return false
}

// Update the location of a hitbox.
func (hb *circleHitbox) update(move CoordDelta) {
	hb.orientation += move.Rotation
	dX := move.Distance * math.Cos(hb.orientation)
	dY := move.Distance * math.Sin(hb.orientation)
	hb.centre.update(dX, dY)
}

// Update the radius of the hitbox.
func (hb *circleHitbox) setRadius(radius float64) {
	hb.radius = radius
}

// Get the entity owned by the hitbox.
func (hb *circleHitbox) getEntity() entity.Entity {
	return hb.entity
}

// Get the central co-ordinates of the entity.
func (hb *circleHitbox) getCoord() coord {
	return hb.centre
}

// Get the orientation of the hitbox.
func (hb *circleHitbox) getOrient() float64 {
	return hb.orientation
}

// Print debug information.
func (hb *circleHitbox) printDebug() {
	fmt.Printf("    Active: %v\n", hb.active)
	fmt.Printf("    Orient: %v\n", hb.orientation)
	fmt.Printf("    Centre: (%v, %v)\n", hb.centre.locX, hb.centre.locY)
	fmt.Printf("    Radius: %v\n", hb.radius)
	fmt.Printf("    Entity: %v\n", hb.entity)
}
