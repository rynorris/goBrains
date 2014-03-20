/*
 * Testing for the creature package.
 */

package creature

import (
	"github.com/DiscoViking/goBrains/brain"
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/locationmanager"
	"testing"
)

// Dummy brain for testing.
type testBrain struct {
	fired int
}

func (tb *testBrain) AddInputNode(node *brain.Node) {
	node.AddOutput(tb)
}

func (tb *testBrain) Charge(charge float64) {
	tb.fired++
}

func newTestCreature() *Creature {
	return &Creature{
		lm:       locationmanager.NewLocationManager(),
		brain:    nil,
		inputs:   make([]input, 0),
		vitality: 0,
	}
}

// Basic antenna verification.
func TestAntenna(t *testing.T) {
	errorStr := "[%v] Expected test brain to have received %v firings, actually got %v."
	creature := newTestCreature()
	tBrain := &testBrain{0}
	creature.brain = tBrain

	// Add two antennae to the creature.
	antL := NewAntenna(creature, AntennaLeft)
	antR := NewAntenna(creature, AntennaRight)
	creature.inputs = append(creature.inputs, antL)
	creature.inputs = append(creature.inputs, antR)

	// Trigger detection, ensure nothing detected.
	antL.detect()
	antR.detect()
	if tBrain.fired > 0 {
		t.Errorf(errorStr, 1, 0, tBrain.fired)
	}

	// Add something to detect.  Is it detected?
	creature.lm.AddEntity(entity.TestEntity{100})
	antL.detect()
	if tBrain.fired != 1 {
		t.Errorf(errorStr, 2, 1, tBrain.fired)
	}
	antR.detect()
	if tBrain.fired != 2 {
		t.Errorf(errorStr, 3, 2, tBrain.fired)
	}

	// Detection of multiple entities at once.
	for ii := 0; ii < 99; ii++ {
		creature.lm.AddEntity(entity.TestEntity{100})
	}
	antL.detect()
	if tBrain.fired != 101 {
		t.Errorf(errorStr, 4, 101, tBrain.fired)
	}
}
