package usecase

import (
	"context"

	"github.com/go-baselib/go-plugin/internal/entity"

	"github.com/grpcprotocol/plugin"
)

type PluginUseCase struct {
	repo PluginRepo
	api  PluginAPI
}

func NewPlugin(repo PluginRepo, api PluginAPI) Plugin {
	return &PluginUseCase{repo: repo, api: api}
}

func (p *PluginUseCase) Exec(ctx context.Context, in *plugin.ExecReq) (*plugin.ExecRsp, error) {
	var e entity.Plugin
	defer func() {
		p.repo.Store(ctx, e)
	}()

	e.Req = in.String()
	var out, err = p.api.Exec(ctx, in)
	if err != nil {
		e.Err = err.Error()
		return out, err
	}
	e.Rsp = out.String()
	return out, nil
}
