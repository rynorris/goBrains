/*
 * TestBrain.go
 *
 * A dummy brain for testing purposes.
 */

package creature

import (
	"github.com/DiscoViking/goBrains/brain"
	"github.com/DiscoViking/goBrains/genetics"
)

// Dummy brain for testing.
type testBrain struct {
	fired   int
	nodes   []*brain.Node
	outputs []brain.ChargedWorker
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

func (tb *testBrain) GenesNeeded() int {
	return 0
}

func (tb *testBrain) Work() {
	for _, node := range tb.nodes {
		node.Work()
	}
}

func (tb *testBrain) Charge(charge float64) {
	tb.fired++
}

func (tb *testBrain) Restore(d *genetics.Dna) {
	return
}
