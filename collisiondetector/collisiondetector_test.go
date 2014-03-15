/*
 * Collisiondetector testing.
 */

package collisiondetector

import (
	"github.com/DiscoViking/goBrains/entity"
	"testing"
)

// Dummy entity structure for testing.
type testEntity struct {
	radius float64
}

func (te testEntity) GetRadius() float64 {
	return te.radius
}

// Test co-ordinate handling; coords and DeltaCoords.
func TestCoord(t *testing.T) {
	loc := coord{0, 0}

	// Update location and verify it.
	deltaLoc := CoordDelta{1, 2}
	loc.update(deltaLoc)

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
	hb := circleHitbox{coord{0, 0}, 10, testEntity{}}

	move := CoordDelta{1, 2}
	hb.update(move)

	if hb.centre.locX != 1 {
		t.Errorf("Expected x-location update to %v, got %v.", 1, hb.centre.locX)
	}
	if hb.centre.locY != 2 {
		t.Errorf("Expected y-location update to %v, got %v.", 2, hb.centre.locY)
	}

	// Run checks on points inside and outside the hitbox.
	hb = circleHitbox{coord{0, 0}, 10, testEntity{}}

	loc := coord{1, 2}
	if !hb.isInside(loc) {
		t.Errorf("Expected location (1, 2) to be inside hitbox.  It wasn't.")
	}

	loc = coord{12, 8}
	if hb.isInside(loc) {
		t.Errorf("Expected location (12, 8) to be outside hitbox.  It wasn't.")
	}
}

// Test basic collision detection.
func TestCollisionDetection(t *testing.T) {
	errorStr := "[%v] Expected %v hitboxes, actual: %v"

	// A test location find collisions here.
	var loc CoordDelta

	// Entities at a location.
	var col []entity.Entity

	// Set up a new collision detector.
	cm := newCollisionManager()

	// Add two entities to be managed.
	ent1 := testEntity{5}
	ent2 := testEntity{5}

	cm.addEntity(ent1)
	cm.addEntity(ent2)

	if len(cm.hitboxes) != 2 {
		t.Errorf(errorStr, 1, 2, len(cm.hitboxes))
		cm.printDebug()
	}

	// Check there are two hitboxes found at the origin.
	loc = CoordDelta{0, 0}
	col = cm.getCollisions(loc, ent1)
	if len(col) != 2 {
		t.Errorf(errorStr, 2, 2, len(col))
		cm.printDebug()
	}

	// Move a hitbox and verify it's moved.
	move := CoordDelta{10, 10}
	cm.changeLocation(move, ent2)

	col = cm.getCollisions(loc, ent1)
	if len(col) != 1 {
		t.Errorf(errorStr, 3, 1, len(col))
		cm.printDebug()
	}

	// Verify that we can detect the moved entity.
	loc = CoordDelta{10, 10}
	col = cm.getCollisions(loc, ent1)
	if len(col) != 1 {
		t.Errorf(errorStr, 4, 1, len(col))
		cm.printDebug()
	}

	// Reduce radius of the entity and verify we stop detecting it.
	loc = CoordDelta{2, 0}
	col = cm.getCollisions(loc, ent1)
	if len(col) != 1 {
		t.Errorf(errorStr, 5, 1, len(col))
		cm.printDebug()
	}

	cm.changeRadius(1, ent1)
	col = cm.getCollisions(loc, ent1)
	if len(col) != 0 {
		t.Error(errorStr, 6, 0, len(col))
		cm.printDebug()
	}
}
