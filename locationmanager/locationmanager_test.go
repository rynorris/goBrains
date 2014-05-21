/*
 * LocationManager testing.
 */

package locationmanager

import (
	"github.com/DiscoViking/goBrains/entity"
	"math"
	"testing"
)

// Limit movement to the default tank size.
func (c *circleHitbox) tUpdate(move CoordDelta) {
	c.update(move, coord{TANKSIZEX, TANKSIZEY})
}

// Verify that the number of hitboxes found were as expected.
func HitboxCheck(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected %v hitboxes, actual: %v",
			expected,
			actual)
	}
}

// Verify that the entries in the location manager are as expected.
func StoreCheck(t *testing.T, lm *LocationManager, entries int, actives int) {
	if (entries != len(lm.hitboxes)) || (actives != lm.NumberOwned()) {
		t.Errorf("Expected %v entries in LM (%v active), found %v entries (%v active).",
			entries,
			actives,
			len(lm.hitboxes),
			lm.NumberOwned())
	}
}

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
	hb := &circleHitbox{
		active:      true,
		centre:      coord{0, 0},
		orientation: 0,
		radius:      10,
		entity:      &entity.TestEntity{0},
	}

	move := CoordDelta{1, 0}
	hb.tUpdate(move)

	if hb.centre.locX != 1 {
		t.Errorf("Expected x-location update to %v, got %v.", 1, hb.centre.locX)
		hb.printDebug()
	}

	move = CoordDelta{0, math.Pi / 2}
	hb.tUpdate(move)

	if hb.orientation != (math.Pi / 2) {
		t.Errorf("Expected orientation update to %v, got %v.", (math.Pi / 2), hb.orientation)
		hb.printDebug()
	}

	move = CoordDelta{2, 0}
	hb.tUpdate(move)

	if hb.centre.locY != 2 {
		t.Errorf("Expected y-location update to %v, got %v.", 2, hb.centre.locY)
		hb.printDebug()
	}

	// Run checks on points inside and outside the hitbox.
	hb = &circleHitbox{
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
	lm := New()

	// The entity to query for.
	ent := &entity.TestEntity{5}

	// Query for the entity which LM does not know about.  This must fail.
	res, comb := lm.GetLocation(ent)
	locx, locy, orient := comb.X, comb.Y, comb.Orient

	if res {
		t.Errorf("Lookup of unknown object succeeded; returned: (%v, %v, %v, %v))",
			res, locx, locy, orient)
		lm.PrintDebug()
	}

	// Add the entity and query for it.
	lm.AddEntAtLocation(ent, Combination{0, 0, 0})
	res, comb = lm.GetLocation(ent)
	locx, locy, orient = comb.X, comb.Y, comb.Orient

	if !res || (locx != 0) || (locy != 0) || (orient != 0) {
		t.Errorf("Lookup of known object failed; returned: (%v, %v, %v, %v))",
			res, locx, locy, orient)
		lm.PrintDebug()
	}
}

// Test the testing function - where entities start at the origin.
func TestOrigin(t *testing.T) {

	lmn := New()
	lmo := New()
	lmo.StartAtOrigin()

	entn := &entity.TestEntity{5}
	ento := &entity.TestEntity{5}
	lmn.AddEntity(entn)
	lmo.AddEntity(ento)

	res, comb := lmo.GetLocation(ento)
	if !res || (comb != Combination{0.0, 0.0, 0.0}) {
		t.Errorf("Position of entity incorrect; returned: (%v, %v)), expected: (%v, %v)",
			res, comb, true, Combination{0.0, 0.0, 0.0})
		lmo.PrintDebug()
	}

	res, comb = lmn.GetLocation(entn)
	if !res || (comb == Combination{0.0, 0.0, 0.0}) {
		t.Errorf("Position of entity incorrect; returned: (%v, %v)),  result away from origin",
			res, comb)
		lmn.PrintDebug()
	}

}

