package config

import "sort"

var gCfg *Config

type Config struct {
	GRPC    GRPC     `yaml:"grpc"`
	Plugins []Plugin `yaml:"plugins"`
}

type GRPC struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Plugin struct {
	Name  string
	Level int
}

func Init() {
	gCfg = &Config{
		GRPC: GRPC{Host: "0.0.0.0", Port: 9999},
		Plugins: []Plugin{
			{Name: "hi"},
		},
	}
	sort.Slice(gCfg.Plugins, func(i, j int) bool {
		if gCfg.Plugins[i].Level != gCfg.Plugins[j].Level {
			return gCfg.Plugins[i].Level < gCfg.Plugins[j].Level
		}

		return gCfg.Plugins[i].Name < gCfg.Plugins[j].Name
	})
}

func GetConfig() *Config {
	return gCfg
}
