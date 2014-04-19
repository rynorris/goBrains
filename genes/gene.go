/*
 * Genes.
 *
 * Playing about with single units of genetic information.
 */

package genes

import (
	"bytes"
	"encoding/binary"
	"log"
)

const (
	// A gene contains four bytes for a signed integer.
	MAXGENEVAL = (32767.0)
)

// Pack a value into a gene.
// This takes a floating point number between 1 and -1 and converts it into an integer which can be stored as binary.
func (g *gene) Pack(value float64) {
	if value > 1 {
		value = 1
	} else if value < -1 {
		value = -1
	}

	intVal := int64(value * MAXGENEVAL)
	binary.PutVarint(g.value, intVal)
}

// Unpack a value from a gene.
// This converts the unsigned integer back into a floating point number in the correct range.
func (g *gene) Unpack() float64 {
	rdr := bytes.NewReader(g.value)
	val, err := binary.ReadVarint(rdr)
	if (val == 0) && (err != nil) {
		log.Fatal(err)
	}

	return (float64(val) / MAXGENEVAL)
}

func NewGene(value float64) *gene {
	g := gene{(make([]byte, 4))}
	g.Pack(value)
	return &g
}
