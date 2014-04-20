/*
 * DNA.
 *
 * The magic that happens when genes get together.
 */

package genetics

import "math/rand"

// Channel over which we retrieve values from a sequence of DNA.
func (d *Dna) GetValues() chan float64 {
	c := make(chan float64)

	go func() {
		for _, g := range d.sequence {
			c <- g.Unpack()
		}
		close(c)
	}()

	return c
}

// Breed with an external DNA strand to produce a child sequence.
// Simulate genetic recombination, rather than randomly picking genes from each.
func (dx *Dna) Breed(dy *Dna) *Dna {
	if len(dx.sequence) != len(dy.sequence) {
		// Attempt to breed two sequences which are not compatible.  Abort.
		return nil
	}

	dn := NewDna()
	active := dx
	other := dy

	// Equal bias to start with the mother or father sequence.
	if rand.Intn(1) == 0 {
		active, other = other, active
	}

	for i := 0; i < len(dx.sequence); i++ {
		dn.AddGene(active.sequence[i].Copy())

		// Perform on average once per sequence a recombination switch.
		if rand.Intn(len(dx.sequence)) == 0 {
			active, other = other, active
		}
	}

	return dn
}

// Cloning, like breeding but with less parents and more science.
func (d *Dna) Clone() *Dna {
	dn := NewDna()
	for i := 0; i < len(d.sequence); i++ {
		dn.AddGene(NewGene(d.sequence[i].Unpack()))
	}
	return dn
}

// Add a gene to the genetic sequence.
func (d *Dna) AddGene(g *gene) {
	d.sequence = append(d.sequence, g)
}

// Generate a new sequence of DNA.
// This is initially empty, until populated.
func NewDna() *Dna {
	return &Dna{(make([]*gene, 0))}
}
