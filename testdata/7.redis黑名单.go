package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/service/redis_service/redis_jwt"
	"fmt"
)

func main() {
	flags.Parse() //参数解析
	//fmt.Println(flags.FlagOptions)
	global.Conifg = core.ReadConf()
	core.InitLogrus()
	global.Redis = core.InitRedis()

	//token, err := jwts.GetToken(jwts.Claims{
	//	UserID: 2,
	//	Role:   1,
	//})
	//fmt.Println(token, err)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjEsInJvbGUiOjEsInVzZXJuYW1lIjoiIiwiZXhwIjoxNzQxOTU4NzQzLCJpc3MiOiJoankifQ.N4OV79b1Ujk4MHZ4fDidTFVRpM8B10GyAUMRBcNLPyY"
	redis_jwt.TokenBlack(token, redis_jwt.UserBlackType)
	blk, ok := redis_jwt.HasTokenBlack(token)
	fmt.Println(blk, ok)
}
