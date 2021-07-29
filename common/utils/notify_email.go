package utils

import (
	"fmt"
	"net/smtp"
	"strconv"
)

func SendNotifyEmaill(to_email string, amount float64, WorD int8)  {
	host := "hwsmtp.exmail.qq.com"
	port := 465
	email := "guoxue@hzchainup.com"
	password := "ChainUp.com123"
	toEmail := to_email
	header := make(map[string]string)
	header["From"] = "链上星球" + "<" + email + ">"
	header["To"] = toEmail
	header["Subject"] = "充值提现通知"
	header["Content-Type"] = "text/html; charset=UTF-8"
	var body string
	// 代表充值，
	if WorD == 0 {
		body = "您的充值已经到账，充值金额为: " +strconv.FormatFloat(amount,'g',1,64) + "若有问题，请去 APP 帮助反馈里面联系链上星球客服"
	}
	//代表提现
	if  WorD ==1  {
		body = "您的提现已经到账，提现金额为: " + strconv.FormatFloat(amount,'g',1,64) + "若有问题，请去 APP 帮助反馈里面联系链上星球客服"
	}
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
