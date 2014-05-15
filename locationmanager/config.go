package locationmanager

import "github.com/DiscoViking/goBrains/config"

var (
	TANKSIZEX float64
	TANKSIZEY float64
)

func LoadConfig(cfg *config.Config) {
	TANKSIZEX = float64(cfg.General.ScreenWidth)
	TANKSIZEY = float64(cfg.General.ScreenHeight)
}
