package iomanager

import "github.com/DiscoViking/goBrains/config"

var (
	width  int = 800
	height int = 800
)

func LoadConfig(cfg *config.Config) {
	width = cfg.General.ScreenWidth
	height = cfg.General.ScreenHeight
}
