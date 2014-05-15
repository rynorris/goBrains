// A real-time neural network implementation.
package brain

type Node struct {
	ChargeCarrier
	firingThreshold float64      // When currentCharge crosses this threshold, the node will Fire().
	firingStrength  float64      // The node fires by charging it's output by this amount.
	outputs         []Chargeable // An array of outputs that will get charged when this Node fires.
}

// ChargeCarrier up all outputs by our firingStrength.
func (n *Node) Fire() {
	for _, out := range n.outputs {
		out.Charge(n.firingStrength)
	}
}

// Work is where a Node does all it's processing.
// The Node should not affect anything outside itself anywhere EXCEPT
// in the Work method.
func (n *Node) Work() {
	if n.currentCharge < 0 {
		n.currentCharge = 0
	}
	for n.currentCharge >= n.firingThreshold {
		n.Fire()
		n.currentCharge -= n.firingThreshold
	}
	n.Decay()
}

// Adds a new output to this node.
func (n *Node) AddOutput(out Chargeable) {
	n.outputs = append(n.outputs, out)
}

// Creates a new node with default values.
func NewNode() *Node {
	return &Node{firingThreshold: defaultFiringThreshold,
		firingStrength: defaultFiringStrength}
}
