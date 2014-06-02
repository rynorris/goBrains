package brain

import "github.com/DiscoViking/goBrains/config"

// An abstract collection of electrical ChargeCarrier.
// Used to commonise code between different kinds of
// brain elements.
type ChargeCarrier struct {
	currentCharge float64
}

// ChargeCarrier up this ChargeCarrier by strength ChargeUnits.
func (c *ChargeCarrier) Charge(strength float64) {
	c.currentCharge += strength
}

// Decreases this ChargeCarrier by chargeDecayRate.
// Should be called once per time-step.
func (c *ChargeCarrier) Decay() {
	c.currentCharge -= config.Global.Brain.ChargeDecayRate
	if c.currentCharge < 0 {
		c.currentCharge = 0
	}
}
