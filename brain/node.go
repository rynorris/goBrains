// A real-time neural network implementation.
package brain

const (
    defaultFiringThreshold = 1.0
    defaultFiringStrength  = 0.8
)

type Node struct {
    Charge
    firingThreshold ChargeUnit   // When currentCharge crosses this threshold, the node will Fire().
    firingStrength  ChargeUnit   // The node fires by charging it's output by this amount.
    outputs         []Chargeable // An array of outputs that will get charged when this Node fires.
}

// Charge up all outputs by our firingStrength.
func (n *Node) Fire() {
    for _, out := range n.outputs {
        out.ChargeUp(n.firingStrength)
    }
}

// Work is where a Node does all it's processing.
// The Node should not affect anything outside itself anywhere EXCEPT
// in the Work method.
func (n *Node) Work() {
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
