package genetics

import "github.com/DiscoViking/goBrains/config"

var (
	mutationRate int = 500 // The rate at which bit mutate. Probability of bit flip is 1/n.
)

func LoadConfig(cfg *config.Config) {
	mutationRate = cfg.Genetics.MutationRate
}
