package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/wallet"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
)

type UserWalletController struct {
	beego.Controller
}

// GetWalletAssetList @Title GetAssetList
// @Description 获取资产列表 GetAssetList
// @Success 200 status bool, data interface{}, msg string
// @router /get_wallet_asset_list [post]
func (uw *UserWalletController) GetWalletAssetList() {
	ass_list, err := (&models.Asset{}).GetAssetList()
	if err != nil {
		uw.Data["json"] = RetResource(false, types.GetAssetListError, nil, err.Error())
		uw.ServeJSON()
		return
	}
	asset_list_rep := []wallet.AssetListRep{}
	var fee string
	var min_amount string
	for _, value := range ass_list {
		if value.Name == "USDT" {
			fee = beego.AppConfig.String("usdt_fee")
			min_amount = beego.AppConfig.String("min_usdt")
		}
		if value.Name == "BTC" {
			fee = beego.AppConfig.String("btc_fee")
			min_amount = beego.AppConfig.String("min_btc")
		}
		return_data := wallet.AssetListRep{
			Id: value.Id,
			AssetName: value.Name,
			ChainName: value.ChainName,
			Fee: fee,
			MinWithAmount: min_amount,
		}
		asset_list_rep = append(asset_list_rep, return_data)
	}
	uw.Data["json"] = RetResource(true, types.ReturnSuccess, asset_list_rep, "获取资产列表成功")
	uw.ServeJSON()
	return
}

// UserWalletDeposit @Title UserWalletDeposit
// @Description 充值接口 UserWalletDeposit
// @Success 200 status bool, data interface{}, msg string
// @router /wallet_deposit [post]
func (uw *UserWalletController) UserWalletDeposit() {
	bearerToken := uw.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		uw.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		uw.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user, err := models.GetUserByToken(token)
	if err != nil {
		uw.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		uw.ServeJSON()
		return
	}
	var user_w wallet.UsrWalletDepositReq
	if err := json.Unmarshal(uw.Ctx.Input.RequestBody, &user_w); err != nil {
		uw.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		uw.ServeJSON()
		return
	} else {
		if code, err := user_w.ParamCheck(); err != nil {
			uw.Data["json"] = RetResource(false, code, nil, err.Error())
			uw.ServeJSON()
			return
		}
		var asset models.Asset
		asset.Id = user_w.AssetId
		ast, err := asset.GetAssetById()
		if err != nil {
			uw.Data["json"] = RetResource(false, types.AssetNameIsEmpty, nil, "暂时不支持这个资产")
			uw.ServeJSON()
			return
		}
		usr_wallet, err := models.GetUserWalletByUserId(user.Id, user_w.AssetId)
		if err != nil {
			uw.Data["json"] = RetResource(false, types.NothisWallet, nil, "暂时不支持这个资产")
			uw.ServeJSON()
			return
		}
		deposit_data := wallet.RetDeposit{
			Id: usr_wallet.Id,
			UserId: usr_wallet.UserId,
			AssetName: ast.Name,
			ChainName: usr_wallet.ChainName,
			Address: usr_wallet.Address,
		}
		uw.Data["json"] = RetResource(true, types.ReturnSuccess, deposit_data, "获取用户钱包成功")
		uw.ServeJSON()
		return
	}
}

