package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	db, _ := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/mydb1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	fmt.Println(db)
}
