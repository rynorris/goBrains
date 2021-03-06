/*
 * Location management.
 *
 * The core methods behind detecting collisions.
 */

// Package locationmanager provides all abilities to detect other entities in an environment.
package locationmanager

import (
	"math"
	"math/rand"

	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/entity"
)

// Add a new entity to a random position in the tank.
// This is added to first empty entry in the array, else append a new entry.
func (cm *LocationManager) AddEntity(ent entity.Entity) {
	comb := Combination{0.0, 0.0, 0.0}
	if !cm.spawnOrigin {
		comb = Combination{
			X:      (rand.Float64() * cm.maxPoint.locX),
			Y:      (rand.Float64() * cm.maxPoint.locY),
			Orient: (rand.Float64() * 2 * math.Pi),
		}
	}
	cm.AddEntAtLocation(ent, comb)
}

// Add an entity at a specific position and orientation.
func (cm *LocationManager) AddEntAtLocation(ent entity.Entity, comb Combination) {
	newHitbox := circleHitbox{
		centre:      coord{comb.X, comb.Y},
		orientation: comb.Orient,
		radius:      ent.Radius(),
		entity:      ent,
	}
	// Allocate memory for the zone store and bounding box.
	newHitbox.mZones = make([]*spacialZone, 0, 4)
	newHitbox.bb = make([]coord, 0, 4)

	cm.hitboxes[ent] = &newHitbox

	// Add the hitbox to all zones it belongs to.
	cm.addToZones(&newHitbox)
}

// Remove an entity.
func (cm *LocationManager) RemoveEntity(ent entity.Entity) {
	cm.removeFromZones(cm.hitboxes[ent])
	delete(cm.hitboxes, ent)
}

// Update the location of an entity.
func (cm *LocationManager) ChangeLocation(move CoordDelta, ent entity.Entity) {
	hb := cm.hitboxes[ent]
	if hb == nil {
		return
	}
	cm.removeFromZones(hb)
	hb.update(move, cm.maxPoint)
	cm.addToZones(hb)
}

// Update the radius of an entity.
func (cm *LocationManager) ChangeRadius(radius float64, ent entity.Entity) {
	hb := cm.hitboxes[ent]
	if hb == nil {
		return
	}
	hb.setRadius(radius)
}

// Determine all entities which exist at a specific point.
func (cm *LocationManager) GetCollisions(offset CoordDelta, ent entity.Entity) []entity.Entity {
	collisions := make([]entity.Entity, 0)

	searcher := cm.hitboxes[ent]
	absLoc := searcher.getCoord()

	dX := offset.Distance * math.Cos(searcher.getOrient()+offset.Rotation)
	dY := offset.Distance * math.Sin(searcher.getOrient()+offset.Rotation)
	absLoc.update(dX, dY)

	// Only need to check entities in the same zone.
	zone := cm.findZone(absLoc)

	for _, hb := range *zone {
		if hb.isInside(absLoc) {
			collisions = append(collisions, hb.getEntity())
		}
	}

	return collisions
}

// Get the location and orientation of a specific entity.
func (cm *LocationManager) GetLocation(ent entity.Entity) (bool, Combination) {
	hb := cm.hitboxes[ent]

	if hb == nil {
		return false, Combination{0, 0, 0}
	}

	coordinate := hb.getCoord()
	orientation := hb.getOrient()

	return true, Combination{coordinate.locX, coordinate.locY, orientation}
}

// Returns the number of hitboxes currently owned by the LocationManager.
func (cm *LocationManager) NumberOwned() int {
	return len(cm.hitboxes)
}

// Set LM to spawn all new entities at the origin.  For testing purposes.
func (lm *LocationManager) StartAtOrigin() {
	lm.spawnOrigin = true
}

// Initialise the LocationManager.
// Accepts the x- and y-sizes of the tank the creatures' live in.
func NewLocationManager(x, y float64) *LocationManager {
	lm := &LocationManager{
		spawnOrigin: false,
		hitboxes:    make(map[entity.Entity]locatable),
		maxPoint:    coord{x, y},
	}
	lm.resetZones()
	return lm
}

// Initialize a default locationmanager.
func New() *LocationManager {
	lm := NewLocationManager(
		float64(config.Global.General.ScreenWidth),
		float64(config.Global.General.ScreenHeight))
	return lm
}
