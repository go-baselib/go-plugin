package repo

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-baselib/go-plugin/internal/entity"

	"gorm.io/gorm"
)

type PluginRepo struct {
	db   *gorm.DB
	once *sync.Once
}

func NewPlugin(db *gorm.DB) *PluginRepo {
	return &PluginRepo{db: db, once: &sync.Once{}}
}

func (p *PluginRepo) Store(ctx context.Context, e entity.Plugin) {
	var err error
	p.once.Do(func() {
		if err = p.db.AutoMigrate(&entity.Plugin{}); err != nil {
			fmt.Printf("auto migrate model plugin failed: %+v\n", err)
		}
	})

	if err = p.db.Model(&entity.Plugin{}).Create(&e).Error; err != nil {
		fmt.Printf("save plugin exec info failed: %+v\n", err)
	}
}
