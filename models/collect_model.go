package models

type CollectModel struct {
	Model
	Title       string                    `gorm:"size:32" json:"title"`
	Abstract    string                    `gorm:"size:256" json:"abstract"`
	Cover       string                    `gorm:"size:256" json:"cover"`
	ArticleList []UserArticleCollectModel `gorm:"foreignKey:CollectID" json:"-"` //属于哪一个收藏夹
	UserID      uint                      `json:"userID"`
	UserModel   UserModel                 `gorm:"foreignKey:UserID" json:"-"`
	IsDefault   bool                      `json:"isDefault"` //是否是默认收藏夹
}
