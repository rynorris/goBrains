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
	nodes []*brain.Node
}

func newTestBrain() *testBrain {
	return &testBrain{
		fired: 0,
		nodes: make([]*brain.Node, 0),
	}
}

func (tb *testBrain) AddInputNode(node *brain.Node) {
	node.AddOutput(tb)
	tb.nodes = append(tb.nodes, node)
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
	antL := NewAntenna(creature, AntennaLeft)
	antR := NewAntenna(creature, AntennaRight)
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
