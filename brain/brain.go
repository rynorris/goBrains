package brain

import (
	"github.com/DiscoViking/goBrains/genes"
	"log"
)

/*
 * The Brain is composed of inputs, central nodes, outputs and synapses.
 * They are arranged as such:
 *
 * inputNodes   : o  o  o
 * inSynapses   :  \/ \/
 * centralNodes :  o   o
 * outSynapses  : /\   /\
 * outputs      :o o   o o
 *
 * Charge flows from top to bottom in the diagram.
 *
 * Each tick, the input Nodes should be charged appropriately (by the external stimuli),
 * and then Update() called exactly once.
 */
type Brain struct {
	inputNodes   []*Node
	centralNodes []*Node
	outputs      []ChargedWorker
	inSynapses   []*Synapse
	outSynapses  []*Synapse
}

// Updates all the nodes in the brain.
// Should be called exactly once per tick.
func (b *Brain) Update() {
	for _, node := range b.inputNodes {
		node.Work()
	}
	for _, synapse := range b.inSynapses {
		synapse.Work()
	}
	for _, node := range b.centralNodes {
		node.Work()
	}
	for _, synapse := range b.outSynapses {
		synapse.Work()
	}
	for _, output := range b.outputs {
		output.Work()
	}
}

// Returns a pointer to a new Brain with numCentralNodes central nodes.
// It will have no inputs, no outputs and no synapses.
func NewBrain(numCentralNodes int) *Brain {
	b := Brain{}
	for i := 0; i < numCentralNodes; i++ {
		node := NewNode()
		b.centralNodes = append(b.centralNodes, node)
	}

	return &b
}

// Restore a set of synapses.
func (b *Brain) restoreSynapses(c chan float64, syns []*Synapse) {
	var ok bool
	for _, syn := range syns {
		syn.permittivity, ok = <-c
		if !ok {
			log.Fatal("Not enough information in DNA to restore brain.")
		}
	}
}

// Restore a brain using genetic information.
// This can only occur if the structure matches that of the information.
func (b *Brain) Restore(d *genes.Dna) {
	c := d.GetValues()
	b.restoreSynapses(c, b.inSynapses)
	b.restoreSynapses(c, b.outSynapses)
	_, ok := <-c
	if ok {
		log.Fatal("Unused information in DNA after restore.")
	}
}

// Adds an input node to the brain.
// Automatically connects it to all central nodes with synapses.
func (b *Brain) AddInputNode(input *Node) {
	b.inputNodes = append(b.inputNodes, input)
	for _, node := range b.centralNodes {
		s := Synapse{}
		s.output = node
		input.AddOutput(&s)
		b.inSynapses = append(b.inSynapses, &s)
	}
}

// Adds an output node to the brain.
// Automatically connects it to all central nodes with synapses.
func (b *Brain) AddOutput(output ChargedWorker) {
	b.outputs = append(b.outputs, output)
	for _, node := range b.centralNodes {
		s := Synapse{}
		s.output = output
		node.AddOutput(&s)
		b.outSynapses = append(b.outSynapses, &s)
	}
}
