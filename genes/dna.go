/*
 * DNA.
 *
 * The magic that happens when genes get together.
 */

package genes

// Add a value to the genetic sequence.
func (d *Dna) AddValue(value float64) {
	g := NewGene(value)
	d.sequence = append(d.sequence, g)
}

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

// Generate a new sequence of DNA.
// This is initially empty, until populated.
func NewDna() *Dna {
	return &Dna{(make([]*gene, 0))}
}
