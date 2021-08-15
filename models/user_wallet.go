package models

import (
	"blockshop/common"
	"blockshop/http"
	"blockshop/types"
	"blockshop/types/wallet"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type UserWallet struct {
	BaseModel
	Id          int64     `orm:"pk;column(id);auto;size(11)" description:"钱包ID" json:"id"`
	UserId      int64     `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
	AssetId     int64     `orm:"column(asset_id);index" description:"资产ID" json:"asset_id"`
	ChainName   string    `orm:"column(chain_name);default(Bitcoin)" description:"链的名称" json:"chain_name"`
	Address     string    `orm:"column(address);size(256)" description:"地址" json:"address"`
	Balance     float64   `orm:"column(balance);default(150);digits(22);decimals(8)" description:"钱包余额" json:"balance"`
	InAmount    float64   `orm:"column(in_amount);default(150);digits(22);decimals(8)" description:"收入统计" json:"in_amount"`
	OutAmount   float64   `orm:"column(out_amount);default(150);digits(22);decimals(8)" description:"支出统计" json:"out_amount"`
}

func (this *UserWallet) TableName() string {
	return common.TableName("user_wallet")
}

func (this *UserWallet) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *UserWallet) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *UserWallet) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *UserWallet) SearchField() []string {
  return []string{"chain"}
}


func (w *UserWallet) GetUserWalletListByUserId() ([]*UserWallet, int, error) {
	var u_wallet []*UserWallet
	_, err := orm.NewOrm().QueryTable(w.TableName()).Filter("user_id", w.UserId).RelatedSel().All(&u_wallet)
	if err != nil {
		return nil, types.NothisWallet, err
	}
	return u_wallet, types.ReturnSuccess, nil
}

func (w *UserWallet) GetUserWalletByUser() (*UserWallet, int, error) {
	err := w.Query().Filter("user_id", w.UserId).Filter("asset_id", w.AssetId).One(w)
	if err != nil {
		return nil, types.UserIdEmptyError, err
	}
	return w, types.ReturnSuccess, nil
}

func GetUserWalletByUserId(user_id, asset_id int64) (*UserWallet, error) {
	var u_wallet UserWallet
	err := orm.NewOrm().QueryTable(UserWallet{}).Filter("user_id", user_id).Filter("asset_id", asset_id).One(&u_wallet)
	if err != nil {
		return nil, errors.New("获取钱包地址失败")
	}
	return &u_wallet, nil
}

func GetUserWalletBalance(asset_id, user_id int64) (float64, int, error) {
	var user_wallet UserWallet
	err := orm.NewOrm().QueryTable(&UserWallet{}).Filter("asset_id", asset_id).Filter("user_id", user_id).RelatedSel().One(&user_wallet)
	if err != nil {
		return 0, types.NothisWallet, err
	}
	return float64(user_wallet.Balance), types.ReturnSuccess, nil
}


func UpdateWalletBalance(db orm.Ormer, asset_id, user_id int64, pay_amount float64) (bool, int, error) {
	var user_wallet UserWallet
	err := db.QueryTable(&UserWallet{}).Filter("asset_id", asset_id).Filter("user_id", user_id).RelatedSel().One(&user_wallet)
	if err != nil {
		return false, types.NothisWallet, err
	}
	sql := fmt.Sprintf(`
	UPDATE
		%s
	SET
		balance = balance - %f
	WHERE
		user_id = '%d' AND asset_id = %d`, user_wallet.TableName(), pay_amount, user_id, asset_id)
	if _, err := db.Raw(sql).Exec(); err != nil {
		return false, types.NothisWallet, err
	}
	return true, types.ReturnSuccess, nil
}


func CreateWalletAddress(user_id, wallet_id int64) error {
	wallet_url := beego.AppConfig.String("wallet_url")
	request_url := wallet_url + "create_address"
	data := wallet.AddressReq{
		UserId:   user_id,
		WalletId: wallet_id,
	}
	response := http.HttpPost(request_url, data, "application/json")
	var addres_rep wallet.WalletAddressRep
	if err := json.Unmarshal([]byte(response), &addres_rep); err != nil {
		logs.Error("decode json fail")
		return err
	}
	if addres_rep.Status == false {
		request_url := wallet_url + "get_address"
		response := http.HttpPost(request_url, data, "application/json")
		if err := json.Unmarshal([]byte(response), &addres_rep); err != nil {
			logs.Error("decode json fail")
			return err
		}
	}
	for _, value := range addres_rep.Data {
		var asset Asset
		var user_w UserWallet
		asset.Name = value.AssetName
		assts, err := asset.GetAssetByName()
		if err != nil {
			logs.Error("get asset by name fail")
			return err
		}
		err = orm.NewOrm().QueryTable(UserWallet{}).Filter("user_id", user_id).Filter("asset_id", assts.Id).One(&user_w)
		if err != nil {
			logs.Error("get wallet id fail")
			return err
		}
		var chain_name string
		if value.ChainName == "Ethereum" && value.AssetName == "USDT" {
			chain_name = "Erc20"
		}
		if value.ChainName == "TRX" && value.AssetName == "USDT" {
			chain_name = "Trc20"
		}
		if value.ChainName == "Bitcoin" && value.AssetName == "BTC" {
			chain_name = "Bitcoin"
		}
		wallet_address := UserWallet{
			UserId:  user_id,
			ChainName: chain_name,
			AssetId: assts.Id,
			Address: value.Address,
		}
		err1 := wallet_address.Insert()
		if err1 != nil {
			logs.Error("insert wallet fail")
			return err
		}
	}
	return nil
}

func GeneratedUserWallet(user_id int64) int {
	var asset Asset
	asset_list, _ := asset.GetAssetList()
	for _, v := range asset_list {
		uw := new(UserWallet)
		uw.AssetId = v.Id
		uw.UserId = user_id
		err := uw.Insert()
		if err != nil {
			return types.CreateWalletFail
		}
	}
	err := CreateWalletAddress(user_id, 100)
	if err != nil {
		return types.CreateWalletFail
	}
	return types.ReturnSuccess
}

//获得钱包信息
func (w *UserWallet) GetUserWalletAsset() (*UserWallet, error) {
	var user_wallet UserWallet
	err := orm.NewOrm().QueryTable(&UserWallet{}).Filter("asset_id", w.AssetId).Filter("user_id", w.UserId).RelatedSel().One(&user_wallet)
	if err != nil {
		return nil, err
	}
	return &user_wallet, nil
}

//获得特定钱包的信息
func (w *UserWallet) GetUserWallet(condition *orm.Condition) (float64, error) {
	var user_wallet UserWallet
	err := orm.NewOrm().QueryTable(&UserWallet{}).SetCond(condition).One(&user_wallet)
	if err != nil {
		return 0, err
	}
	return user_wallet.Balance, nil
}
