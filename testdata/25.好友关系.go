package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/models"
	"fmt"
)

func main() {
	flags.Parse()
	global.Conifg = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()

	fmt.Println(models.CalcUserRelationship(3, 7))
	fmt.Println(models.CalcUserRelationship(4, 3))
	//fmt.Println(models.CalcUserRelationship(4, 1))
	//fmt.Println(models.CalcUserRelationship(5, 1))
	//fmt.Println(models.CalcUserRelationship(3, 4))

}
