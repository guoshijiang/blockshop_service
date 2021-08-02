package user

import (
	"blockshop/types"
	"github.com/pkg/errors"
)

type Register struct {
	UserName string `json:"user_name"` // 用户名
	Password string `json:"password"`  // 密码
	ConfirmPassword string `json:"confirm_password"`  // 确认密码
	PinCode string `json:"pin_code"`   // pin 码
}

func (this Register) ParamCheck() (int, error) {
	if this.UserName == "" || this.Password == "" || this.ConfirmPassword == "" || this.PinCode == ""{
		return types.ParamEmptyError, errors.New("注册请求参数为空，请检查之后再注册")
	}
	if this.Password != this.ConfirmPassword {
		return types.PasswordNotEqual, errors.New("密码和确认密码密码不一样，请核对之后再注册")
	}
	return types.ReturnSuccess, nil
}

type Login struct {
	UserName  string `json:"user_name"`   // 用户名
	Password  string `json:"password"`    // 密码
	LoginTime int64  `json:"login_time"`  // 登录时长
	TimeUnit  int    `json:"time_unit"`   // 0: 分钟， 1: 小时， 2: 天
}

func (this Login) ParamCheck() (int, error) {
	if this.UserName == "" || this.Password == "" || this.LoginTime == 0 || this.TimeUnit == 0 {
		return types.ParamEmptyError, errors.New("注册请求参数为空，请检查之后再注册")
	}
	return types.ReturnSuccess, nil
}

type ReqUserRegister struct {
  UserName string   `json:"user_name"`
  Password1  string `json:"password1"`
  Password2  string `json:"password2"`
}