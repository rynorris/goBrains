package config

import "code.google.com/p/gcfg"

var (
	Global Config
)

func Load(filename string) (*Config, error) {
	Global = Config{}

	err := gcfg.ReadFileInto(&Global, filename)

	return &Global, err
}
