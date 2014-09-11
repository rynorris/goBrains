/*
 * Testing for the creature package.
 */

package creature

import (
	"testing"

	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/food"
	"github.com/DiscoViking/goBrains/genetics"
	"github.com/DiscoViking/goBrains/locationmanager"
	"github.com/DiscoViking/goBrains/testutils"
)

// Dummy detection method
type testInput struct {
	triggered int
}

func (ti *testInput) detect() {
	ti.triggered++
}

func CheckVelocity(t *testing.T, id int, c *creature, actual velocity) {
	if c.movement.move != actual.move || c.movement.rotate != actual.rotate {
		t.Errorf("[%v] Expected movement (%v, %v), got (%v, %v).",
			id, c.movement.move, c.movement.rotate, actual.move, actual.rotate)
	}
}

// Verify that a movement structure is as expected for a booster.
func CheckMove(t *testing.T, tb *booster, actual velocity, expected float64) {
	if (tb.btype == BoosterLinear) && (!testutils.FloatsAreEqual(actual.move, expected)) {
		t.Errorf("Expected linear velocity %v, got %v",
			expected,
			actual.move)
	}
	if (tb.btype == BoosterAngular) && (!testutils.FloatsAreEqual(actual.rotate*LinPerAng, expected)) {
		t.Errorf("Expected rotational velocity %v, got %v",
			expected,
			actual.rotate*LinPerAng)
	}
}

// Verify that DNA matches the creature.
func checkDnaLen(t *testing.T, c *creature) {
	gn := c.brain.GenesNeeded()
	cv := c.dna.GetValues()
	ga := 0
	for jj := range cv {
		jj = jj
		ga++
	}

	if gn != ga {
		t.Errorf("Mismatch of DNA with creature.  Expected %v, actual %v", gn, ga)
	}
}

// Verify collection of values from a creature.
func TestValues(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	c := New(locationmanager.New()).(*creature)
	c.Color()
	c.Radius()
}

// Basic antenna verification.
func TestAntenna(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	errorStr := "[%v] Expected test brain to have received %v firings, actually got %v."
	lm := locationmanager.New()
	lm.StartAtOrigin()
	creature := New(lm).(*creature)
	tBrain := newTestBrain()
	creature.brain = tBrain

	// Add two antennae to the creature.
	antL := creature.AddAntenna(AntennaLeft)
	antR := creature.AddAntenna(AntennaRight)

	// Trigger detection, ensure nothing detected.
	antL.detect()
	antR.detect()
	tBrain.Work()
	if tBrain.fired > 0 {
		t.Errorf(errorStr, 1, 0, tBrain.fired)
	}

	// Add something to detect.  Is it detected?
	// As the antenna has three inputs and it will charge three times.
	creature.lm.AddEntity(&entity.TestEntity{TeRadius: 100})
	antL.detect()
	tBrain.Work()
	if tBrain.fired != 3 {
		t.Errorf(errorStr, 2, 3, tBrain.fired)
	}
	antR.detect()
	tBrain.Work()
	if tBrain.fired != 6 {
		t.Errorf(errorStr, 3, 6, tBrain.fired)
	}

	if t.Failed() {
		// Later tests will be incredibly verbose, so stop here if we have failed already.
		return
	}

	// Detection of multiple entities at once.
	// Add another 99, for a total of 100 entities to detect.
	tBrain.fired = 0
	for ii := 0; ii < 99; ii++ {
		creature.lm.AddEntity(&entity.TestEntity{TeRadius: 100})
	}
	antL.detect()
	tBrain.Work()
	if tBrain.fired != 300 {
		t.Errorf(errorStr, 4, 300, tBrain.fired)
	}
}

