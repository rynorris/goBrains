// A real-time neural network implementation.
package brain

const (
	defaultFiringThreshold = 1.0
	defaultFiringStrength  = 0.8
	defaultChargeDecayRate = 0.02
)

type Node struct {
	firingThreshold float64 // When currentCharge crosses this threshold, the node will Fire().
	firingStrength  float64 // The node fires by charging it's output by this amount.
	chargeDecayRate float64 // The amount currentCharge decreases by per Update().
	currentCharge   float64 // The current electrical charge stored in the Node.
	outputs         []*Node // An array of outputs that will get charged when this Node fires.
}

// Charge up all outputs by our firingStrength.
func (n *Node) Fire() {
	for _, out := range n.outputs {
		out.Charge(n.firingStrength)
	}
}

// Charge up this node by strength.
// If this node then has more charge than the firing threshold, fire.
func (n *Node) Charge(strength float64) {
	n.currentCharge += strength
	for n.currentCharge >= n.firingThreshold {
		n.Fire()
		n.currentCharge -= n.firingThreshold
	}
}

// Decreases this node's charge by chargeDecayRate.
// Should be called once per time-step.
func (n *Node) Update() {
	n.currentCharge -= n.chargeDecayRate
	if n.currentCharge <= 0 {
		n.currentCharge = 0
	}
}

// Adds a new output to this node.
func (n *Node) AddOutput(out *Node) {
	n.outputs = append(n.outputs, out)
}

// Creates a new node with default values.
func NewNode() *Node {
	return &Node{firingThreshold: defaultFiringThreshold,
		firingStrength:  defaultFiringStrength,
		chargeDecayRate: defaultChargeDecayRate}
}
