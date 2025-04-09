package core

import (
	"blogx_server/global"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func EsConnect() *elastic.Client {
	es := global.Conifg.ES
	if !es.Enable || es.Addr == "" {
		logrus.Infof("未启用es连接")
		return nil
	}
	client, err := elastic.NewClient(
		elastic.SetURL(es.Url()),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(es.Username, es.Password),
	)
	if err != nil {
		logrus.Panic("es连接失败：", err)
		fmt.Println(err)
		return nil
	}
	logrus.Infof("es连接成功：")
	return client
}
