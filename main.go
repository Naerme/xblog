package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/router"
)

func main() {

	flags.Parse() //参数解析
	//fmt.Println(flags.FlagOptions)
	global.Conifg = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	global.Redis = core.InitRedis()
	global.ESClient = core.EsConnect()

	flags.Run()
	router.Run()

}
