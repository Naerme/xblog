package global

import (
	"blogx_server/conf"
	"gorm.io/gorm"
)

var (
	Conifg *conf.Config
	DB     *gorm.DB
)
