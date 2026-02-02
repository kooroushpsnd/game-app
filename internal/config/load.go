package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load(configPath string) Config {
	var k = koanf.New(".")

	k.Load(confmap.Provider(defaultConfig ,".") ,nil)
	k.Load(file.Provider(configPath) ,yaml.Parser())

	var cfg Config
	if err := k.Unmarshal("" ,&cfg); err != nil {
		panic(err)
	}

	return cfg
}