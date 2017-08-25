package config

import "gopkg.in/gcfg.v1"

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
