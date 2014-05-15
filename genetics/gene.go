/*
 * Genes.
 *
 * Playing about with single units of genetic information.
 */

package genetics

import "math/rand"

const (

	// Length of a gene, in bytes.
	GENELEN = 4

	// A gene contains four bytes for a signed integer.
	MAXGENEVAL = (2147483647.0)
)

// Pack a value into a gene.
// This takes a floating point number between 1 and -1 and converts it into an integer which can be stored as binary.
func (g *gene) Pack(value float64) {
	if value > 1 {
		value = 1
	} else if value < -1 {
		value = -1
	}

	g.value = int32(value * MAXGENEVAL)
}

// Unpack a value from a gene.
// This converts the unsigned integer back into a floating point number in the correct range.
func (g *gene) Unpack() float64 {
	return (float64(g.value) / MAXGENEVAL)
}

// Generate a new gene from an existing one, with mutations.
func (g *gene) Copy() *gene {
	ng := NewGene(0.0)
	ng.value = g.value

	// Generate a map of mutations to make.  0.05% chance mutation per bit.
	mutateMap := int32(0)

	for i := 0; i < (GENELEN * 8); i++ {
		mutateMap <<= 1
		if rand.Intn(mutationRate) == 0 {
			mutateMap |= 1
		}
	}

	// Apply the mutations to the new gene.
	ng.value ^= mutateMap

	return ng
}

// Generate a new random gene.
func NewRandomGene() *gene {
	val := rand.Float64()

	// This value may be negative, but rand only generates positive numbers.
	if rand.Intn(2) == 0 {
		val = -1 * val
	}

	return NewGene(val)
}

func NewGene(value float64) *gene {
	g := gene{0.0}
	g.Pack(value)
	return &g
}
