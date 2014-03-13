/*
 * Collision detection.
 *
 * The core methods behind detecting collisions.
 */

// Package collisiondetector provides all abilities to detect other entities in an environment.
package collisiondetector

import "github.com/DiscoViking/goBrains/entity"

// Add a new entity.
func (cm CollisionManager) addEntity(ent entity.Entity) {
	newHitbox := circleHitbox{
		centre: coord{0, 0},
		radius: ent.GetRadius(),
		entity: ent,
	}

	cm.hitboxes = append(cm.hitboxes, &newHitbox)
}

// Update the location of an entity.
func (cm CollisionManager) changeLocation(move CoordDelta, ent entity.Entity) {
	hb := cm.findHitbox(ent)
	hb.update(move)
}

// Update the radius of an entity.
func (cm CollisionManager) changeRadius(radius float64, ent entity.Entity) {
	hb := cm.findHitbox(ent)
	hb.setRadius(radius)
}

// Determine all entities which exist at a specific point.
func (cm CollisionManager) getCollisions(offset CoordDelta, ent entity.Entity) []entity.Entity {
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
func (cm CollisionManager) findHitbox(ent entity.Entity) locatable {
	for _, hb := range cm.hitboxes {
		if hb.getEntity() == ent {
			return hb
		}
	}
	return nil
}

// Initialise the CollisionManager.
func newCollisionManager() *CollisionManager {
	return &CollisionManager{hitboxes: make([]locatable, 0)}
}
