package core

import (
	"blogx_server/global"
	river "blogx_server/service/river_service"
	"github.com/sirupsen/logrus"
)

func InitMysqlES() {
	if !global.Conifg.River.Enable {
		logrus.Infof("关闭musql同步操作")
		return
	}
	if !global.Conifg.ES.Enable {
		logrus.Infof("未配置es，关闭musql同步操作")
		return
	}
	r, err := river.NewRiver()
	if err != nil {
		logrus.Fatal(err)
	}
	go r.Run()
}
