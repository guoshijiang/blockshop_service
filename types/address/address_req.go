package address


import (
	"blockshop/types"
	"github.com/pkg/errors"
	"regexp"
)

const (
	PhoneNumRule = "^(1[3|4|5|6|7|8|9][0-9]\\d{4,8})$"
)

type UserAddressAddCheck struct {
	UserId    int64   `json:"user_id"`
	UserName  string  `json:"user_name"`
	Phone     string  `json:"phone"`
	Address   string  `json:"address"`
	IsSet     int8    `json:"is_set"`
}

func (this UserAddressAddCheck) UserAddressAddCheckParamValidate() (int, error) {
	if this.UserId <= 0 {
		return types.UserIsNotExist, errors.New("用户不存在, 请联系客服处理")
	}
	if this.UserName == "" {
		return types.UserIsNotExist, errors.New("用户名为空，请务必填写")
	}
	if this.Phone == "" {
		return types.PhoneEmptyError, errors.New("手机号码不能为空，请务必填写")
	}
	result, _ := regexp.MatchString(PhoneNumRule, this.Phone)
	if !result {
		return types.PhoneFormatError, errors.New("手机号码格式不正确")
	}
	if this.Address == "" {
		return types.AddressIsEmpty, errors.New("地址不能为空，请务必填写")
	}
	return types.ReturnSuccess, nil
}

type UserUpdAddressCheck struct {
	AddressId  int64   `json:"address_id"`
	UserName  string  `json:"user_name"`
	Phone      string  `json:"phone"`
	Address    string  `json:"address"`
	IsSet      int8    `json:"is_set"`
}

func (this UserUpdAddressCheck) UserUpdAddressCheckParamValidate() (int, error) {
	if this.AddressId <= 0 {
		return types.AddressIdLessEqError, errors.New("地址不存在, 请联系客服处理")
	}
	if this.UserName == "" {
		return types.UserIdEmptyError, errors.New("用户名为空，请务必填写")
	}
	if this.Phone == "" {
		return types.PhoneEmptyError, errors.New("手机号码不能为空，请务必填写")
	}
	result, _ := regexp.MatchString(PhoneNumRule, this.Phone)
	if !result {
		return types.PhoneFormatError, errors.New("手机号码格式不正确")
	}
	if this.Address == "" {
		return types.AddressIsEmpty, errors.New("地址不能为空，请务必填写")
	}
	return types.ReturnSuccess, nil
}

type UserAdddressDelCheck struct {
	AddressId  int64   `json:"address_id"`
}

func (this UserAdddressDelCheck) UserAdddressDelParamValidate() (int, error) {
	if this.AddressId <= 0 {
		return types.AddressIdLessEqError, errors.New("地址不存在, 请联系客服处理")
	}
	return types.ReturnSuccess, nil
}
