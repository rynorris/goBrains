package iomanager

import "github.com/DiscoViking/goBrains/config"

var (
	width  int
	height int
)

func LoadConfig(cfg *config.Config) {
	width = cfg.General.ScreenWidth
	height = cfg.General.ScreenHeight
}
