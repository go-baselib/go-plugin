package main

import (
	"github.com/go-baselib/go-plugin/pkg/hi"
	"github.com/go-baselib/go-plugin/pkg/server"
)

func main() {
	server.Run("hi", &hi.Hi{})
}
