/*
 * Location management.
 *
 * The core methods behind detecting collisions.
 */

// Package locationmanager provides all abilities to detect other entities in an environment.
package locationmanager

import (
	"fmt"
	"github.com/DiscoViking/goBrains/entity"
	"math"
)

// Add a new entity.
// This is added to first empty entry in the array, else append a new entry.
func (cm *LocationManager) AddEntity(ent entity.Entity) {
	newHitbox := circleHitbox{
		active:      true,
		centre:      coord{0, 0},
		orientation: 0,
		radius:      ent.GetRadius(),
		entity:      ent,
	}

	entry := cm.findEmptyHitbox()
	if entry == nil {
		cm.hitboxes = append(cm.hitboxes, &newHitbox)
	} else {
		entry = &newHitbox
	}
}

// Remove an entity.
func (cm *LocationManager) RemoveEntity(ent entity.Entity) {
	hb := cm.findHitbox(ent)
	hb.setActive(false)
}

// Update the location of an entity.
func (cm *LocationManager) ChangeLocation(move CoordDelta, ent entity.Entity) {
	hb := cm.findHitbox(ent)
	hb.update(move)
}

// Update the radius of an entity.
func (cm *LocationManager) ChangeRadius(radius float64, ent entity.Entity) {
	hb := cm.findHitbox(ent)
	hb.setRadius(radius)
}

// Determine all entities which exist at a specific point.
func (cm *LocationManager) GetCollisions(offset CoordDelta, ent entity.Entity) []entity.Entity {
	collisions := make([]entity.Entity, 0)

	searcher := cm.findHitbox(ent)
	absLoc := searcher.getCoord()

	dX := offset.Distance * math.Cos(searcher.getOrient())
	dY := offset.Distance * math.Sin(searcher.getOrient())
	absLoc.update(dX, dY)

	for _, hb := range cm.hitboxes {
		if hb.isInside(absLoc) {
			collisions = append(collisions, hb.getEntity())
		}
	}

	return collisions
}

// Get the location and orientation of a specific entity.
func (cm *LocationManager) GetLocation(ent entity.Entity) (bool, float64, float64, float64) {
	hb := cm.findHitbox(ent)

	if hb == nil {
		return false, 0, 0, 0
	}

	coordinate := hb.getCoord()
	orientation := hb.getOrient()

	return true, coordinate.locX, coordinate.locY, orientation
}

// Find the hitbox associated with an entity.
func (cm *LocationManager) findHitbox(ent entity.Entity) locatable {
	for _, hb := range cm.hitboxes {
		if hb.getActive() && (hb.getEntity() == ent) {
			return hb
		}
	}
	return nil
}

// Find the first unused hitbox structure.
func (cm *LocationManager) findEmptyHitbox() locatable {
	for _, hb := range cm.hitboxes {
		if !hb.getActive() {
			return hb
		}
	}
	return nil
}

// Print debug information about information stored in the LocationManager.
func (cm *LocationManager) printDebug() {
	fmt.Printf("Location Manager: %v\n", cm)
	for ii, hb := range cm.hitboxes {
		fmt.Printf("  Hitbox %v\n", ii)
		hb.printDebug()
	}
	fmt.Printf("\n")
}

// Initialise the LocationManager.
func NewLocationManager() *LocationManager {
	return &LocationManager{hitboxes: make([]locatable, 0)}
}
