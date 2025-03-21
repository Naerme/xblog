package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/service/email_service"
)

func main() {

	flags.Parse() //参数解析
	//fmt.Println(flags.FlagOptions)
	global.Conifg = core.ReadConf()
	core.InitLogrus()

	email_service.SendRegisterCode("hujunyin1016@163.com", "2333")
}
