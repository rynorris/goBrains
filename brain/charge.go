package brain

// An abstract unit of electrical charge.
type ChargeUnit float64

// The fixed at which charge disperses into the environment.
const chargeDecayRate = 0.02

// An abstract collection of electrical charge.
// Used to commonise code between different kinds of
// brain elements.
type Charge struct {
    currentCharge ChargeUnit
}

// Charge up this charge by strength ChargeUnits.
func (c *Charge) ChargeUp(strength ChargeUnit) {
    c.currentCharge += strength
}

// Decreases this charge by chargeDecayRate.
// Should be called once per time-step.
func (c *Charge) Decay() {
    c.currentCharge -= chargeDecayRate
    if c.currentCharge < 0 {
        c.currentCharge = 0
    }
}
