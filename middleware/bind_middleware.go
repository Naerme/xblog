package middleware

import (
	"blogx_server/common/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func BindJsonMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindJSON(&cr)
	fmt.Println(cr)
	fmt.Println(err)

	if err != nil {
		fmt.Println("进入err")
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("request", cr)
	fmt.Println("c:", c)
	return
}

func BindQueryMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("request", cr)
	return
}

func GetBind[T any](c *gin.Context) (cr T) {
	return c.MustGet("request").(T)
}
