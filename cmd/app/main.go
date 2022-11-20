package main

import (
	_ "embed"

	"github.com/go-baselib/go-plugin/config"
	"github.com/go-baselib/go-plugin/internal/app"
	_ "github.com/go-baselib/go-plugin/pkg/hi"
)

//go:embed app.yaml
var appConf []byte

func main() {
	config.Init(appConf)
	app.Run(config.GetConfig())
}
