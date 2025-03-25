package es_service

import (
	"blogx_server/global"
	"context"
	"github.com/sirupsen/logrus"
)

func CreateIndexV2(index, mapping string) {
	if ExistsIndex(index) {
		DeleteIndex(index)
	}
	CreateIndex(index, mapping)
}

func CreateIndex(index, mapping string) {
	_, err := global.ESClient.
		CreateIndex(index).
		BodyString(mapping).Do(context.Background())
	if err != nil {
		logrus.Errorf("%s 索引创建失败%s", index, err)
		return
	}
	logrus.Infof("%s 索引创建成功", index)
}

// ExistsIndex 判断索引是否存在
func ExistsIndex(index string) bool {
	exists, _ := global.ESClient.IndexExists(index).Do(context.Background())
	return exists
}

//func DocCreate() {
//	user := models.UserModel{
//		ID:        12,
//		UserName:  "lisi",
//		Age:       23,
//		NickName:  "夜空中最亮的lisi",
//		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
//		Title:     "今天天气很不错",
//	}
//	indexResponse, err := global.ESClient.Index().Index(user.Index()).BodyJson(user).Do(context.Background())
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Printf("%#v\n", indexResponse)
//}

func DeleteIndex(index string) {
	_, err := global.ESClient.
		DeleteIndex(index).Do(context.Background())
	if err != nil {
		logrus.Errorf("%s 索引删除失败%s", index, err)
		return
	}
	logrus.Infof("%s 索引删除成功", index)
}