// Basic mouth verification.
func TestMouth(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	errorStrHost := "[%v] Expected host vitality of %v, actually got %v."
	errorStrFood := "[%v] Expected food content of %v, actually got %v."
	lm := locationmanager.New()
	lm.StartAtOrigin()
	creature := New(lm).(*creature)
	mot := creature.AddMouth()

	// This should be as expected, or this test will most definitely fail.
	if creature.vitality != InitialVitality {
		t.Errorf(errorStrHost, 1, InitialVitality, creature.vitality)
	}

	// Attempt to consume when we overlap with only ourselves present.  No change in vitality.
	mot.detect()
	if creature.vitality != InitialVitality {
		t.Errorf(errorStrHost, 2, InitialVitality, creature.vitality)
	}

	// Add some food to eat.  Try and eat it.  We do not deal with the *other* end.
	fd := food.New(lm, 10)
	lm.ChangeLocation(mot.location, fd)

	mot.detect()
	if creature.vitality != InitialVitality+1 {
		t.Errorf(errorStrHost, 3, InitialVitality+1, creature.vitality)
	}
	if fd.GetContent() != 9 {
		t.Errorf(errorStrFood, 4, 9, fd.GetContent())
	}

	// And repeat. The first time might have been a fluke, right?
	mot.detect()
	if creature.vitality != InitialVitality+2 {
		t.Errorf(errorStrHost, 5, InitialVitality+2, creature.vitality)
	}
	if fd.GetContent() != 8 {
		t.Errorf(errorStrFood, 6, 8, fd.GetContent())
	}
	if t.Failed() {
		// Later tests will be incredibly verbose, so stop here if we have failed already.
		return
	}

	// Deplete the food entirely.
	for ii := 0; ii < 10; ii++ {
		mot.detect()
	}
	if creature.vitality != InitialVitality+10 {
		t.Errorf(errorStrHost, 7, InitialVitality+10, creature.vitality)
	}
	if fd.GetContent() != 0 {
		t.Errorf(errorStrFood, 8, 0, fd.GetContent())
	}

	// This food is done.
	lm.RemoveEntity(fd)

	// Test eating lots of food at once.
	// Reset the vitailty first.
	creature.vitality = 10
	muchFood := make([]*food.Food, 100)
	for ii := 0; ii < 100; ii++ {
		muchFood[ii] = food.New(lm, 10)
		lm.ChangeLocation(mot.location, muchFood[ii])
	}

	mot.detect()
	for ii := 0; ii < 100; ii++ {
		if muchFood[ii].GetContent() != 9 {
			t.Errorf("[%v] Failed on food number %v", 9, ii)
			t.Errorf(errorStrFood, 9, 9, muchFood[ii].GetContent())
		}
	}
	if creature.vitality != 110 {
		t.Errorf(errorStrHost, 10, 110, creature.vitality)
	}
}

// Booster behaviour verification.
func TestBoosters(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	host := New(locationmanager.New()).(*creature)
	tBrain := newTestBrain()
	host.brain = tBrain

	linBoost, angBoost := host.AddBoosters()

	testBoosters := []*booster{
		linBoost,
		angBoost,
	}

	for _, testBooster := range testBoosters {

		// Reset host velocity.
		host.movement = velocity{0, 0}

		// Trigger the booster with no charge.  No velocity as a result.
		testBooster.Work()
		CheckMove(t, testBooster, host.movement, 0.0)

		// Add change.  Check that this results in movement after the work phase.
		testBooster.Charge(0.1)
		CheckMove(t, testBooster, host.movement, 0.0)

		testBooster.Work()
		CheckMove(t, testBooster, host.movement, 0.02)

		// Ensure that the charge has definitely depleted after use.
		testBooster.Work()
		CheckMove(t, testBooster, host.movement, 0.02)
	}

	// Test that activating both boosters at once works.
	host.movement = velocity{0, 0}
	for _, testBooster := range testBoosters {
		testBooster.Charge(0.1)
	}
	for _, testBooster := range testBoosters {
		testBooster.Work()
	}
	for _, testBooster := range testBoosters {
		CheckMove(t, testBooster, host.movement, 0.02)
	}

	// Test overcharging boosters (in both directions), and that they are limited.
	host.movement = velocity{0, 0}
	for _, testBooster := range testBoosters {
		testBooster.Charge(9999)
		testBooster.Work()
	}
	CheckMove(t, linBoost, host.movement, MaxLinVel)
	CheckMove(t, angBoost, host.movement, MaxLinVel)

	host.movement = velocity{0, 0}
	for _, testBooster := range testBoosters {
		testBooster.Charge(-9999)
		testBooster.Work()
	}
	CheckMove(t, linBoost, host.movement, -MaxLinVel)
	CheckMove(t, angBoost, host.movement, -MaxLinVel)
}

// Test creature movement.
func TestMovement(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	c := New(locationmanager.New()).(*creature)

	// Creatures start with no velocity.
	CheckVelocity(t, 1, c, velocity{0, 0})

	// Add velocity.  Get the immediate result.
	vel := velocity{0.5, 0.5}
	c.movement = vel
	CheckVelocity(t, 2, c, vel)

	// Setup some speeds.  Ensure that they degrade after work.
	vel.move, vel.rotate = vel.move*SpeedDegredation, vel.rotate*SpeedDegredation
	c.Work()
	c.Check()
	CheckVelocity(t, 3, c, vel)

	// Degrade some more.
	vel.move, vel.rotate = vel.move*SpeedDegredation, vel.rotate*SpeedDegredation
	c.Work()
	c.Check()
	CheckVelocity(t, 4, c, vel)
}

