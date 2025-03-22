package models

import (
	"blogx_server/models/enum"
	"time"
)

type UserModel struct {
	Model
	Username       string                  `gorm:"size:32" json:"username"`
	Nickname       string                  `gorm:"size:32" son:"nickname"`
	Avatar         string                  `gorm:"size:256" json:"avatar"`
	Abstract       string                  `gorm:"size:256" json:"abstract"`
	RegisterSource enum.RegisterSourceType `json:"registerSource"`
	Password       string                  `gorm:"size:64" json:"-"`
	Email          string                  `gorm:"size:256" json:"email"`
	OpenID         string                  `gorm:"size:64" json:"openID"` //第三方登录的唯一ID
	Role           enum.RoleType           `json:"role"`                  //角色1admin2user3visitor
	UserConfModel  *UserConfModel          `gorm:"foreignKey:UserID" json:"-"`
	//UserConfModel  *UserConfModel          `gorm:"foreignKey:UserID" json:"-"`
	//CodeAge        int                     `json:"codeAge"`
	//LikeTags       []string `gorm:"type:longtext;serializer:json" json:"likeTags"` //兴趣标签
}

type UserConfModel struct {
	UserID             uint       `gorm:"unique" json:"userID"`
	UserModel          UserModel  `gorm:"foreignKey:UserID" json:"-"`
	LikeTags           []string   `gorm:"type:longtext;serializer:json" json:"likeTags"`
	UpdateUsernameDate *time.Time `json:"updateUsernameDate"` // 上次修改用户名的时间
	OpenCollect        bool       `json:"openCollect"`        // 公开我的收藏
	OpenFollow         bool       `json:"openFollow"`         // 公开我的关注
	OpenFans           bool       `json:"openFans"`           // 公开我的粉丝
	HomeStyleID        uint       `json:"homeStyleID"`        // 主页样式的id
}
