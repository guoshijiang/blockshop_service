package models


import (
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
	orm.RegisterModel(new(User), new(UserInfo), new(UserWallet), new(UserIntegral), new(UserCoupon),
		new(AdminUser), new(AdminMenu), new(AdminRole), new(Goods), new(GoodsCar), new(Merchant),
		new(GoodsComment), new(GoodsCat), new(GoodsImage), new(GoodsOrder), new(OrderProcess), new(GroupOrder),
		new(GroupHelper), new(ImageFile),  new(IntegralRecord), new(IntegralTrade), new(UserAddress),
		new(Version), new(WalletRecord), new(Banner), new(CustomerService), new(Questions), new(UserAccount),
		new(AssetDebt), new(MerchantSettle), new(MerchantWallet), new(MerchantWithdraw), new(GoodsType),new(MerchantSettleAccount),new(MerchantSettleDaily))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	err := orm.RunSyncdb(mysqlConfig["db_alias"], false, true)
	if err != nil {
		logs.Error(err.Error())
	}
}
