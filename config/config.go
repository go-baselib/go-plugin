package config

import (
	"fmt"
	"sort"

	"gopkg.in/yaml.v3"
)

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
	Name  string `yaml:"name"`
	Level int    `yaml:"level"`
}

func Init(conf []byte) {
	var (
		v = &Config{
			GRPC: GRPC{Host: "0.0.0.0", Port: 9999},
			Plugins: []Plugin{
				{Name: "hi"},
			},
		}
		err = yaml.Unmarshal(conf, v)
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("conf: %+v", v)

	gCfg = v
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
