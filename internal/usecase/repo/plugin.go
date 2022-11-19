package repo

import (
	"context"

	"github.com/go-baselib/go-plugin/internal/entity"
)

type PluginRepo struct {
}

func NewPlugin() *PluginRepo {
	return &PluginRepo{}
}

func (p *PluginRepo) Store(ctx context.Context, e entity.Plugin) {

}
