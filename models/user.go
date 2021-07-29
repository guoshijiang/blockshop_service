package models

import (
  "blockshop/common"
  "github.com/astaxie/beego/orm"
  "github.com/pkg/errors"
)

type User struct {
	BaseModel
	Id             int64         `json:"id"`
	Phone          string        `orm:"size(64);index" json:"phone"`
	UserName       string        `orm:"size(128)" json:"user_name"`
	Avator         string        `orm:"size(150);default(/static/upload/default/user-default-60x60.png)"`
	Password       string        `orm:"size(128)" json:"password"`
	FundPassword   string        `orm:"size(128)" json:"fund_password"`           // 钱包资金密码
	Email          string        `orm:"size(128);index" json:"email"`
	LoginCount     int64         `orm:"default(0);index" json:"login_count"`
	Token          string        `orm:"size(128)" json:"token"`
	IsAuth         int8          `orm:"default(0);index" json:"is_auth"`          // 0 未实名认证，1: 实名认证中；2:实名认证成功；3实名认证失败
	MemberLevel    int8          `orm:"default(1);index" json:"member_level"`     // 0:v0:普通会员 1:V1:白银会员，2:V2:白金会员，3:V3:黄金会员; 4:V4:砖石会有; 5:V5:皇冠会员
	MyInviteCode   string        `orm:"size(128)" json:"my_invite_code"`          // 用户自己网体邀请码
	InviteMeUserId int64         `orm:"size(64);index" json:"invite_me_user_id"`  // 网体上级用户id
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


func GetUserById(id int64) (User, error) {
	var query_user User
	err := query_user.Query().Filter("Id", id).Limit(1).One(&query_user)
	if err != nil {
		return query_user, errors.New("user is not exist")
	}
	return query_user, nil
}
