package models


import (
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)


func init() {
	mysqlConfig, _ := beego.AppConfig.GetSection("mysql")
	dburl := mysqlConfig["db_user"] + ":" + mysqlConfig["db_pass"] + "@tcp(" + mysqlConfig["db_host"] + ":" + mysqlConfig["db_port"] + ")/" + mysqlConfig["db_name"] + "?charset=utf8&loc=Asia%2FShanghai"
	if err := orm.RegisterDataBase(mysqlConfig["db_alias"], mysqlConfig["db_type"], dburl); err != nil {
		panic(errors.Wrap(err, "register data base model"))
	}
	orm.RegisterModel(new(User))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	err := orm.RunSyncdb(mysqlConfig["db_alias"], false, true)
	if err != nil {
		logs.Error(err.Error())
	}
}


type BaseModel struct {
	IsRemoved      int8       `orm:"default(0);index"`  // 0: 正常，1: 删除
	CreatedAt      time.Time  `orm:"auto_now_add;type(datetime);index" json:"created_at"`
	UpdatedAt      time.Time  `orm:"auto_now_add;type(datetime);index" json:"updated_at"`
}
