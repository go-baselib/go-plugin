package api

import (
	"context"
	"fmt"
	"os/exec"

	bp "github.com/go-baselib/go-plugin"
	"github.com/go-baselib/go-plugin/pkg/server"

	"github.com/grpcprotocol/plugin"
	hp "github.com/hashicorp/go-plugin"
)

type PluginAPI struct {
}

func NewPlugin() *PluginAPI {
	return &PluginAPI{}
}

func (p *PluginAPI) Exec(ctx context.Context, in *plugin.ExecReq) (*plugin.ExecRsp, error) {
	var typeURL = in.GetPayload().GetTypeUrl()
	var pl, ok = bp.GetPlugin(typeURL)
	if !ok {
		return nil, fmt.Errorf("TypeURL:%s 没有可处理插件", in.GetPayload().GetTypeUrl())
	}

	var name = bp.GetName(typeURL)
	var client = hp.NewClient(&hp.ClientConfig{
		HandshakeConfig:  bp.Handshake,
		Plugins:          hp.PluginSet{name: pl},
		Cmd:              exec.Command("sh", "-c", "./bin/"+name),
		AllowedProtocols: []hp.Protocol{hp.ProtocolGRPC},
	})
	defer client.Kill()

	var rpcClient, err = client.Client()
	if err != nil {
		return nil, err
	}

	var raw interface{}
	if raw, err = rpcClient.Dispense(name); err != nil {
		fmt.Printf("Dispense err: %+v\n", err)
		return nil, err
	}

	var out *plugin.ExecRsp
	if out, err = raw.(*server.GRPCClient).Exec(ctx, in); err != nil {
		fmt.Printf("Exec err: %+v\n", err)
		return nil, err
	}

	return out, nil
}