// Test basic collision detection interface.
func TestDetection(t *testing.T) {

	// A test location find collisions here.
	var loc CoordDelta

	// Entities at a location.
	var col []entity.Entity

	// Set up a new location manager.
	cm := New()

	// Add two entities to be managed.
	ent1 := &entity.TestEntity{5}
	ent2 := &entity.TestEntity{5}

	cm.AddEntAtLocation(ent1, Combination{0.0, 0.0, 0.0})
	cm.AddEntAtLocation(ent2, Combination{0.0, 0.0, 0.0})

	HitboxCheck(t, 2, len(cm.hitboxes))

	// Check there are two hitboxes found at the origin.
	loc = CoordDelta{0, 0}
	col = cm.GetCollisions(loc, ent1)
	HitboxCheck(t, 2, len(col))

	// Move a hitbox and verify it's moved.
	move := CoordDelta{10, 0}
	cm.ChangeLocation(move, ent2)

	col = cm.GetCollisions(loc, ent1)
	HitboxCheck(t, 1, len(col))

	// Verify that we can detect the moved entity.
	loc = CoordDelta{10, 0}
	col = cm.GetCollisions(loc, ent1)
	HitboxCheck(t, 1, len(col))

	// Reduce radius of the entity at the origin and verify we stop detecting it.
	loc = CoordDelta{2, 0}
	col = cm.GetCollisions(loc, ent1)
	HitboxCheck(t, 1, len(col))

	cm.ChangeRadius(1, ent1)
	col = cm.GetCollisions(loc, ent1)
	HitboxCheck(t, 0, len(col))

	// A radius reduced to zero cannot be detected at all.
	loc = CoordDelta{0, 0}
	col = cm.GetCollisions(loc, ent1)
	HitboxCheck(t, 1, len(col))

	cm.ChangeRadius(0, ent1)
	col = cm.GetCollisions(loc, ent1)
	HitboxCheck(t, 0, len(col))
}

// Test object storage within LocationManager.
func TestStorage(t *testing.T) {
	// Set up a new location manager.
	cm := New()

	// Add two entities to be managed.
	ent1 := &entity.TestEntity{5}
	ent2 := &entity.TestEntity{5}

	cm.AddEntity(ent1)
	cm.AddEntity(ent2)

	StoreCheck(t, cm, 2, 2)

	// Remove the entities from the CM.
	// This doesn't reduce the length of the internal list, as the entries are re-used.
	cm.RemoveEntity(ent1)
	cm.RemoveEntity(ent2)
	StoreCheck(t, cm, 2, 0)

	// Add a new entry.
	// This re-uses the entries from earlier, so the list is not extended.
	cm.AddEntity(ent1)
	StoreCheck(t, cm, 2, 1)

	// Extend the list again.
	ent3 := &entity.TestEntity{5}
	cm.AddEntity(ent2)
	cm.AddEntity(ent3)
	StoreCheck(t, cm, 3, 3)
}

// Test error cases.
func TestErrors(t *testing.T) {
	lm := New()

	// Two new entities, one of which is added to LM.
	ent1 := &entity.TestEntity{5}
	ent2 := &entity.TestEntity{5}
	lm.AddEntity(ent1)

	// Successful changes.
	lm.ChangeLocation(CoordDelta{0, 0}, ent1)
	lm.ChangeRadius(0, ent1)

	// Error changes.
	lm.ChangeLocation(CoordDelta{0, 0}, ent2)
	lm.ChangeRadius(0, ent2)
}

type tankSize struct {
	x, y float64
}

// Test how far we can go in a direction.
func tankDirection(t *testing.T, s tankSize, angle float64, exp coord) {

	cm := NewLocationManager(s.x, s.y)
	ent := &entity.TestEntity{5}
	cm.AddEntAtLocation(ent, Combination{0.0, 0.0, angle})
	move := CoordDelta{TANKSIZEX * 99, 0}
	cm.ChangeLocation(move, ent)
	res, comb := cm.GetLocation(ent)
	locx, locy := comb.X, comb.Y

	if !res || (locx != exp.locX) || (locy != exp.locY) {
		t.Errorf("Expected entity to update location to maximum at (%v, %v), actually moved to (%v, %v)",
			s.x,
			s.y,
			locx,
			locy)
		cm.PrintDebug()
	}
}

// Ensure that we cannot exceed min or max range.
func tankLimit(t *testing.T, s tankSize) {

	// Move towards the maximum coordinate.
	tankDirection(t, s, (math.Pi / 4.0), coord{s.x, s.y})

	// Move away from the maximum coordinate.
	tankDirection(t, s, (math.Pi * (-3.0 / 4.0)), coord{0, 0})
}

// Test that units are held within the tank.
func TestTank(t *testing.T) {

	// Test various tank sizes.
	tankSizes := []tankSize{
		tankSize{TANKSIZEX, TANKSIZEY},      // Normal size.
		tankSize{TANKSIZEX, TANKSIZEX},      // Square.
		tankSize{TANKSIZEX, 10 * TANKSIZEY}, // Lopsided.
		tankSize{10 * TANKSIZEX, TANKSIZEY}, // Also lopsided.
		tankSize{0.0, 0.0},                  // Tiny.
	}

	for _, limit := range tankSizes {
		t.Logf("Test tank with size: (%v, %v)", limit.x, limit.y)

		// Attempt to reach max range.
		tankLimit(t, limit)
	}
}
