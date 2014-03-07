// A real-time neural network implementation.
package brain

const (
    defaultFiringThreshold = 1.0
    defaultFiringStrength = 0.8
    defaultChargeDecayRate = 0.02
)

type Node struct {
	firingThreshold float64
	firingStrength  float64
	chargeDecayRate float64
	currentCharge   float64
	outputs         []*Node
}

func (n *Node) Fire() {
	for _, out := range n.outputs {
		out.Charge(n.firingStrength)
	}
}

func (n *Node) Charge(strength float64) {
	n.currentCharge += strength
	for n.currentCharge >= n.firingThreshold {
		n.Fire()
		n.currentCharge -= n.firingThreshold
	}
}

func (n *Node) Update() {
	n.currentCharge -= n.chargeDecayRate
	if n.currentCharge <= 0 {
		n.currentCharge = 0
	}
}

func (n *Node) AddOutput(out *Node) {
	n.outputs = append(n.outputs, out)
}

func NewNode() (*Node) {
    return &Node{firingThreshold: defaultFiringThreshold, 
                 firingStrength: defaultFiringStrength, 
                 chargeDecayRate: defaultChargeDecayRate}
}
