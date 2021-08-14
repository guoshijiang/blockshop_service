package models

import (
	"blockshop/common"
	"blockshop/types"
	type_merchant "blockshop/types/merchant"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"time"
)

type Merchant struct {
	BaseModel
	Id              int64     `orm:"pk;column(id);auto;size(11)" description:"商品属性ID"  json:"id"`
	UserId          int64     `orm:"column(user_id);size(64);index" description:"用户ID" json:"user_id"`
	Logo            string    `orm:"column(logo);size(150);default(/static/upload/default/user-default-60x60.png)" description:"商家Logo" json:"logo"`
	MerchantName    string    `orm:"column(merchant_name);size(512);index" description:"商家名称" json:"merchant_name"`
	MerchantIntro   string    `orm:"column(merchant_intro);size(512);index" description:"商家简介" json:"merchant_intro"`
	MerchantDetail  string    `orm:"column(merchant_detail);type(text)" description:"商家详情" json:"merchant_detail"`
	ContactUser     string    `orm:"column(contact_user);size(128);index" description:"商家联系人" json:"contact_user"`
	Phone           string    `orm:"column(phone);size(64);index" description:"商家联系电话" json:"phone"`
	WeChat          string    `orm:"column(we_chat);size(64);index" description:"商家联系微信" json:"we_chat"`
	Address         string    `orm:"column(address);size(512);index" description:"店铺地址" json:"address"`
	GoodsNum        int64     `orm:"column(goods_num)" description:"商品总数" json:"goods_num"`
	MerchantWay     int8      `orm:"column(merchant_way);default(0);index" description:"商家类别" json:"merchant_way"`   // 0:自营商家； 1:认证商家  2:普通商家
	SettlePercent   float64   `orm:"column(settle_percent);default(0);digits(22);decimals(8)" description:"结算比例"  json:"settle_percent"`
	ShopLevel       int8      `orm:"column(shop_level)" description:"店铺等级" json:"shop_level"`
	ShopServer      int8      `orm:"column(shop_server)" description:"店铺服务" json:"shop_server"`
	ShopScore       int64  	  `orm:"column(shop_score)" description:"店铺评分" json:"shop_score"`
	MonthSellNum    int64  	  `orm:"column(month_sell_num)" description:"本月销售量" json:"month_sell_num"`
	MonthSellAmount float64   `orm:"column(month_sell_amount)" description:"本月销售总额" json:"month_sell_amount"`
    TotalSellNum    int64     `orm:"column(total_sell_num)" description:"本月销售量" json:"total_sell_num"`
	TotalSellAmount float64	  `orm:"column(total_sell_amount)" description:"本月销售量" json:"total_sell_amount"`
	IsShow          int8      `orm:"column(is_show);default(0)" description:"是否在首页展示" json:"is_show"` // 0:不展示，1:展示
}

func (this *Merchant) TableName() string {
	return common.TableName("merchant")
}

func (this *Merchant) SearchField() []string {
  return []string{"merchant_name"}
}

func (this *Merchant) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func GetMerchantList(page, pageSize int, is_show int8) ([]*Merchant, int64, error) {
	offset := (page - 1) * pageSize
	merchant_list := make([]*Merchant, 0)
	query := orm.NewOrm().QueryTable(Merchant{})
	if is_show == 1 {
		query = query.Filter("is_show", is_show)
	}
	total, _ := query.Count()
	_, err := query.OrderBy("-GoodsNum").Limit(pageSize, offset).All(&merchant_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return merchant_list, total, nil
}

func GetMerchantDetail(id int64) (*Merchant, int, error) {
	var merchant Merchant
	if err := orm.NewOrm().QueryTable(Merchant{}).Filter("Id", id).RelatedSel().One(&merchant); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &merchant, types.ReturnSuccess, nil
}

func OpenMerchant(open_mct type_merchant.OpenMerchantReq) (msg string, err error, id int64) {
	db := orm.NewOrm()
	if err := db.Begin(); err != nil {
		err := errors.Wrap(err, "开启支付事物失败")
		return "开启支付事物失败", err, 0
	}
	defer func() {
		if err != nil {
			err = errors.Wrap(db.Rollback(), "回滚事物失败")
		} else {
			err = errors.Wrap(db.Commit(), "提交事物失败")
		}
	}()
	merchant_config := MerchantConfig{}
	if err = db.QueryTable(merchant_config.TableName()).Filter("config_type", 0).One(&merchant_config); err != nil {
		err := errors.New("查询配置失败")
		return "查询配置失败", err, 0
	}
	user := User{}
	if err = db.QueryTable(user).RelatedSel().Filter("id", open_mct.UserId).One(&user); err != nil {
		err = errors.New("查询开通商家用户失败")
		return "查询开通商家用户失败", err, 0
	}
	var pay_asset *Asset
	var pay_coin_aount float64
	if open_mct.PayWay == PayWayUSDT {
		var ast Asset
		ast.Name = "USDT"
		asst, _ := ast.GetAsset()
		pay_asset = asst
		total, _, err := GetUserWalletBalance(asst.Id, open_mct.UserId)
		if err != nil {
			err = errors.New("用户没有钱包")
			return "用户没有钱包", err, 0
		}
		if total < merchant_config.UsdtAmount {
			err = errors.New("账户没有足够的资金，请去充值")
			return "账户没有足够的资金，请去充值", err, 0
		}
		pay_coin_aount = merchant_config.UsdtAmount
	} else {
		var ast Asset
		ast.Name = "BTC"
		asst, _ := ast.GetAsset()
		pay_asset = asst
		total, _, err := GetUserWalletBalance(asst.Id, open_mct.UserId)
		if err != nil {
			err = errors.New("用户没有钱包")
			return "用户没有钱包", err, 0
		}
		if total < merchant_config.BtcAmount {
			err = errors.New("账户没有足够的资金，请去充值")
			return "账户没有足够的资金，请去充值", err, 0
		}
		pay_coin_aount = merchant_config.BtcAmount
	}
	_, _, err = UpdateWalletBalance(db, pay_asset.Id, open_mct.UserId, pay_coin_aount)
	if err != nil {
		err = errors.New("更新钱包余额失败")
		return "更新钱包余额失败", err, 0
	}
	mct := Merchant{
		UserId: open_mct.UserId,
		Logo: open_mct.MctLogo,
		MerchantName: open_mct.MctName,
		MerchantIntro: open_mct.MctAbstruct,
		MerchantDetail: open_mct.MctDetail,
		ContactUser: open_mct.MctCrtName,
		Phone:open_mct.MctCrtPhone,
	}
	err, id = mct.Insert()
	if err != nil {
		err = errors.New("新增商家信息失败")
		return "新增商家信息失败", err, 0
	}
	now := time.Now()
	mct_record := MerchantOpenRecord{
		UserId: open_mct.UserId,
		MerchantId: id,
		PayCoinAmount: pay_coin_aount,
		PayWay: open_mct.PayWay,
		PayAt: &now,
	}
	err, _ = mct_record.Insert()
	if err != nil {
		err = errors.New("新增商家信息失败")
		return "新增商家信息失败", err, 0
	}
	return "", nil, id
}

func GetMerchantByUserId(user_id int64) (*Merchant, int, error) {
	var merchant Merchant
	if err := orm.NewOrm().QueryTable(Merchant{}).Filter("user_id", user_id).RelatedSel().One(&merchant); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &merchant, types.ReturnSuccess, nil
}