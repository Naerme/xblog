package flag_user

import (
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/utils/pwd"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type FlagUser struct {
}

func (FlagUser) Create() {
	var role enum.RoleType
	fmt.Println("选择角色 1超级管理员 2用户 3访客")
	_, err := fmt.Scan(&role)
	if err != nil {
		logrus.Errorf("输入错误", err)
		return
	}
	if !(role == 1 || role == 2 || role == 3) {
		logrus.Errorf("输入角色错误", err)
		return
	}
	var username string
	fmt.Println("请输入用户名：")
	fmt.Scan(&username)
	//查用户名是否存在
	var model models.UserModel
	err = global.DB.Take(&model, "username = ?", username).Error
	if err == nil {
		logrus.Errorf("此用户名已存在")
	}

	fmt.Println("请输入密码")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd())) // 读取用户输入的密码
	if err != nil {
		fmt.Println("读取密码时出错:", err)
		return
	}
	fmt.Println("请再次输入密码")
	repassword, err := terminal.ReadPassword(int(os.Stdin.Fd())) // 读取用户输入的密码
	if err != nil {
		fmt.Println("读取密码时出错:", err)
		return
	}
	if string(password) != string(repassword) {
		fmt.Println("两次密码不一致")
		return
	}
	hashPwd, _ := pwd.GenerateFromPassword(string(password))
	//创建用户
	err = global.DB.Create(&models.UserModel{
		Username:       username,
		Nickname:       "yong",
		RegisterSource: enum.RegisterTerminalSourceType,
		Password:       hashPwd,
		Role:           role,
	}).Error
	if err != nil {
		fmt.Println("创建用户失败", err)
	}
	logrus.Infof("创建成功")
}
