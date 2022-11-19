package main

import (
	"github.com/go-baselib/go-plugin/config"
	"github.com/go-baselib/go-plugin/internal/app"
	_ "github.com/go-baselib/go-plugin/pkg/hi"
)

func main() {
	config.Init()
	app.Run(config.GetConfig())
}
