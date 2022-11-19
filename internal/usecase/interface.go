package usecase

import (
	"context"

	"github.com/go-baselib/go-plugin/internal/entity"

	"github.com/grpcprotocol/plugin"
)

type Plugin interface {
	Exec(context.Context, *plugin.ExecReq) (*plugin.ExecRsp, error)
}

type PluginRepo interface {
	Store(context.Context, entity.Plugin)
}

type PluginAPI interface {
	Exec(context.Context, *plugin.ExecReq) (*plugin.ExecRsp, error)
}
