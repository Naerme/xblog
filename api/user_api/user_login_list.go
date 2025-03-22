package user_api

import (
	"blogx_server/common"
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type UserLoginListRequest struct {
	common.PageInfo
	UserID    uint   `form:"userID"`
	Ip        string `form:"ip"`
	Addr      string `form:"addr"`
	StartTime string `form:"startTime"`
	EndTime   string `form:"endTime"`
	Type      int8   `form:"type" binding:"required,oneof=1 2"` //1用户2管理员
}
type UserLoginListResponse struct {
	models.UserLoginModel
	UserNickname string `json:"userNickname,omitempty"`
	UserAvatar   string `json:"userAvatar,omitempty"`
}

func (UserApi) UserLoginListView(c *gin.Context) {
	var cr UserLoginListRequest
	err := c.ShouldBindQuery(&cr)
	fmt.Println("cr:", cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	claims := jwts.GetClaims(c)
	fmt.Println("cr.Type:", cr.Type)
	if cr.Type == 1 {
		cr.UserID = claims.UserID
		fmt.Println("执行type=1， claims.Role:", claims.Role, "cr.UserID:", cr.UserID)
	}
	var query = global.DB.Where("")
	if cr.StartTime != "" {
		_, err := time.Parse("2006-01-02 15:04:05", cr.StartTime)
		if err != nil {
			res.FailWithMsg("开始时间格式错误", c)
			return
		}
		query.Where("created_at >= ?", cr.StartTime)
	}
	if cr.EndTime != "" {
		_, err := time.Parse("2006-01-02 15:04:05", cr.EndTime)
		if err != nil {
			res.FailWithMsg("结束间格式错误", c)
			return
		}
		query.Where("created_at <= ?", cr.EndTime)
	}
	var preloads []string
	if cr.Type == 2 && claims.Role == 2 {
		res.FailWithMsg("权限不足", c)
		return
	}

	if cr.Type == 2 && claims.Role == 1 {
		preloads = []string{"UserModel"}
		fmt.Println("执行type=2， claims.Role:", claims.Role)
	}
	fmt.Println("claims.Role:", claims.Role)
	_list, count, _ := common.ListQuery(models.UserLoginModel{
		UserID: cr.UserID,
		IP:     cr.Ip,
		Addr:   cr.Addr,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Where:    query,
		Preloads: preloads,
	})
	var list = make([]UserLoginListResponse, 0)
	for _, model := range _list {
		list = append(list, UserLoginListResponse{
			UserLoginModel: model,
			UserNickname:   model.UserModel.Nickname,
			UserAvatar:     model.UserModel.Avatar,
		})
	}
	res.OkWithList(list, count, c)
}
