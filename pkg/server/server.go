package server

import (
	"context"

	bp "github.com/go-baselib/go-plugin"
	"github.com/go-baselib/go-plugin/internal/usecase"

	proto "github.com/grpcprotocol/plugin"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type server struct {
	usecase.Plugin
	proto.UnsafePluginServer
}

type GRPCPlugin struct {
	Impl usecase.PluginAPI
	plugin.Plugin
}

func NewGRPCPlugin(impl usecase.PluginAPI) *GRPCPlugin {
	return &GRPCPlugin{Impl: impl}
}

type GRPCServer struct {
	Impl usecase.PluginAPI
}

type GRPCClient struct {
	proto.PluginClient
}

func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterPluginServer(s, server{p.Impl, proto.UnimplementedPluginServer{}})
	return nil
}

func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{proto.NewPluginClient(c)}, nil
}

func Run(name string, up usecase.Plugin) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: bp.Handshake,
		Plugins: map[string]plugin.Plugin{
			name: &GRPCPlugin{Impl: up},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
