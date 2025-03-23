package email_service

import (
	"blogx_server/global"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)

func SendRegisterCode(to string, code string) error {
	em := global.Conifg.Email
	subject := fmt.Sprintf("[%s]账号注册", em.SendNickname)
	text := fmt.Sprintf("账号注册的验证码%s，十分钟内有效", code)

	return SendEmail(to, subject, text)
}

func SendResetPwdCode(to string, code string) error {
	em := global.Conifg.Email
	subject := fmt.Sprintf("[%s]重置密码", em.SendNickname)
	text := fmt.Sprintf("账号重置的验证码%s，十分钟内有效", code)

	return SendEmail(to, subject, text)
}

func SendEmail(to, subject, text string) (err1 error) {
	em := global.Conifg.Email

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", em.SendNickname, em.SendEmail)
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	err1 = e.Send(fmt.Sprintf("%s:%d", em.Domain, em.Port), smtp.PlainAuth("", em.SendEmail, em.AuthCode, em.Domain))
	if err1 != nil && !strings.Contains(err1.Error(), "short response:") {
		return err1
	}
	return nil
}

// 绑定邮箱
func SendBIndEmailCode(to string, code string) error {
	em := global.Conifg.Email
	subject := fmt.Sprintf("【%s】绑定邮箱", em.SendNickname)
	text := fmt.Sprintf("你正在进行账号邮箱绑定操作，这是你的验证码%s ，十分钟内有效", code)
	return SendEmail(to, subject, text)
}
