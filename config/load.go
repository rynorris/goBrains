package config

import "code.google.com/p/gcfg"

func Load(filename string) (*Config, error) {
	cfg := &Config{}

	err := gcfg.ReadFileInto(cfg, filename)

	return cfg, err
}
