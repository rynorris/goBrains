/*
 * Methods for hitboxes.
 *
 * These methods are the behaviours of hitboxes, and allow determination of whether they interact with a point or other entity.
 */

// Package locationmanager provides all abilities to detect other entities in an environment.
package locationmanager

import (
	"math"

	"github.com/DiscoViking/goBrains/entity"
)

// Calculation of whether a co-ordinate is within a circular hitbox.
func (hb *circleHitbox) isInside(loc coord) bool {

	// A radius of zero means that a hitbox is unhittable.
	if hb.radius == 0 {
		return false
	}

	dX := hb.centre.locX - loc.locX
	dY := hb.centre.locY - loc.locY
	diffDist := dX*dX + dY*dY

	if diffDist < hb.radius*hb.radius {
		return true
	}
	return false
}

// Update the location of a hitbox.
func (hb *circleHitbox) update(move CoordDelta, max coord) {
	hb.orientation += move.Rotation
	dX := move.Distance * math.Cos(hb.orientation)
	dY := move.Distance * math.Sin(hb.orientation)
	hb.centre.update(dX, dY)
	hb.centre.limit(&coord{0, 0}, &max)
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

// Get the corners of a box bounding this hitbox.
func (hb *circleHitbox) boundingBox() []coord {
	hb.bb = hb.bb[:0]
	c := coord{hb.centre.locX + hb.radius, hb.centre.locX + hb.radius}
	hb.bb = append(hb.bb, c)
	c = coord{hb.centre.locX + hb.radius, hb.centre.locX - hb.radius}
	hb.bb = append(hb.bb, c)
	c = coord{hb.centre.locX - hb.radius, hb.centre.locX + hb.radius}
	hb.bb = append(hb.bb, c)
	c = coord{hb.centre.locX - hb.radius, hb.centre.locX - hb.radius}
	hb.bb = append(hb.bb, c)

	return hb.bb
}

// zones returns a slice of the spacial zones this hitbox belongs to.
func (hb *circleHitbox) zones() []*spacialZone {
	return hb.mZones
}

// clearZones empties the zone store.
func (hb *circleHitbox) clearZones() {
	hb.mZones = hb.mZones[:0]
}

// addZone adds a zone to the zone store.
func (hb *circleHitbox) addZone(z *spacialZone) {
	hb.mZones = append(hb.mZones, z)
}
