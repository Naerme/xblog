package middleware

import (
	"blogx_server/common/res"
	"blogx_server/models/enum"
	"blogx_server/service/redis_service/redis_jwt"
	"blogx_server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	blcType, ok := redis_jwt.HashTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(blcType.Msg(), c)
		c.Abort()
		return
	}

	c.Set("claims", claims)
	return

}

func AdminMiddleware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	blcType, ok := redis_jwt.HashTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(blcType.Msg(), c)
		c.Abort()
		return
	}
	if claims.Role != enum.AdminRole {
		res.FailWithMsg("权限错误", c)
		c.Abort()
		return
	}
	fmt.Println("claims.Role", claims.Role)
	c.Set("claims", claims)
	return
}
