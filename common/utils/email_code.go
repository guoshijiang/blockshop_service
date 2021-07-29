package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
	"log"
	"net"
	"net/smtp"
	"strconv"
)


func SendEmail(email string, verify_code int) int {
	host_name := beego.AppConfig.String("host_name")
	eamil_addr := beego.AppConfig.String("eamil_addr")
	email_name := beego.AppConfig.String("email_name")
	email_passwd := beego.AppConfig.String("email_passwd")
	auth := smtp.PlainAuth("", email_name, email_passwd, host_name)
	to := []string{email}
	str := fmt.Sprintf("From:%s\r\nTo:,%s,\r\nSubject:verify code is\r\n\r\n %d \r\n", email_name, email, verify_code)
	msg := []byte(str)
	err := smtp.SendMail(eamil_addr, auth, email_name, to, msg)
	if err != nil {
		logs.Error("邮件发送失败", err)
	}
	return verify_code
}


func SendEmailByBeego(email string, verify_code int) int {
	config := `{"username":"guoxue@hzchainup.com","password":"ChainUp.com123","host":"smtp.exmail.qq.com","port":25}`
	temail := utils.NewEMail(config)
	temail.To = []string{email}
	temail.From = "guoxue@hzchainup.com"
	temail.Subject = "验证码"
	temail.HTML = "【链上星球】您的邮箱验证码是"+strconv.Itoa(verify_code) + "请勿泄露，否则会造成资产安全"
	err := temail.Send()
	if err != nil {
		logs.Error(err)
		return 0
	}
	return 0
}

func SendSSLEmaill(to_email string, verify_code int)  {
	host := "hwsmtp.exmail.qq.com"
	port := 465
	email := "guoshijiang@gingernet.vip"
	password := "Guo20123762"
	toEmail := to_email
	header := make(map[string]string)
	header["From"] = "来自市集" + "<" + email + ">"
	header["To"] = toEmail
	header["Subject"] = "验证码"
	header["Content-Type"] = "text/html; charset=UTF-8"
	body := "【市集商城】您的邮箱验证码是" +strconv.Itoa(verify_code) + "请勿泄露，否则会造成资产安全"
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


func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}


func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {
	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}