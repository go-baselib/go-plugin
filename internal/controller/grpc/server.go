package grpc

import (
	"github.com/go-baselib/go-plugin/internal/usecase"

	"github.com/grpcprotocol/plugin"
	"google.golang.org/grpc"
)

type Plugin struct {
	usecase.Plugin
	plugin.UnsafePluginServer
}

func NewServer(uc usecase.Plugin) *grpc.Server {
	var (
		s  = grpc.NewServer()
		ps = Plugin{uc, plugin.UnimplementedPluginServer{}}
	)

	plugin.RegisterPluginServer(s, ps)

	return s
}
