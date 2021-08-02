package models


import (
	"blockshop/common"
	"blockshop/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"time"
)

type UserAddress struct {
	Id        int64     `json:"id"`
	UserId    int64     `orm:"default(150000)" json:"user_id"`
	UserName  string    `orm:"size(128);index" json:"user_name"` // 收件名字
	Phone     string    `orm:"size(32);index" json:"phone"`      // 手机号
	Address   string    `orm:"size(512);index" json:"address"`   // 地址
	IsSet     int8      `orm:"index" json:"is_set"`              // 0: 正常，1: 默认地址
	Status    int8      `orm:"index" json:"status"`              // 0: 正常，1: 删除
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);index" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime);index" json:"updated_at"`
}

func (this *UserAddress) SearchField() []string{
	return []string{"username","phone","user_id"}
}

func (this *UserAddress) TableName() string {
	return common.TableName("user_address")
}

func (this *UserAddress) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *UserAddress) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *UserAddress) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func (this *UserAddress) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *UserAddress) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *UserAddress) GetGoodsAddress() (*UserAddress, int64, string) {
	var address UserAddress
	if err := orm.NewOrm().QueryTable(this.TableName()).Filter("UserId", this.UserId).Filter("IsSet", 1).OrderBy("-id").Limit(1).RelatedSel().One(&address); err != nil {
		return nil, types.SystemDbErr, "数据库查询失败，请联系客服处理"
	}
	return &address, types.ReturnSuccess, "获取地址成功"
}

func (this *UserAddress) GetAddressById() (*UserAddress, int64, string) {
	var address UserAddress
	if err := orm.NewOrm().QueryTable(this.TableName()).Filter("Id", this.Id).RelatedSel().One(&address); err != nil {
		return nil, types.SystemDbErr, "数据库查询失败，请联系客服处理"
	}
	return &address, types.ReturnSuccess, ""
}

func (this *UserAddress) UpdateAddressInfo() bool {
	var crfr_addr UserAddress
	crfr_addr.Id = this.Id
	addres, _, _ := crfr_addr.GetAddressById()
	addres.UserName = this.UserName
	addres.Phone = this.Phone
	addres.Address = this.Address
	addres.IsSet = this.IsSet
	err := addres.Update()
	if err != nil {
		return true
	} else {
		return true
	}
}

func (this *UserAddress) GetUserAddressList() ([]*UserAddress, int64, string) {
	var address_list []*UserAddress
	if _, err := orm.NewOrm().QueryTable(this.TableName()).Filter("UserId", this.UserId).All(&address_list); err != nil {
		return nil, types.SystemDbErr, "数据库查询失败，请联系客服处理"
	}
	return address_list, types.ReturnSuccess, "获取地址成功"
}


func GetUserAddressDefault(user_id int64) (*UserAddress, int, error) {
	address := UserAddress{}
	if err := orm.NewOrm().QueryTable(UserAddress{}).Filter("UserId", user_id).Filter("IsSet", 1).Limit(1).One(&address); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &address, types.ReturnSuccess, nil
}