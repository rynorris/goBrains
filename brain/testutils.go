package brain

import "math"

const Delta = 0.00001

func ChargesAreEqual(x, y ChargeUnit) bool {
    if dist := math.Abs(float64(x) - float64(y)); dist > Delta {
        return false
    }

    return true
}
