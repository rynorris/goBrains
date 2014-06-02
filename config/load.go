package config

import "code.google.com/p/gcfg"

var (
	Global Config
)

func Load(filename string) {
	Global = Config{}

	err := gcfg.ReadFileInto(&Global, filename)

	if err != nil {
		panic(err)
	}
}
