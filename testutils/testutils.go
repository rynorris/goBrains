package testutils

import "math"

const Delta = 0.0000001

func FloatsAreEqual(x, y float64) bool {
	if dist := math.Abs(x - y); dist > Delta {
		return false
	}

	return true
}
