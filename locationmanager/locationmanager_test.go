/*
 * LocationManager testing.
 */

package locationmanager

import (
	"fmt"
	"math"
	"testing"

	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/entity"
)

const (
	BASE_WIDTH  = 800
	BASE_HEIGHT = 800
)

// Limit movement to the default tank size.
func (c *circleHitbox) tUpdate(move CoordDelta) {
	c.update(move, coord{BASE_WIDTH, BASE_HEIGHT})
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
func StoreCheck(t *testing.T, lm *LocationManager, entries int) {
	if entries != lm.NumberOwned() {
		t.Errorf("Expected %v entries in LM, found %v entries.",
			entries,
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
		centre:      coord{0, 0},
		orientation: 0,
		radius:      10,
		entity:      &entity.TestEntity{TeRadius: 0},
	}

	move := CoordDelta{1, 0}
	hb.tUpdate(move)

	if hb.centre.locX != 1 {
		t.Errorf("Expected x-location update to %v, got %v.", 1, hb.centre.locX)
		fmt.Printf("  Hitbox: %v\n", hb)
	}

	move = CoordDelta{0, math.Pi / 2}
	hb.tUpdate(move)

	if hb.orientation != (math.Pi / 2) {
		t.Errorf("Expected orientation update to %v, got %v.", (math.Pi / 2), hb.orientation)
		fmt.Printf("  Hitbox: %v\n", hb)
	}

	move = CoordDelta{2, 0}
	hb.tUpdate(move)

	if hb.centre.locY != 2 {
		t.Errorf("Expected y-location update to %v, got %v.", 2, hb.centre.locY)
		fmt.Printf("  Hitbox: %v\n", hb)
	}

	// Run checks on points inside and outside the hitbox.
	hb = &circleHitbox{
		centre:      coord{0, 0},
		orientation: (math.Pi * 2 / 6),
		radius:      10,
		entity:      &entity.TestEntity{TeRadius: 0},
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
	lm := NewLocationManager(BASE_WIDTH, BASE_HEIGHT)

	// The entity to query for.
	ent := &entity.TestEntity{TeRadius: 5}

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

	lmn := NewLocationManager(BASE_WIDTH, BASE_HEIGHT)
	lmo := NewLocationManager(BASE_WIDTH, BASE_HEIGHT)
	lmo.StartAtOrigin()

	entn := &entity.TestEntity{TeRadius: 5}
	ento := &entity.TestEntity{TeRadius: 5}
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
	cm := NewLocationManager(BASE_WIDTH, BASE_HEIGHT)

	// Add two entities to be managed.
	ent1 := &entity.TestEntity{TeRadius: 5}
	ent2 := &entity.TestEntity{TeRadius: 5}

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
	cm := NewLocationManager(BASE_WIDTH, BASE_HEIGHT)

	// Add two entities to be managed.
	ent1 := &entity.TestEntity{TeRadius: 5}
	ent2 := &entity.TestEntity{TeRadius: 5}

	cm.AddEntity(ent1)
	cm.AddEntity(ent2)

	StoreCheck(t, cm, 2)

	// Remove the entities from the CM.
	cm.RemoveEntity(ent1)
	cm.RemoveEntity(ent2)
	StoreCheck(t, cm, 0)

	// Add a new entry.
	cm.AddEntity(ent1)
	StoreCheck(t, cm, 1)

	// Extend the list again.
	ent3 := &entity.TestEntity{TeRadius: 5}
	cm.AddEntity(ent2)
	cm.AddEntity(ent3)
	StoreCheck(t, cm, 3)
}

// Test error cases.
func TestErrors(t *testing.T) {
	lm := NewLocationManager(BASE_WIDTH, BASE_HEIGHT)

	// Two new entities, one of which is added to LM.
	ent1 := &entity.TestEntity{TeRadius: 5}
	ent2 := &entity.TestEntity{TeRadius: 5}
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
	ent := &entity.TestEntity{TeRadius: 5}
	cm.AddEntAtLocation(ent, Combination{0.0, 0.0, angle})
	move := CoordDelta{s.x * 99, 0}
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
		tankSize{BASE_WIDTH, BASE_HEIGHT},      // Normal size.
		tankSize{BASE_WIDTH, BASE_HEIGHT},      // Square.
		tankSize{BASE_WIDTH, 10 * BASE_HEIGHT}, // Lopsided.
		tankSize{10 * BASE_WIDTH, BASE_HEIGHT}, // Also lopsided.
		tankSize{1.0, 1.0},                     // Tiny.
	}

	for _, limit := range tankSizes {
		t.Logf("Test tank with size: (%v, %v)", limit.x, limit.y)

		// Attempt to reach max range.
		tankLimit(t, limit)
	}
}

// Debug functions.
func (cm *LocationManager) PrintDebug() {
	fmt.Printf("Location Manager: %v\n", cm)
	for ii, hb := range cm.hitboxes {
		fmt.Printf("  Hitbox %v\n", ii)
		fmt.Printf("    %v\n", hb)
	}
	fmt.Printf("\n")
}

// Tests we can load a new LM with settings from global config.
func TestNew(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	lm := New()

	exp := coord{800, 800}
	if lm.maxPoint != exp {
		t.Errorf("Wrong max point. Expected %v, got %v.", exp, lm.maxPoint)
	}
}
