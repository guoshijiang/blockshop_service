package models

import (
	"blockshop/common"
	"blockshop/types"
	"blockshop/types/user"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	BaseModel
	Id             int64         `orm:"pk;column(id);auto;size(11)" description:"用户ID" json:"id"`
	UserName       string        `orm:"column(user_name);size(128)" description:"用户名称" json:"user_name"`
	Avator         string        `orm:"column(avator);size(150);default(/static/upload/default/user-default-60x60.png)" description:"用户头像" json:"avator"`
	Password       string        `orm:"column(password);size(128)" description:"密码" json:"password"`
	PinCode        string        `orm:"column(pin_code);size(128)" description:"提款的pin码" json:"pin_code"`
	FundPassword   string        `orm:"column(fund_password);size(128)" description:"钱包资金密码" json:"fund_password"`
	LoginCount     int64         `orm:"column(login_count);default(0);index" description:"登录次数" json:"login_count"`
	Token          string        `orm:"column(token);size(128)" description:"用户 token" json:"token"`
	IsMerchant     int8          `orm:"column(is_merchant);default(0);index" description:"是否开通商户" json:"is_merchant"` // ：0 不是，1: 是
	MemberLevel    int8          `orm:"column(member_level);default(1);index" description:"会员级别" json:"member_level"`  // 0:v0:普通会员 1:V1:白银会员，2:V2:白金会员，3:V3:黄金会员; 4:V4:砖石会有; 5:V5:皇冠会员
	UserPrivateKey string		 `orm:"column(user_private_key);size(512)" description:"用户的私钥" json:"user_private_key"`
	UserPublicKey  string        `orm:"column(user_public_key);size(512)" description:"用户的公钥" json:"user_public_key"`
	Factor         string        `orm:"column(factor);size(512)" description:"用户因子" json:"factor"`
	IsOpen         int8          `orm:"column(is_open);default(0);index" description:"是否开因子认证" json:"is_open"` // ：0 不是，1: 是
	ActiveTime     int64         `orm:"column(active_time);default(0);index" description:"用户活跃时间" json:"active_time"`
}

type OptionList struct {
  Id							int						`json:"id"`
  Name						string					`json:"name"`
}

func (this *User) TableName() string {
	return common.TableName("user")
}

func (this *User) SearchField() []string {
	return []string{"user_name"}
}

func (this *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *User) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *User) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func GetUserByToken(token string) (*User, error) {
	u := User{}
	if err := orm.NewOrm().QueryTable(u.TableName()).RelatedSel().Filter("token", token).One(&u); err != nil {
		return nil, errors.Wrap(err, "error in GetUserByToken")
	}
	return &u, nil
}

func (this *User) ExistByName (user_name string) bool {
	return orm.NewOrm().QueryTable(this).Filter("user_name", user_name).Exist()
}

func UpdateFactor(user_id int64, factor string)(code int, msg string){
	var fa2_user User
	err := fa2_user.Query().Filter("id", user_id).One(fa2_user)
	if err != nil {
		return types.UserNoExist, "没有这个用户"
	}
	fa2_user.Factor = factor
	err = fa2_user.Update()
	if err != nil {
		return types.UserNoExist, "更新factor失败"
	}
	return types.ReturnSuccess, "更新factor成功"
}

func GetUserInfo(user_name string) (user *User) {
	var login_user User
	err := login_user.Query().Filter("user_name", user_name).OrderBy("-id").One(&login_user)
	if err != nil {
		return nil
	}
	return &login_user
}

func UserRegister(register user.Register) (code int, msg string) {
	u := User{}
	if u.ExistByName(register.UserName) {
		return types.UserExist, "该用户名已经被注册"
	}
	is_open := int8(0)
	if register.PublicKey != "" {
		is_open = 1
	}
	user_data := User {
		UserName: register.UserName,
		Password: common.ShaOne(register.Password),
		Token: uuid.NewV4().String(),
		PinCode: register.PinCode,
		UserPublicKey: register.PublicKey,
		IsOpen: is_open,
	}
	err, _ := user_data.Insert()
	if err != nil {
		return types.SystemDbErr, "创建用户失败"
	}
	return types.ReturnSuccess,  "用户注册成功"
}

func UserLogin(login user.Login) (login_rep *user.LoginRep, code int, msg string) {
	var login_user User
	err := login_user.Query().Filter("user_name", login.UserName).Limit(1).One(&login_user)
	if err != nil {
		return nil, types.UserNoExist, "你还没有注册，请去注册"
	}
	if login_user.IsOpen == 1 {
		if login.Factor != login_user.Factor {
			return nil, types.FactorIsNotSame, "验证因子出错"
		}
	}
	return &user.LoginRep{
		Id:       login_user.Id,
		UserName: login_user.UserName,
		Token:    login_user.Token,
	}, types.ReturnSuccess, "登录成功"
}

func OpenCloseFactor(is_open int8, user_id int64, public_key string)(code int, msg string) {
	var fa2_user User
	err := fa2_user.Query().Filter("id", user_id).OrderBy("-id").One(&fa2_user)
	if err != nil {
		return types.UserNoExist, "没有这个用户"
	}
	fa2_user.UserPublicKey = public_key
	fa2_user.IsOpen = is_open
	err = fa2_user.Update()
	if err != nil {
		return types.UserNoExist, "开通或者关闭双因子认证失败"
	}
	return types.ReturnSuccess, "开通或者关闭双因子认证成功"
}

func UpdatePassword(upd_pwd user.UpdatePasswordReq) (code int, msg string) {
	var upd_user User
	err := upd_user.Query().Filter("id", upd_pwd.UserId).One(&upd_user)
	if err != nil {
		return types.UserNoExist, "没有这个用户"
	}
	if upd_user.Password != common.ShaOne(upd_pwd.OldPassword) {
		return types.PasswordError, "输入的原密码错误"
	}
	upd_user.Password = upd_pwd.NewPassword
	err = upd_user.Update()
	if err != nil {
		return types.UserNoExist, "修改密码失败"
	}
	return types.ReturnSuccess, "修改密码成功"
}

func ForgetPassword(fpt_pwd user.ForgetPasswordReq) (code int, msg string) {
	var fpt_user User
	err := fpt_user.Query().Filter("id", fpt_pwd.UserId).One(&fpt_user)
	if err != nil {
		return types.UserNoExist, "没有这个用户"
	}
	if fpt_user.PinCode != fpt_pwd.PinCode {
		return types.PasswordError, "输入的Pin码错误"
	}
	fpt_user.Password = fpt_pwd.NewPassword
	err = fpt_user.Update()
	if err != nil {
		return types.UserNoExist, "修改密码失败"
	}
	return types.ReturnSuccess, "修改密码成功"
}

