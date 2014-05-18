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
	"math/rand"
)

const (
	TANKSIZEX = 800.0
	TANKSIZEY = 800.0
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
		active:      true,
		centre:      coord{comb.X, comb.Y},
		orientation: comb.Orient,
		radius:      ent.Radius(),
		entity:      ent,
	}

	inserted := cm.replaceEmptyHitbox(&newHitbox)

	if !inserted {
		cm.hitboxes = append(cm.hitboxes, &newHitbox)
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
	hb.update(move, cm.maxPoint)
	if hb == nil {
		return
	}
}

// Update the radius of an entity.
func (cm *LocationManager) ChangeRadius(radius float64, ent entity.Entity) {
	hb := cm.findHitbox(ent)
	if hb == nil {
		return
	}
	hb.setRadius(radius)
}

// Determine all entities which exist at a specific point.
func (cm *LocationManager) GetCollisions(offset CoordDelta, ent entity.Entity) []entity.Entity {
	collisions := make([]entity.Entity, 0)

	searcher := cm.findHitbox(ent)
	absLoc := searcher.getCoord()

	dX := offset.Distance * math.Cos(searcher.getOrient()+offset.Rotation)
	dY := offset.Distance * math.Sin(searcher.getOrient()+offset.Rotation)
	absLoc.update(dX, dY)

	for _, hb := range cm.hitboxes {
		if hb.isInside(absLoc) {
			collisions = append(collisions, hb.getEntity())
		}
	}

	return collisions
}

// Get the location and orientation of a specific entity.
func (cm *LocationManager) GetLocation(ent entity.Entity) (bool, Combination) {
	hb := cm.findHitbox(ent)

	if hb == nil {
		return false, Combination{0, 0, 0}
	}

	coordinate := hb.getCoord()
	orientation := hb.getOrient()

	return true, Combination{coordinate.locX, coordinate.locY, orientation}
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

// Replace the first unused hitbox structure.  Return a boolean for whether the operation was successful.
func (cm *LocationManager) replaceEmptyHitbox(loc locatable) bool {
	for ii := range cm.hitboxes {
		hb := cm.hitboxes[ii]
		if !hb.getActive() {
			cm.hitboxes[ii] = loc
			return true
		}
	}
	return false
}

// Returns the number of hitboxes currently owned by the LocationManager.
func (cm *LocationManager) NumberOwned() int {
	ii := 0
	for _, hb := range cm.hitboxes {
		if hb.getActive() {
			ii++
		}
	}
	return ii
}

// Print debug information about information stored in the LocationManager.
func (cm *LocationManager) PrintDebug() {
	fmt.Printf("Location Manager: %v\n", cm)
	for ii, hb := range cm.hitboxes {
		fmt.Printf("  Hitbox %v\n", ii)
		hb.printDebug()
	}
	fmt.Printf("\n")
}

// Set LM to spawn all new entities at the origin.  For testing purposes.
func (lm *LocationManager) StartAtOrigin() {
	lm.spawnOrigin = true
}

// Initialise the LocationManager.
// Accepts the x- and y-sizes of the tank the creatures' live in.
func NewLocationManager(x, y float64) *LocationManager {
	return &LocationManager{
		spawnOrigin: false,
		hitboxes:    make([]locatable, 0),
		maxPoint:    coord{x, y},
	}
}

// Initialize a default locationmanager.
func New() *LocationManager {
	lm := NewLocationManager(TANKSIZEX, TANKSIZEY)
	return lm
}
