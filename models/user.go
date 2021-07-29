package models

import (
	"blockshop/common"
	"blockshop/types"
	"blockshop/types/user"
	"github.com/astaxie/beego/orm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	BaseModel
	Id             int64         `orm:"column(id)" description:"用户ID" json:"id"`
	UserName       string        `orm:"column(user_name);size(128)" description:"用户名称" json:"user_name"`
	Avator         string        `orm:"column(avator);size(150);default(/static/upload/default/user-default-60x60.png)" description:"用户头像" json:"avator"`
	Password       string        `orm:"column(password);size(128)" description:"密码" json:"password"`
	PinCode        string        `orm:"column(pin_code);size(128)" description:"提款的pin码" json:"pin_code"`
	FundPassword   string        `orm:"column(fund_password);size(128)" description:"钱包资金密码" json:"fund_password"`
	LoginCount     int64         `orm:"column(login_count);default(0);index" description:"登录次数" json:"login_count"`
	Token          string        `orm:"column(token);size(128)" description:"用户 token" json:"token"`
	IsMerchant     int8          `orm:"column(is_merchant);default(0);index" description:"是否开通商户" json:"is_merchant"` // ：0 不是，1: 是
	MemberLevel    int8          `orm:"column(member_level);default(1);index" description:"会员级别" json:"member_level"`  // 0:v0:普通会员 1:V1:白银会员，2:V2:白金会员，3:V3:黄金会员; 4:V4:砖石会有; 5:V5:皇冠会员
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

func (this *User) ExistByName (user_name string) bool {
	return orm.NewOrm().QueryTable(this).Filter("user_name", user_name).Exist()
}

func UserRegister(register user.Register) (code int, msg string) {
	u := User{}
	if u.ExistByName(register.UserName) {
		return types.UserExist, "该用户名已经被注册"
	}
	user_data := User {
		UserName: register.UserName,
		Password: common.ShaOne(register.Password),
		Token: uuid.NewV4().String(),
		PinCode: register.PinCode,
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
	return &user.LoginRep{
		Id:       login_user.Id,
		UserName: login_user.UserName,
		Token:    login_user.Token,
	}, types.ReturnSuccess, "登录成功"
}