func checkDetection(t *testing.T, id, expected int, ti *testInput) {
	if ti.triggered != expected {
		t.Errorf("[%v] Expected %v triggerings, got %v.",
			id, expected, ti.triggered)
	}
}

// Test that inputs are called to detect things.
func TestDetection(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	c := New(locationmanager.New()).(*creature)
	ti := &testInput{}

	// No inputs, nothing going.
	c.Work()
	checkDetection(t, 0, 0, ti)

	// All for one, and one for all.
	c.inputs = append(c.inputs, ti)
	c.Work()
	checkDetection(t, 1, 1, ti)

	// Getting a bit crowded in here.
	ti.triggered = 0
	for ii := 0; ii < 99; ii++ {
		c.inputs = append(c.inputs, ti)
	}
	c.Work()
	checkDetection(t, 2, 100, ti)
}

// All things must die, eventually.
func TestMortality(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	errorStrLm := "[%v] Expected %v entities in LM, found %v."
	errorStrDead := "[%v] Creature expected %v, actually %v."

	lm := locationmanager.New()
	if lm.NumberOwned() != 0 {
		t.Errorf(errorStrLm, 1, 0, lm.NumberOwned())
	}

	// The new creature should have registered with the LM.
	creature := NewSimple(lm).(*creature)
	if lm.NumberOwned() != 1 {
		t.Errorf(errorStrLm, 2, 1, lm.NumberOwned())
	}

	// It should be alive and happy and not immediately keel over dead.
	// This will impede our testing somewhat.
	if creature.Check() {
		t.Errorf(errorStrDead, 3, "alive", "dead")
		return
	}

	// Vitality is capped.
	creature.vitality = MaxVitality + 10
	creature.Check()
	if creature.vitality > MaxVitality {
		t.Errorf("Expected vitality should be capped at %v, but actually at %v.",
			creature.vitality,
			MaxVitality)
	}

	// If the creature runs out of vitality it will die.
	// This should also remove it from LM.
	creature.vitality = 0
	if !creature.Check() {
		t.Errorf(errorStrDead, 4, "dead", "alive")
	}
	if lm.NumberOwned() != 0 {
		t.Errorf(errorStrLm, 5, 0, lm.NumberOwned())
	}
}

// Cannibalism.  AKA. Hot creature-on-creature action.
func TestCannibalism(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	lm := locationmanager.New()
	creature := NewSimple(lm).(*creature)

	// Creatures cannot eat other creatures (yet).  Attempts to eat other creatures results in a no-op.
	if creature.Consume() != 0 {
		t.Errorf("Creature successfully eaten.")
	}
}

// Creaturesmakecreatures.  <Insert unsuitable joke here.>
func TestBreeding(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	lm := locationmanager.New()
	mother := NewSimple(lm).(*creature)
	father := NewSimple(lm).(*creature)
	newChild := false

	// New child should be a mixture of the parents.
	// Test twice, as there is a small but finite chance it's a clone.
	for i := 0; i < 2; i++ {
		c, err := mother.Breed(father)
		child := c.(*creature)
		if err != nil {
			t.Errorf("Error whilst breeding: %v\n", err)
		}
		if !genetics.CompareSequence(mother.dna, child.dna) {
			newChild = true
		}
		if !genetics.CompareSequence(father.dna, child.dna) {
			newChild = true
		}
	}

	if !newChild {
		t.Errorf("Child was a clone of it's parent.")
	}
}

// Test the cloning facilities.
func TestCloning(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	lm := locationmanager.New()
	original := NewSimple(lm).(*creature)
	clone := original.Clone().(*creature)

	if !genetics.CompareSequence(original.dna, clone.dna) {
		t.Errorf("Clone does not match original creature.")
	}
}

// Test random DNA generation works.
func TestPrepare(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	var c *creature
	lm := locationmanager.New()

	// Verify that an empty creature has the correct DNA.
	c = New(lm).(*creature)
	c.Prepare()
	checkDnaLen(t, c)

	// Verify that a simple creature has the correct DNA.
	c = NewSimple(lm).(*creature)
	checkDnaLen(t, c)
}
