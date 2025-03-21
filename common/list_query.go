package common

import (
	"blogx_server/global"
	"fmt"
	"gorm.io/gorm"
)

type PageInfo struct {
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Order string `form:"order"` //前端可以覆盖
}

func (p *PageInfo) GetPage() int {
	if p.Page > 20 || p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PageInfo) GetLimit() int {
	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *PageInfo) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

type Options struct {
	PageInfo     PageInfo
	Likes        []string
	Preloads     []string
	Where        *gorm.DB
	Debug        bool
	DefaultOrder string
}

func ListQuery[T any](model T, option Options) (list []T, count int, err error) {
	//基础查询
	query := global.DB.Model(model).Where(model)

	//日志
	if option.Debug {
		query = query.Debug()
	}

	//预加载
	for _, preload := range option.Preloads {
		query = query.Preload(preload)
	}

	//模糊匹配
	if len(option.Preloads) > 0 && option.PageInfo.Key != "" {
		likes := global.DB.Where("")

		//for _, column := range option.Likes {
		//	likes.Or(
		//		fmt.Sprintf("%s like %s", column),
		//		fmt.Sprintf("%%%s%%", option.PageInfo.Key))
		//}

		for _, column := range option.Likes {
			likes.Or(
				fmt.Sprintf("%s like ?", column),
				fmt.Sprintf("%%%s%%", option.PageInfo.Key))
		}
		query = query.Where(likes)
	}
	//高级查询
	if option.Where != nil {
		query = query.Where(option.Where)

	}

	//定制化
	if option.Where != nil {
		query.Where(option.Where)
	}

	//查总数
	var _c int64
	query.Count(&_c)
	count = int(_c)

	//分页
	limit := option.PageInfo.GetLimit()
	offset := option.PageInfo.GetOffset()

	//排序
	if option.PageInfo.Order != "" {
		query = query.Order(option.PageInfo.Order)
	} else {
		if option.DefaultOrder != "" {
			query = query.Order(option.DefaultOrder)

		}
	}

	err = query.Offset(offset).Limit(limit).Find(&list).Error
	return
}
