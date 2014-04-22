/*
 * Genetic structures.
 *
 * For all your storing-genetic-information needs.
 */

package genetics

// A single unit of genetic information.
// This contains a single signed integer.
type gene struct {
	value int32
}

// A full sequence of genetic information.
// This is variable in length, as the amount of genetic information is dependent on the brain's structure.
type Dna struct {
	sequence []*gene
}
