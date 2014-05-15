package brain

import "github.com/DiscoViking/goBrains/config"

var (
	// Rate at which charge disperses into the environment.
	chargeDecayRate float64

	// The maximum ChargeCarrier a synapse can hold. If it is charged beyond this limit
	// any extra ChargeCarrier is lost.
	synapseMaxCharge float64

	// The proportion of a Synapse's ChargeCarrier which it will attempt to convey
	// to it's output each time it does work.
	// Note: the actual conveyed ChargeCarrier will be further modified by permittivity.
	synapseOutputScale float64

	defaultFiringThreshold float64
	defaultFiringStrength  float64
)

func LoadConfig(cfg *config.Config) {
	chargeDecayRate = cfg.Brain.ChargeDecayRate
	synapseMaxCharge = cfg.Brain.SynapseMaxCharge
	synapseOutputScale = cfg.Brain.SynapseOutputScale
	defaultFiringThreshold = cfg.Brain.NodeFiringThreshold
	defaultFiringStrength = cfg.Brain.NodeFiringStrength
}
