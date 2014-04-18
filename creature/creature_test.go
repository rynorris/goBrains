/*
 * Testing for the creature package.
 */

package creature

import (
	"github.com/DiscoViking/goBrains/brain"
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/food"
	"github.com/DiscoViking/goBrains/locationmanager"
	"testing"
)

// Dummy brain for testing.
type testBrain struct {
	fired   int
	nodes   []*brain.Node
	outputs []brain.ChargedWorker
}

// Verify that a movement structure is as expected for a booster.
func CheckMove(t *testing.T, tb *booster, actual velocity, expected float64) {
	if (tb.btype == BoosterLinear) && (actual.move != expected) {
		t.Errorf("Expected linear velocity %v, got %v",
			expected,
			actual.move)
	}
	if (tb.btype == BoosterAngular) && (actual.rotate != expected) {
		t.Errorf("Expected rotational velocity %v, got %v",
			expected,
			actual.rotate)
	}
}

func newTestBrain() *testBrain {
	return &testBrain{
		fired:   0,
		nodes:   make([]*brain.Node, 0),
		outputs: make([]brain.ChargedWorker, 0),
	}
}

func (tb *testBrain) AddInputNode(node *brain.Node) {
	node.AddOutput(tb)
	tb.nodes = append(tb.nodes, node)
}

func (tb *testBrain) AddOutput(oput brain.ChargedWorker) {
	tb.outputs = append(tb.outputs, oput)
}

func (tb *testBrain) Work() {
	for _, node := range tb.nodes {
		node.Work()
	}
}

func (tb *testBrain) Charge(charge float64) {
	tb.fired++
}

// Basic antenna verification.
func TestAntenna(t *testing.T) {
	errorStr := "[%v] Expected test brain to have received %v firings, actually got %v."
	lm := locationmanager.NewLocationManager()
	creature := NewCreature(lm)
	tBrain := newTestBrain()
	creature.brain = tBrain

	// Add two antennae to the creature.
	antL := newAntenna(creature, AntennaLeft)
	antR := newAntenna(creature, AntennaRight)
	creature.inputs = append(creature.inputs, antL)
	creature.inputs = append(creature.inputs, antR)

	// Trigger detection, ensure nothing detected.
	antL.detect()
	antR.detect()
	tBrain.Work()
	if tBrain.fired > 0 {
		t.Errorf(errorStr, 1, 0, tBrain.fired)
	}

	// Add something to detect.  Is it detected?
	creature.lm.AddEntity(entity.TestEntity{100})
	antL.detect()
	tBrain.Work()
	if tBrain.fired != 1 {
		t.Errorf(errorStr, 2, 1, tBrain.fired)
	}
	antR.detect()
	tBrain.Work()
	if tBrain.fired != 2 {
		t.Errorf(errorStr, 3, 2, tBrain.fired)
	}

	// Detection of multiple entities at once.
	// Add another 99, for a total of 100 entities to detect.
	tBrain.fired = 0
	for ii := 0; ii < 99; ii++ {
		creature.lm.AddEntity(entity.TestEntity{100})
	}
	antL.detect()
	tBrain.Work()
	if tBrain.fired != 100 {
		t.Errorf(errorStr, 4, 100, tBrain.fired)
	}
}

// Basic mouth verification.
func TestMouth(t *testing.T) {
	errorStrHost := "[%v] Expected host vitality of %v, actually got %v."
	errorStrFood := "[%v] Expected food content of %v, actually got %v."
	lm := locationmanager.NewLocationManager()
	creature := NewCreature(lm)
	mot := newMouth(creature)
	creature.inputs = append(creature.inputs, mot)

	// This should be as expected, or this test will most definitely fail.
	if creature.vitality != 10 {
		t.Errorf(errorStrHost, 1, 10, creature.vitality)
	}

	// Attempt to consume when we overlap with only ourselves present.  No change in vitality.
	mot.detect()
	if creature.vitality != 10 {
		t.Errorf(errorStrHost, 2, 10, creature.vitality)
	}

	// Add some food to eat.  Try and eat it.  We do not deal with the *other* end.
	fd := food.New(lm, 10)

	mot.detect()
	if creature.vitality != 11 {
		t.Errorf(errorStrHost, 3, 11, creature.vitality)
	}
	if fd.GetContent() != 9 {
		t.Errorf(errorStrFood, 4, 9, fd.GetContent())
	}

	// And repeat. The first time might have been a fluke, right?
	mot.detect()
	if creature.vitality != 12 {
		t.Errorf(errorStrHost, 5, 12, creature.vitality)
	}
	if fd.GetContent() != 8 {
		t.Errorf(errorStrFood, 6, 8, fd.GetContent())
	}

	// Deplete the food entirely.
	for ii := 0; ii < 10; ii++ {
		mot.detect()
	}
	if creature.vitality != 20 {
		t.Errorf(errorStrHost, 7, 20, creature.vitality)
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
	}

	mot.detect()
	for ii := 0; ii < 100; ii++ {
		if muchFood[ii].GetContent() != 9 {
			t.Errorf("[%v] Failed on food number %v", 9, ii)
			t.Errorf(errorStrFood, 9, 9, muchFood[ii].GetContent())
			lm.PrintDebug()
		}
	}
	if creature.vitality != 110 {
		t.Errorf(errorStrHost, 10, 110, creature.vitality)
		lm.PrintDebug()
	}
}

// Booster behaviour verification.
func TestBoosters(t *testing.T) {
	host := NewCreature(locationmanager.NewLocationManager())
	tBrain := newTestBrain()
	host.brain = tBrain

	linBoost := newLinearBooster(host)
	angBoost := newAngularBooster(host)

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
		CheckMove(t, testBooster, host.movement, 0.1)

		// Ensure that the charge has definitely depleted after use.
		testBooster.Work()
		CheckMove(t, testBooster, host.movement, 0.1)
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
		CheckMove(t, testBooster, host.movement, 0.1)
	}

}

// High-level creature verification.
func TestCreature(t *testing.T) {
	errorStrLm := "[%v] Expected %v entities in LM, found %v."
	errorStrDead := "[%v] Creature expected %v, actually %v."

	lm := locationmanager.NewLocationManager()
	if lm.NumberOwned() != 0 {
		t.Errorf(errorStrLm, 1, 0, lm.NumberOwned())
	}

	// The new creature should have registered with the LM.
	creature := NewCreature(lm)
	if lm.NumberOwned() != 1 {
		t.Errorf(errorStrLm, 2, 1, lm.NumberOwned())
	}

	// It should be alive and happy and not immediately keel over dead.
	// This will impede our testing somewhat.
	if creature.Check() {
		t.Errorf(errorStrDead, 3, "alive", "dead")
		return
	}

	// If the creature runs out of vitality it will die.
	// This should also remove it from LM.
	creature.vitality = 0
	if !creature.Check() {
		t.Errorf(errorStrDead, 4, "dead", "alive")
	}
	if lm.NumberOwned() != 0 {
		t.Errorf(errorStrDead, 5, 0, lm.NumberOwned())
	}
}

// Cannibalism.  AKA. Hot creature-on-creature action.
func TestCannibalism(t *testing.T) {
	lm := locationmanager.NewLocationManager()
	creature := NewCreature(lm)

	// Creatures cannot eat other creatures (yet).  Attempts to eat other creatures result in a no-op.
	if creature.Consume() != 0 {
		t.Errorf("Creature successfully eaten.")
	}
}
