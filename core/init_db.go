package core

import (
	"blogx_server/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

func InitDB() *gorm.DB {
	if len(global.Conifg.DB) == 0 {
		logrus.Fatalf("未配置数据库")
	}

	dc := global.Conifg.DB[0]

	// TODO: pgsql的支持
	db, err := gorm.Open(mysql.Open(dc.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成外键约束
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败 %s", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logrus.Infof("数据库连接成功！")
	fmt.Println("len(global.Conifg.DB) :", len(global.Conifg.DB))
	if len(global.Conifg.DB) > 1 {
		// 读写库不为空，就注册读写分离的配置
		var readList []gorm.Dialector
		i := 0
		for _, d := range global.Conifg.DB[1:] {
			fmt.Println(i)
			i++
			readList = append(readList, mysql.Open(d.DSN()))
		}
		fmt.Println("readList", readList)
		err = db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dc.DSN())}, // 写
			Replicas: readList,                               // 读
			Policy:   dbresolver.RandomPolicy{},
		}))
		if err != nil {
			logrus.Fatalf("读写配置错误 %s", err)
		}
	}
	return db

}