// WalletWithdrawAmount @Title WalletWithdrawAmount
// @Description 可提取的币种数量 WalletWithdrawAmount
// @Success 200 status bool, data interface{}, msg string
// @router /wallet_withdraw_amount [post]
func (uw *UserWalletController) WalletWithdrawAmount() {
	bearerToken := uw.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		uw.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		uw.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user, err := models.GetUserByToken(token)
	if err != nil {
		uw.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		uw.ServeJSON()
		return
	}
	var user_w wallet.UsrWalletDepositReq
	if err := json.Unmarshal(uw.Ctx.Input.RequestBody, &user_w); err != nil {
		uw.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		uw.ServeJSON()
		return
	} else {
		var userwallet models.UserWallet
		if code, err := user_w.ParamCheck(); err != nil {
			uw.Data["json"] = RetResource(false, code, nil, err.Error())
			uw.ServeJSON()
			return
		}
		var asset models.Asset
		asset.Id = user_w.AssetId
		ast, err := asset.GetAssetById()
		if err != nil {
			uw.Data["json"] = RetResource(false, types.AssetNameIsEmpty, nil, "暂时不支持这个资产")
			uw.ServeJSON()
			return
		}
		userwallet.UserId = user.Id
		userwallet.AssetId = ast.Id
		wallet_u, _, err := userwallet.GetUserWalletByUser()
		if err != nil {
			uw.Data["json"] = RetResource(false, types.NothisWallet, nil, "暂时不支持这个资产")
			uw.ServeJSON()
			return
		}
		deposit_data := wallet.RetWithdrawAmount{
			Id: wallet_u.Id,
			UserId: wallet_u.UserId,
			AssetName: ast.Name,
			Amount: wallet_u.Balance,
		}
		uw.Data["json"] = RetResource(true, types.ReturnSuccess, deposit_data, "获取可提现金额成功")
		uw.ServeJSON()
		return
	}
}

// WalletAssetWithdraw @Title WalletAssetWithdraw
// @Description 提现接口 WalletAssetWithdraw
// @Success 200 status bool, data interface{}, msg string
// @router /wallet_withdraw [post]
func (uw *UserWalletController) WalletAssetWithdraw() {
	bearerToken := uw.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		uw.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		uw.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user, err := models.GetUserByToken(token)
	if err != nil {
		uw.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		uw.ServeJSON()
		return
	}
	var withdrawdata wallet.WalletWithDrawReq
	if err := json.Unmarshal(uw.Ctx.Input.RequestBody, &withdrawdata); err != nil {
		uw.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		uw.ServeJSON()
		return
	} else {
		var userwallet models.UserWallet
		if code, err := withdrawdata.ParamCheck(); err != nil {
			uw.Data["json"] = RetResource(false, code, nil, err.Error())
			uw.ServeJSON()
			return
		}
		var asset models.Asset
		asset.Id = withdrawdata.AssetId
		ast, err := asset.GetAssetById()
		if err != nil {
			uw.Data["json"] = RetResource(false, types.AssetNameIsEmpty, nil, "暂时不支持这个资产")
			uw.ServeJSON()
			return
		}
		if ast.Name == "BTC" {
			config_amount := beego.AppConfig.String("min_btc")
			float_config_amount, _ := strconv.ParseFloat(config_amount, 64)
			if withdrawdata.Amount <= float_config_amount {
				uw.Data["json"] = RetResource(false, types.WithdrawAmountLessMin, nil, "BTC 提币数量小于最小提币限制")
				uw.ServeJSON()
				return
			}
		}
		if ast.Name == "USDT" {
			min_usdt_amount := beego.AppConfig.String("min_usdt")
			float_usdt_amount, _ := strconv.ParseFloat(min_usdt_amount, 64)
			if withdrawdata.Amount <= float_usdt_amount {
				uw.Data["json"] = RetResource(false, types.WithdrawAmountLessMin, nil, "USDT 提币数量小于最小提币限制")
				uw.ServeJSON()
				return
			}
		}
		userwallet.UserId = user.Id
		userwallet.AssetId = ast.Id
		wallet_us, _, err := userwallet.GetUserWalletByUser()
		if err != nil {
			uw.Data["json"] = RetResource(false, types.NothisWallet, nil, "没有这个钱包，请联系客服处理")
			uw.ServeJSON()
			return
		}
		if withdrawdata.Amount > wallet_us.Balance {
			uw.Data["json"] = RetResource(false, types.WalletBalanceNotEnough, nil, "钱包余额不足")
			uw.ServeJSON()
			return
		}
		if ast.Name == "BTC" {
			btc_fee, _ := beego.AppConfig.Float("btc_fee")
			if withdrawdata.Fee < btc_fee {
				uw.Data["json"] = RetResource(false, types.TxFeeNotEnough, nil, "提币手续费太小")
				uw.ServeJSON()
				return
			}
		} else if ast.Name == "USDT" {
			usdt_fee, _ := beego.AppConfig.Float("usdt_fee")
			if withdrawdata.Fee < usdt_fee {
				uw.Data["json"] = RetResource(false, types.TxFeeNotEnough, nil, "提币手续费太小")
				uw.ServeJSON()
				return
			}
		}
		withdraw := models.WalletRecord{
			AssetId: ast.Id,
			UserId: user.Id,
			ChainName: withdrawdata.ChainName,
			ToAddress: withdrawdata.ToAddress,
			Amount: withdrawdata.Amount,
			TxFee: withdrawdata.Fee,
			Status: 0,
			Comment:withdrawdata.Comment,
		}
		err = withdraw.Insert()
		var user_wallet models.UserWallet
		user_wallet.UserId = user.Id
		user_wallet.AssetId = ast.Id
		u_wallet, code, err := user_wallet.GetUserWalletByUser()
		if err != nil {
			uw.Data["json"] = RetResource(false, code, nil, err.Error())
			uw.ServeJSON()
			return
		}
		u_wallet.Balance = u_wallet.Balance - withdrawdata.Amount
		err = u_wallet.Update()
		logs.Info(err)
		uw.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "提币成功")
		uw.ServeJSON()
		return
	}
}

