package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/service/log_service"
)

func main() {
	flags.Parse() //参数解析
	//fmt.Println(flags.FlagOptions)
	global.Conifg = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()

	log := log_service.NewRuntimeLog("同步文章数据", log_service.RuntimeDataHour)
	log.SetItem("文章1", 11)
	log.Save()
	log.SetItem("文章2", 12)
	log.Save()
	log.SetItem("文章3", 13)
	log.Save()

}
