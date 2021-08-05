package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"os/exec"
	"strings"
	"time"
)


func init() {
	mysqlConfig, _ := beego.AppConfig.GetSection("mysql")
	dburl := mysqlConfig["db_user"] + ":" + mysqlConfig["db_pass"] + "@tcp(" + mysqlConfig["db_host"] + ":" + mysqlConfig["db_port"] + ")/" + mysqlConfig["db_name"] + "?charset=utf8&loc=Asia%2FShanghai"
	if err := orm.RegisterDataBase(mysqlConfig["db_alias"], mysqlConfig["db_type"], dburl); err != nil {
		panic(errors.Wrap(err, "register data base model"))
	}
	orm.RegisterModel(new(User),new(AdminLog),new(AdminLogData),new(AdminUser),new(AdminRole),new(AdminMenu),new(Asset),new(Forum),new(ForumReply),new(GoodsCat),new(GoodsImage),new(GoodsOrder),new(GoodsType),new(GoodsAttr),new(Merchant),new(UserWallet),new(WalletRecord),new(Goods),new(News),new(Message),new(MerchantDataStat),new(UserDataStat),new(GoodsComment),new(ForumCat))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	err := orm.RunSyncdb(mysqlConfig["db_alias"], false, true)
	if err != nil {
		logs.Error(err.Error())
	}
	insertRole()
	insertAdmin()
	loadMenu()
}


type BaseModel struct {
	IsRemoved      int8       `orm:"default(0);index"`  // 0: 正常，1: 删除
	CreatedAt      time.Time  `orm:"auto_now_add;type(datetime);index" json:"created_at"`
	UpdatedAt      time.Time  `orm:"auto_now_add;type(datetime);index" json:"updated_at"`
}

func insertRole() {
  var err error
  _, err = orm.NewOrm().Raw("INSERT INTO `admin_role` VALUES (1, '管理员', '后台管理员角色', '1,19,20,21,22,23,24,25,53,54,76', 1);").Exec()
  _, err = orm.NewOrm().Raw("INSERT INTO `admin_role` VALUES (2, '商户A', '商户A', '1,2,18,19,20', 1);").Exec()
  _, err = orm.NewOrm().Raw("INSERT INTO `admin_role` VALUES (6, '商户B', '商户B', '1,2,18,19,20', 1);").Exec()
  fmt.Println("err---", err)
}

func insertAdmin(){
  var err error
  _, err = orm.NewOrm().Raw("INSERT INTO `admin_user` VALUES (1, 'admin', 'JDJhJDEwJFdRaU5qRlpLUmZ1dG8uUXdpaXNaaS40SkIwdXNhQmRZOTZsMmc5by53SldMUi9qTjVLc1dp', '超级管理员', '/static/uploads/attachment/aecb9fb7-871b-43fc-9414-a4265d0cb72d.png', '1', 1, 0,0);").Exec()
  _, err = orm.NewOrm().Raw("INSERT INTO `admin_user` VALUES (2, 'aaa', 'JDJhJDEwJEhHaWZ0LkdzaTRtYzRRMWNvNncxTC5HL0NEZnk5bkpJdmw1bzdiRDE2OEVSOXROamk2MWxX', 'aaa', '/static/admin/images/avatar.png', '2', 1, 0,0);").Exec()
  _, err = orm.NewOrm().Raw("INSERT INTO `admin_user` VALUES (3, 'bbb', 'JDJhJDEwJFpKVElLZVpBLjV5YXRObC5FUDdMVy5sQ1F4ekx0VjVzd3laQ0p1L05ERU1kZDlvNTFJcnhh', 'bbb', '/static/admin/images/avatar.png', '6', 1, 0,0);").Exec()
  fmt.Println("err---", err)
}

func loadMenu() {
  sqls,_ := beego.AppConfig.GetSection("source")
  fmt.Println(sqls)
  mysqlConfig, _ := beego.AppConfig.GetSection("mysql")
  command := "mysql -P {port} -h {address} -u{username} -p{password} {database} < {source}"
  command = strings.Replace(command, "{username}", mysqlConfig["db_user"], 1)
  command = strings.Replace(command, "{password}", mysqlConfig["db_pass"], 1)
  command = strings.Replace(command, "{database}", mysqlConfig["db_name"], 1)
  command = strings.Replace(command, "{address}", mysqlConfig["db_host"], 1)
  command = strings.Replace(command, "{source}", sqls["source"], 1)
  command = strings.Replace(command, "{port}", mysqlConfig["db_port"], 1)
  cmd := exec.Command("/bin/sh", "-c", command)
  out, err := cmd.CombinedOutput()
  if err != nil {
    fmt.Println("err",err)
  }
  fmt.Println(string(out))
}

type WalletExtra struct {
  UserName				string					`json:"user_name"`
  AssetName       string          `json:"asset_name"`
}

type UserWalletList struct {
  UserWallet
  WalletExtra
}

type WalletRecordList struct {
  WalletRecord
  WalletExtra
  AdminName       string          `json:"admin_name"`
}

type MessageData struct {
  Message
  UserName					string					`json:"user_name"`
}