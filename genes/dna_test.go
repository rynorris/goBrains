/*
 * DNA testing.
 *
 * Because while it may not look that way, you probably do want all of those genes.
 */

package genes

import (
	"math"
	"testing"
)

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
		d.AddValue(jj)
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
