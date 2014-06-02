/*
 * DNA testing.
 *
 * Because while it may not look that way, you probably do want all of those genes.
 */

package genetics

import (
	"math"
	"testing"

	"github.com/DiscoViking/goBrains/config"
)

// Generate a reasonable-length DNA sequence.
func NewTestDna() *Dna {
	d := NewDna()
	for _, jj := range []float64{0.0, 0.1, -0.2, 0.3, -0.4, 0.5, -0.6, 0.7, -0.9, 1.0, -1.0} {
		d.AddGene(NewGene(jj))
	}
	return d
}

// Retreive the sequence in a DNA object.
func Retrieve(d *Dna) []float64 {
	c := d.GetValues()
	res := make([]float64, 0)
	for jj := range c {
		res = append(res, jj)
	}
	return res
}

// Compare values to those expected.
func Compare(t *testing.T, result, expected []float64) {
	diff := false

	if len(result) != len(expected) {
		diff = true
	} else {
		for i := range result {
			// Due to loss of precision on compression, check to 6 decimal places only.
			if int(math.Pow(result[i], 6)) != int(math.Pow(expected[i], 6)) {
				diff = true
			}
		}
	}

	if diff {
		t.Errorf("DNA has been unexpectedly mutated by packing and unpacking.\nBefore:\n%v\nAfter:\n%v",
			expected,
			result)
	}
}

// Verify that a sequence can be packed into DNA and unpacked again.
func VerifySequence(t *testing.T, seq []float64) {
	var res []float64
	d := NewDna()

	// Pack values.
	for _, jj := range seq {
		d.AddGene(NewGene(jj))
	}

	// Unpack values and compare.
	res = Retrieve(d)
	Compare(t, res, seq)

	// Unpack again.  DNA should not deplete.
	res = Retrieve(d)
	Compare(t, res, seq)
}

func TestSequences(t *testing.T) {
	VerifySequence(t, []float64{})
	VerifySequence(t, []float64{0.0})
	VerifySequence(t, []float64{0.0, 0.1, -0.2, 0.3, -0.4, 0.5, -0.6, 0.7, -0.9, 1.0, -1.0})
}

// Do some basic kick-the-tires testing of generating new DNA sequences.
func TestBreeding(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	d := NewTestDna()

	// Inbreed a few generations.
	// Verify that at least one generation has no change, and that one shows a mutation.
	identical := false
	mutated := false

	for i := 0; i < 100; i++ {
		nd := d.Breed(d)
		if CompareSequence(d, nd) {
			identical = true
		} else {
			mutated = true
		}
		nd = d
	}

	if !identical {
		t.Errorf("Breeding failed to produce an identical child.")
	}
	if !mutated {
		t.Errorf("Breeding failed to produce a mutated child.")
	}
}

// Ensure that cloning produces identical DNA.
func TestCloning(t *testing.T) {
	config.Load("../config/test_config.gcfg")
	d := NewTestDna()
	newD := d.Clone()
	if !CompareSequence(d, newD) {
		t.Errorf("Cloning produced a non-identical sequence.")
	}
}

// Test that we cannot breed two different-length sequences.
func TestBadBreeding(t *testing.T) {
	da := NewDna()
	db := NewDna()

	for i := 0; i < 5; i++ {
		da.AddGene(NewGene(0.0))
		db.AddGene(NewGene(0.0))
	}

	da.AddGene(NewGene(0.0))

	dn := da.Breed(db)
	if dn != nil {
		t.Errorf("Successfully bred two incompatible sequences.")
	}
}
