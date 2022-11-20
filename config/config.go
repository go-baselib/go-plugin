package config

import (
	"fmt"
	"sort"

	"gopkg.in/yaml.v3"
)

var gCfg *Config

type Config struct {
	GRPC    GRPC              `yaml:"grpc"`
	Plugins []Plugin          `yaml:"plugins"`
	plugins map[string]Plugin `yaml:"-"`
}

type GRPC struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Plugin struct {
	Enabled bool   `yaml:"enabled"`
	Name    string `yaml:"name"`
	Level   int    `yaml:"level"`
	Path    string `yaml:"path"`
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
	v.plugins = make(map[string]Plugin, len(v.Plugins))
	for _, p := range v.Plugins {
		v.plugins[p.Name] = p
	}

	gCfg = v
	fmt.Printf("conf: %+v", gCfg)
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

func GetPlugin(name string) Plugin {
	return gCfg.plugins[name]
}
