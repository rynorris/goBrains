/*
 * Gene testing.
 *
 * Ensuring your genes are stable since 1989.
 */

package genes

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

func TestRepacking(t *testing.T) {
	VerifyRepack(t, 0.0, 0.0)
	VerifyRepack(t, 0.45, 0.45)
	VerifyRepack(t, -0.45, -0.45)
	VerifyRepack(t, 1.0, 1.0)
	VerifyRepack(t, -1.0, -1.0)
	VerifyRepack(t, 100.0, 1.0)
	VerifyRepack(t, -100.0, -1.0)
}
