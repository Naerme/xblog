package user_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/utils/mps"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AdminInfoUpdateRequest struct {
	UserId   uint           `json:"userID" binding:"required"`
	Nickname *string        `json:"nickname" s-u:"nickname"`
	Username *string        `json:"username" s-u:"username"`
	Avatar   *string        `json:"avatar" s-u:"avatar"`
	Abstract *string        `json:"abstract" s-u:"abstract"`
	Role     *enum.RoleType `json:"role" s-u:"role"`
}

func (UserApi) AdminInfoUpdateView(c *gin.Context) {
	var cr AdminInfoUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	userMap := mps.StructToMap(cr, "s-u")
	var user models.UserModel
	err = global.DB.Take(&user, cr.UserId).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	err = global.DB.Model(&user).Updates(userMap).Error
	if err != nil {
		fmt.Println("usermap:")
		res.FailWithMsg("用户信息修改失败", c)
		return
	}

	res.OkWithMsg("用户信息修改成功", c)
}
