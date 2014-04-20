/*
 * Utilities testing.
 */

package genetics

import "testing"

func TestCompare(t *testing.T) {
	dnaShort := NewDna()
	dna0 := NewDna()
	dna1 := NewDna()
	dna0.AddGene(NewGene(0.0))
	dna1.AddGene(NewGene(1.0))

	// Verify that two identical sequences match.
	if !CompareSequence(dna0, dna0) {
		t.Errorf("Two identical sequences did not match.")
	}

	// Verify that two different sequences do not match.
	if CompareSequence(dna0, dna1) {
		t.Errorf("Two different sequences matched.")
	}

	// Different length sequences do not match.
	if CompareSequence(dnaShort, dna0) {
		t.Errorf("Two different-length sequences matched.")
	}
}
