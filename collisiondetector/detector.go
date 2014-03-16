/*
 * Collision detection.
 *
 * The core methods behind detecting collisions.
 */

// Package collisiondetector provides all abilities to detect other entities in an environment.
package collisiondetector

import (
	"fmt"
	"github.com/DiscoViking/goBrains/entity"
)

// Add a new entity.
// This is added to first empty entry in the array, else append a new entry.
func (cm *CollisionManager) AddEntity(ent entity.Entity) {
	newHitbox := circleHitbox{
		active: true,
		centre: coord{0, 0},
		radius: ent.GetRadius(),
		entity: ent,
	}

	entry := cm.findEmptyHitbox()
	if entry == nil {
		cm.hitboxes = append(cm.hitboxes, &newHitbox)
	} else {
		entry = &newHitbox
	}
}

// Remove an entity.
func (cm *CollisionManager) RemoveEntity(ent entity.Entity) {
	hb := cm.findHitbox(ent)
	hb.setActive(false)
}

// Update the location of an entity.
func (cm *CollisionManager) ChangeLocation(move CoordDelta, ent entity.Entity) {
	hb := cm.findHitbox(ent)
	hb.update(move)
}

// Update the radius of an entity.
func (cm *CollisionManager) ChangeRadius(radius float64, ent entity.Entity) {
	hb := cm.findHitbox(ent)
	hb.setRadius(radius)
}

// Determine all entities which exist at a specific point.
func (cm *CollisionManager) GetCollisions(offset CoordDelta, ent entity.Entity) []entity.Entity {
	collisions := make([]entity.Entity, 0)

	searcher := cm.findHitbox(ent)
	absLoc := searcher.getCoord()
	absLoc.update(offset)

	for _, hb := range cm.hitboxes {
		if hb.isInside(absLoc) {
			collisions = append(collisions, hb.getEntity())
		}
	}

	return collisions
}

// Find the hitbox associated with an entity.
func (cm *CollisionManager) findHitbox(ent entity.Entity) locatable {
	for _, hb := range cm.hitboxes {
		if hb.getActive() && (hb.getEntity() == ent) {
			return hb
		}
	}
	return nil
}

// Find the first unused ehitbox structure.
func (cm *CollisionManager) findEmptyHitbox() locatable {
	for _, hb := range cm.hitboxes {
		if !hb.getActive() {
			return hb
		}
	}
	return nil
}

// Print debug information about information stored in the CollisionManager.
func (cm *CollisionManager) printDebug() {
	fmt.Printf("Collision Manager: %v\n", cm)
	for ii, hb := range cm.hitboxes {
		fmt.Printf("  Hitbox %v\n", ii)
		hb.printDebug()
	}
	fmt.Printf("\n")
}

// Initialise the CollisionManager.
func NewCollisionManager() *CollisionManager {
	return &CollisionManager{hitboxes: make([]locatable, 0)}
}