// WalletWDRecord @Title WalletWDRecord
// @Description 提现记录 WalletWDRecord
// @Success 200 status bool, data interface{}, msg string
// @router /wallet_record [post]
func (uw *UserWalletController) WalletWDRecord() {
	bearerToken := uw.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		uw.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		uw.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user, err := models.GetUserByToken(token)
	if err != nil {
		uw.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		uw.ServeJSON()
		return
	}
	var wallet_record_req wallet.WalletRecordListReq
	var wallet_record_m models.WalletRecord
	if err := json.Unmarshal(uw.Ctx.Input.RequestBody, &wallet_record_req); err != nil {
		uw.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		uw.ServeJSON()
		return
	} else {
		wallet_record_m.Status = wallet_record_req.Status
		wallet_record_m.UserId = user.Id
		wallet_record_m.WOrD = wallet_record_req.IsWD
		withdraw, code, err := wallet_record_m.GetWalletWithdrawList(wallet_record_req.AssetName, int64(wallet_record_req.Page), int64(wallet_record_req.PageSize))
		if err != nil {
			uw.Data["json"] = RetResource(false, code, nil, "获取提币记录失败")
			uw.ServeJSON()
			return
		}
		var withdraw_rlist []wallet.WithdrawRecordRep
		for _, value := range withdraw {
			var assets models.Asset
			assets.Id = value.AssetId
			ast, _ := assets.GetAssetById()
			withdraw_r := wallet.WithdrawRecordRep{
				Id:          value.Id,
				AssetName:   ast.Name,
				ChainName:   value.ChainName,
				FromAddress: value.FromAddress,
				ToAddress:   value.ToAddress,
				Amount:      value.Amount,
				TxHash:      value.TxHash,
				TxFee:       value.TxFee,
				TransFee:    value.TxFee,
				Comment:     value.Comment,
				WOrD:        value.WOrD,
				Status:      value.Status,    // 0:审核中(未锁定)；1:交易中 2:已发出 3:成功 4:失败 5:锁定未审核 6:审核通过 7:审核拒绝
				IsRemoved:   value.IsRemoved, // 0: 正常，1: 删除
				CreatedAt:   value.CreatedAt,
				UpdatedAt:   value.UpdatedAt,
			}
			withdraw_rlist = append(withdraw_rlist, withdraw_r)
		}
		uw.Data["json"] = RetResource(true, types.ReturnSuccess, withdraw_rlist, "获取提币记录成功")
		uw.ServeJSON()
		return
	}
}


// WalletFundAsset @Title WalletFundAsset
// @Description 钱包资产 WalletFundAsset
// @Success 200 status bool, data interface{}, msg string
// @router /wallet_fund_asset [post]
func (uw *UserWalletController) WalletFundAsset() {

}

