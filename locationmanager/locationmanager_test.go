/*
 * LocationManager testing.
 */

package locationmanager

import (
	"github.com/DiscoViking/goBrains/entity"
	"math"
	"testing"
)

// Test co-ordinate handling; coords and DeltaCoords.
func TestCoord(t *testing.T) {
	loc := coord{0, 0}

	// Update location and verify it.
	loc.update(1, 2)

	if loc.locX != 1 {
		t.Errorf("Expected x-location update to %v, got %v.", 1, loc.locX)
	}
	if loc.locY != 2 {
		t.Errorf("Expected y-location update to %v, got %v.", 2, loc.locY)
	}
}

// Test the circle hitboxes.
func TestCircleHitbox(t *testing.T) {

	// Update the location of the hitbox.
	hb := circleHitbox{
		active:      true,
		centre:      coord{0, 0},
		orientation: 0,
		radius:      10,
		entity:      &entity.TestEntity{0},
	}

	move := CoordDelta{1, 0}
	hb.update(move)

	if hb.centre.locX != 1 {
		t.Errorf("Expected x-location update to %v, got %v.", 1, hb.centre.locX)
		hb.printDebug()
	}

	move = CoordDelta{0, math.Pi / 2}
	hb.update(move)

	if hb.orientation != (math.Pi / 2) {
		t.Errorf("Expected orientation update to %v, got %v.", (math.Pi / 2), hb.orientation)
		hb.printDebug()
	}

	move = CoordDelta{2, 0}
	hb.update(move)

	if hb.centre.locY != 2 {
		t.Errorf("Expected y-location update to %v, got %v.", 2, hb.centre.locY)
		hb.printDebug()
	}

	// Run checks on points inside and outside the hitbox.
	hb = circleHitbox{
		active:      true,
		centre:      coord{0, 0},
		orientation: (math.Pi * 2 / 6),
		radius:      10,
		entity:      &entity.TestEntity{0},
	}

	loc := coord{1, 2}
	if !hb.isInside(loc) {
		t.Errorf("Expected location (1, 2) to be inside hitbox.  It wasn't.")
	}

	loc = coord{12, 8}
	if hb.isInside(loc) {
		t.Errorf("Expected location (12, 8) to be outside hitbox.  It wasn't.")
	}
}

// Test the location interface.
func TestLocation(t *testing.T) {

	// Set up a new location manager.
	lm := NewLocationManager()

	// The entity to query for.
	ent := &entity.TestEntity{5}

	// Query for the entity which LM does not know about.  This must fail.
	res, locx, locy, orient := lm.GetLocation(ent)

	if res {
		t.Errorf("Lookup of unknown object succeeded; returned: (%v, %v, %v, %v))",
			res, locx, locy, orient)
		lm.printDebug()
	}

	// Add the entity and query for it.
	lm.AddEntity(ent)
	res, locx, locy, orient = lm.GetLocation(ent)

	if !res || (locx != 0) || (locy != 0) || (orient != 0) {
		t.Errorf("Lookup of known object failed; returned:x (%v, %v, %v, %v))",
			res, locx, locy, orient)
		lm.printDebug()
	}
}

// Test basic collision detection interface.
func TestDetection(t *testing.T) {
	errorStr := "[%v] Expected %v hitboxes, actual: %v"

	// A test location find collisions here.
	var loc CoordDelta

	// Entities at a location.
	var col []entity.Entity

	// Set up a new location manager.
	cm := NewLocationManager()

	// Add two entities to be managed.
	ent1 := &entity.TestEntity{5}
	ent2 := &entity.TestEntity{5}

	cm.AddEntity(ent1)
	cm.AddEntity(ent2)

	if len(cm.hitboxes) != 2 {
		t.Errorf(errorStr, 1, 2, len(cm.hitboxes))
		cm.printDebug()
	}

	// Check there are two hitboxes found at the origin.
	loc = CoordDelta{0, 0}
	col = cm.GetCollisions(loc, ent1)
	if len(col) != 2 {
		t.Errorf(errorStr, 2, 2, len(col))
		cm.printDebug()
	}

	// Move a hitbox and verify it's moved.
	move := CoordDelta{10, 0}
	cm.ChangeLocation(move, ent2)

	col = cm.GetCollisions(loc, ent1)
	if len(col) != 1 {
		t.Errorf(errorStr, 3, 1, len(col))
		cm.printDebug()
	}

	// Verify that we can detect the moved entity.
	loc = CoordDelta{10, 0}
	col = cm.GetCollisions(loc, ent1)
	if len(col) != 1 {
		t.Errorf(errorStr, 4, 1, len(col))
		cm.printDebug()
	}

	// Reduce radius of the entity at the origin and verify we stop detecting it.
	loc = CoordDelta{2, 0}
	col = cm.GetCollisions(loc, ent1)
	if len(col) != 1 {
		t.Errorf(errorStr, 5, 1, len(col))
		cm.printDebug()
	}

	cm.ChangeRadius(1, ent1)
	col = cm.GetCollisions(loc, ent1)
	if len(col) != 0 {
		t.Error(errorStr, 6, 0, len(col))
		cm.printDebug()
	}

	// A radius reduced to zero cannot be detected at all.
	loc = CoordDelta{0, 0}
	col = cm.GetCollisions(loc, ent1)
	if len(col) != 1 {
		t.Errorf(errorStr, 5, 1, len(col))
		cm.printDebug()
	}

	cm.ChangeRadius(0, ent1)
	col = cm.GetCollisions(loc, ent1)
	if len(col) != 0 {
		t.Error(errorStr, 6, 0, len(col))
		cm.printDebug()
	}

	// Remove the entities from the CM.
	// This doesn't reduce the length of the internal list, as the entries are re-used.
	cm.RemoveEntity(ent1)
	cm.RemoveEntity(ent2)
	if len(cm.hitboxes) != 2 {
		t.Errorf(errorStr, 7, 2, len(cm.hitboxes))
		cm.printDebug()
	}

	// Add a new entry.
	// This re-uses the entries from earlier, so the list is not extended.
	ent3 := &entity.TestEntity{5}
	cm.AddEntity(ent3)
	if len(cm.hitboxes) != 2 {
		t.Errorf(errorStr, 8, 2, len(cm.hitboxes))
		cm.printDebug()
	}
}
