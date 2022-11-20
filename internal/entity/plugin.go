package entity

import (
	"gorm.io/gorm"
)

type Plugin struct {
	gorm.Model
	Req string `gorm:"column:f_req"`
	Rsp string `gorm:"column:f_rsp"`
	Err string `gorm:"column:f_err"`
}
