/*
 * Collisiondetector testing.
 */

package collisiondetector

import "testing"

// Dummy entity structure for testing.
type testEntity struct {
	radius float64
}

// Test co-ordinate handling; coords and DeltaCoords.
func TestCoord(t *testing.T) {
	loc = coord{0, 0}

	// Update location and verify it.
	deltaLoc = CoordDelta{1, 2}
	loc.update(deltaLoc)

	if loc.locX != 1 {
		t.Errorf("Expected x-location update to %v, got %v.", loc.locX, 1)
	}
	if loc.locY != 2 {
		t.Errorf("Expected y-location update to %v, got %v.", loc.locY, 2)
	}
}

// Test the circle hitboxes.
func TestCircleHitbox(t *testing.T) {

	// Update the location of the hitbox.
	hb := circleHitbox{Coord{0, 0}, 10, testEntity{}}

	move := CoordDelta{1, 2}
	hb.update(move)

	if hb.centre.locX != 1 {
		t.Errorf("Expected x-location update to %v, got %v.", loc.locX, 1)
	}
	if hb.centre.locY != 2 {
		t.Errorf("Expected y-location update to %v, got %v.", loc.locY, 2)
	}

	// Run checks on points inside and outside the hitbox.
	hb = circleHitbox{Coord{0, 0}, 10, testEntity{}}

	loc := coord{1, 2}
	if !hb.isInside(loc) {
		t.Errorf("Expected location (1, 2) to be inside hitbox.  It wasn't.")
	}

	loc = coord{12, 8}
	if !hb.isInside(loc) {
		t.Errorf("Expected location (12, 8) to be outside hitbox.  It wasn't.")
	}
}

// Test the collision detection.
func TestCollisionDetection(t *testing.T) {
	errorStr := "[%v] Expected %v hitboxes to be set up, actual: %v"

	// Set up a new collision detector.
	cm := newCollisionManager()

	// Add two entities to be managed.
	ent1 := testEntity{5}
	ent2 := testEntity{5}

	cm.addEntity(ent1)
	cm.addEntity(ent2)

	if len(cm.hitboxes) != 2 {
		t.Errorf(errorStr, 1, 2, len(cm.hitboxes))
	}

	// Check there are two hitboxes found at the origin.
	loc := CoordDelta{0, 0}
	num := cm.getCollisions(loc, ent1)
	if num != 2 {
		t.Errorf(errorStr, 2, 2, num)
	}

	// Move a hitbox and verify it's moved.
	move := CoordDelta{10, 10}
	cm.changeLocation(move, ent2)

	num = cm.getCollisions(loc, ent1)
	if num != 1 {
		t.Errorf(errorStr, 3, 1, num)
	}

	// Verify that we can detect the moved entity.
	loc = CoordDelta{10, 10}
	num = cm.getCollisions(loc, ent1)
	if num != 1 {
		t.Errorf(errorStr, 4, 1, num)
	}

	// Reduce radius of the entity and verify we stop detecting it.
	loc = CoordDelta{2, 0}
	num = cm.getCollisions(loc, ent1)
	if num != 1 {
		t.Errorf(errorStr, 5, 1, num)
	}

	cm.changeRadius(1, ent1)
	num = cm.getCollisions(loc, ent1)
	if num != 0 {
		t.Error(errorStr, 6, 0, num)
	}
}
