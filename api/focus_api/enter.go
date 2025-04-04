package focus_api

import (
	"blogx_server/common"
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type FocusApi struct {
}
type FocusUserRequest struct {
	FocusUserID uint `json:"focusUserID" binding:"required"`
}

// FocusUserApi 登录人关注用户
func (FocusApi) FocusUserView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserRequest](c)

	claims := jwts.GetClaims(c)
	if cr.FocusUserID == claims.UserID {
		res.FailWithMsg("你时刻都在关注自己", c)
		return
	}
	// 查关注的用户是否存在
	var user models.UserModel
	err := global.DB.Take(&user, cr.FocusUserID).Error
	if err != nil {
		res.FailWithMsg("关注用户不存在", c)
		return
	}

	// 查之前是否已经关注过他了
	var focus models.UserFocusModel
	err = global.DB.Take(&focus, "user_id = ? and focus_user_id = ?", claims.UserID, user.ID).Error
	if err == nil {
		res.FailWithMsg("请勿重复关注", c)
		return
	}

	// 每天关注是不是应该有个限度？
	// 每天的取关也要有个限度？

	// 关注
	global.DB.Create(&models.UserFocusModel{
		UserID:      claims.UserID,
		FocusUserID: cr.FocusUserID,
	})

	res.OkWithMsg("关注成功", c)
	return
}

type FocusUserListRequest struct {
	common.PageInfo
	FocusUserID uint `form:"focusUserID"`
	UserID      uint `form:"userID"` // 查用户的关注
}
type FocusUserListResponse struct {
	FocusUserID       uint      `json:"focusUserID"`
	FocusUserNickname string    `json:"focusUserNickname"`
	FocusUserAvatar   string    `json:"focusUserAvatar"`
	FocusUserAbstract string    `json:"focusUserAbstract"`
	CreatedAt         time.Time `json:"createdAt"`
}

// FocusUserListView 我的关注和用户的关注
func (FocusApi) FocusUserListView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserListRequest](c)

	if cr.UserID != 0 {
		// 传了用户id，我就查这个人关注的用户列表
		var userConf models.UserConfModel
		fmt.Println("cr.UserID:", cr.UserID)
		err := global.DB.Take(&userConf, "user_id = ?", cr.UserID).Error
		if err != nil {
			res.FailWithMsg("用户配置信息不存在", c)
			return
		}
		if !userConf.OpenFollow {
			res.FailWithMsg("此用户未公开我的关注", c)
			return
		}
	} else {
		claims, err := jwts.ParseTokenByGin(c)
		if err != nil || claims == nil {
			res.FailWithMsg("请登录", c)
			return
		}
		cr.UserID = claims.UserID
	}
	fmt.Println("cr.UserID:", cr.UserID)

	_list, count, _ := common.ListQuery(models.UserFocusModel{
		FocusUserID: cr.FocusUserID,
		UserID:      cr.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Preloads: []string{"FocusUserModel"},
	})

	var list = make([]FocusUserListResponse, 0)
	for _, model := range _list {
		list = append(list, FocusUserListResponse{
			FocusUserID:       model.FocusUserID,
			FocusUserNickname: model.FocusUserModel.Nickname,
			FocusUserAvatar:   model.FocusUserModel.Avatar,
			FocusUserAbstract: model.FocusUserModel.Abstract,
			CreatedAt:         model.CreatedAt,
		})
	}

	res.OkWithList(list, count, c)
}

type FansUserListResponse struct {
	FansUserID       uint      `json:"fansUserID"`
	FansUserNickname string    `json:"fansUserNickname"`
	FansUserAvatar   string    `json:"fansUserAvatar"`
	FansUserAbstract string    `json:"fansUserAbstract"`
	CreatedAt        time.Time `json:"createdAt"`
}

// FansUserListView 我的粉丝和用户的粉丝
func (FocusApi) FansUserListView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserListRequest](c)

	if cr.UserID != 0 {
		// 传了用户id，我就查这个人的粉丝列表
		var userConf models.UserConfModel
		err := global.DB.Take(&userConf, "user_id = ?", cr.UserID).Error
		if err != nil {
			res.FailWithMsg("用户配置信息不存在", c)
			return
		}
		if !userConf.OpenFans {
			res.FailWithMsg("此用户未公开我的粉丝", c)
			return
		}
	} else {
		claims, err := jwts.ParseTokenByGin(c)
		if err != nil || claims == nil {
			res.FailWithMsg("请登录", c)
			return
		}
		cr.UserID = claims.UserID
	}

	_list, count, _ := common.ListQuery(models.UserFocusModel{
		FocusUserID: cr.UserID,
		UserID:      cr.FocusUserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Preloads: []string{"UserModel"},
	})

	var list = make([]FansUserListResponse, 0)
	for _, model := range _list {
		list = append(list, FansUserListResponse{
			FansUserID:       model.UserID,
			FansUserNickname: model.UserModel.Nickname,
			FansUserAvatar:   model.UserModel.Avatar,
			FansUserAbstract: model.UserModel.Abstract,
			CreatedAt:        model.CreatedAt,
		})
	}

	res.OkWithList(list, count, c)
}

// UnFocusUserView 登录人取关用户
func (FocusApi) UnFocusUserView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserRequest](c)

	claims := jwts.GetClaims(c)
	if cr.FocusUserID == claims.UserID {
		res.FailWithMsg("你无法取关自己", c)
		return
	}
	// 查关注的用户是否存在
	var user models.UserModel
	err := global.DB.Take(&user, cr.FocusUserID).Error
	if err != nil {
		res.FailWithMsg("取关用户不存在", c)
		return
	}

	// 查之前是否已经关注过他了
	var focus models.UserFocusModel
	err = global.DB.Take(&focus, "user_id = ? and focus_user_id = ?", claims.UserID, user.ID).Error
	if err != nil {
		res.FailWithMsg("未关注此用户", c)
		return
	}
	// 每天的取关也要有个限度？
	// 取关
	global.DB.Delete(&focus)
	res.OkWithMsg("取消关注成功", c)
	return
}
