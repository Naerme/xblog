package user_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"github.com/gin-gonic/gin"
)

//type UserBaseInfoRequest struct {
//}

type UserBaseInfoResponse struct {
	UserID       uint   `json:"userID"`
	CodeAge      int    `json:"codeAge"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	LookCount    int    `json:"lookCount"`
	ArticleCount int    `json:"articleCount"`
	FansCount    int    `json:"fansCount"`
	FollowCount  int    `json:"followCount"`
	Place        string `json:"place"` // ip归属地
	//OpenCollect  bool                       `json:"openCollect"` // 公开我的收藏
	//OpenFollow   bool                       `json:"openFollow"`  // 公开我的关注
	//OpenFans     bool                       `json:"openFans"`    // 公开我的粉丝
	//HomeStyleID  uint                       `json:"homeStyleID"` // 主页样式的id
	//Relation     relationship_enum.Relation `json:"relation"`    // 与登录人的关系

}

func (UserApi) UserBaseInfoView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, cr.ID).Error
	if err != nil {
		res.FailWithMsg("不存在的用户", c)
		return
	}
	data := UserBaseInfoResponse{
		UserID:       user.ID,
		CodeAge:      user.CodeAge(),
		Avatar:       user.Avatar,
		Nickname:     user.Nickname,
		LookCount:    1,
		ArticleCount: 1,
		FansCount:    1,
		FollowCount:  1,
		Place:        user.Addr,
	}
	res.OkWithData(data, c)
}
