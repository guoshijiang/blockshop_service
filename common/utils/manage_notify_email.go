package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/smtp"
	"strconv"
)

const EmailAccount = "403628431@qq.com"

func SendManageDepositeEmaill(user string, amount float64)  {
	host := "hwsmtp.exmail.qq.com"
	port := 465
	email := "guoxue@hzchainup.com"
	password := "ChainUp.com123"
	toEmail := beego.AppConfig.String("deposit_to_eamail")
	header := make(map[string]string)
	header["From"] = "链上星球 APP" + "<" + email + ">"
	header["To"] = EmailAccount
	header["Subject"] = "通知管理员"
	header["Content-Type"] = "text/html; charset=UTF-8"
	body := "用户" + user + "充值，充值金额为: " +strconv.FormatFloat(amount,'g',1,64) + "已经到账"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	auth := smtp.PlainAuth("",  email,  password,  host)
	err := SendMailUsingTLS(fmt.Sprintf("%s:%d", host, port), auth, email, []string{toEmail}, []byte(message))
	if err != nil {
		panic(err)
	}
}