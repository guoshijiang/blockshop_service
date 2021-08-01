package user

import (
	"blockshop/types"
	"github.com/pkg/errors"
)

type Register struct {
	UserName        string `json:"user_name"`         // 用户名
	Password        string `json:"password"`          // 密码
	ConfirmPassword string `json:"confirm_password"`  // 确认密码
	PinCode         string `json:"pin_code"`          // pin 码
	PublicKey       string `json:"public_key"`        // 用户公钥
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
	Factor    string `json:"factor"`      // 登录因子
	TimeUnit  int    `json:"time_unit"`   // 0: 分钟， 1: 小时， 2: 天
}

func (this Login) ParamCheck() (int, error) {
	if this.UserName == "" || this.Password == "" || this.LoginTime == 0 || this.TimeUnit == 0 {
		return types.ParamEmptyError, errors.New("注册请求参数为空，请检查之后再注册")
	}
	return types.ReturnSuccess, nil
}

type TwoFaReq struct {
	UserName  string `json:"user_name"`
}

func (this TwoFaReq) ParamCheck() (int, error) {
	if this.UserName == "" {
		return types.ParamEmptyError, errors.New("用户名为空")
	}
	return types.ReturnSuccess, nil
}

type OpenCloseTwoFa struct {
	UserId    int64  `json:"user_id"`
	IsOpen    int8   `json:"is_open"`    // 0:关闭；1:打开
	PublicKey string `json:"public_key"` // 关闭的时候公钥请传空
}

func (this OpenCloseTwoFa) ParamCheck() (int, error) {
	if this.UserId <= 0 {
		return types.ParamEmptyError, errors.New("无效的用户ID")
	}
	return types.ReturnSuccess, nil
}

type UpdatePasswordReq struct {
	UserId       int64  `json:"user_id"`
	OldPassword  string `json:"old_password"`
	NewPassword  string `json:"new_password"`
	CNewPassword string `json:"c_new_password"`
}

func (this UpdatePasswordReq) ParamCheck() (int, error) {
	if this.UserId <= 0 {
		return types.ParamEmptyError, errors.New("无效的用户ID")
	}
	if this.NewPassword != this.CNewPassword {
		return types.PasswordError, errors.New("两次输入的密码不一样")
	}
	return types.ReturnSuccess, nil
}

type ForgetPasswordReq struct {
	UserId       int64  `json:"user_id"`
	PinCode      string `json:"pin_code"`
	NewPassword  string `json:"new_password"`
	CNewPassword string `json:"c_new_password"`
}

func (this ForgetPasswordReq) ParamCheck() (int, error) {
	if this.UserId <= 0 {
		return types.ParamEmptyError, errors.New("无效的用户ID")
	}
	if this.NewPassword != this.CNewPassword {
		return types.PasswordError, errors.New("两次输入的密码不一样")
	}
	return types.ReturnSuccess, nil
}