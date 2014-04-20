/*
 * Gene testing.
 *
 * Ensuring your genes are stable since 1989.
 */

package genetics

import (
	"math"
	"testing"
)

func VerifyRepack(t *testing.T, value, expected float64) {
	x := NewGene(0.0)

	x.Pack(value)
	result := x.Unpack()

	// Due to loss of precision on compression, check to 6 decimal places only.
	if int(math.Pow(result, 6)) != int(math.Pow(expected, 6)) {
		t.Errorf("Packed and unpacked value %v.  Output %v, expected %v",
			value,
			result,
			expected)
	}
}

// Verify that values are consistent across unpacking and repacking.
func TestRepacking(t *testing.T) {
	VerifyRepack(t, 0.0, 0.0)
	VerifyRepack(t, 0.45, 0.45)
	VerifyRepack(t, -0.45, -0.45)
	VerifyRepack(t, 1.0, 1.0)
	VerifyRepack(t, -1.0, -1.0)
	VerifyRepack(t, 100.0, 1.0)
	VerifyRepack(t, -100.0, -1.0)
}

// Verify behaviour on gene copies.
// We should see a mixture of identical copies and mutated copies.
func TestCopy(t *testing.T) {
	g := NewGene(77)
	identical := false
	mutated := false

	for i := 0; i < 100; i++ {
		ng := g.Copy()
		if ng.value == g.value {
			identical = true
		} else {
			mutated = true
		}
	}

	if !identical {
		t.Errorf("No identical child genes were produced by the copy process.")
	}
	if !mutated {
		t.Errorf("No mutated child genes were produced by the copy process.")
	}
}
