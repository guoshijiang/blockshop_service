package wallet

import (
	"blockshop/types"
	"github.com/pkg/errors"
)

type WalletWithdrawDepositNotify struct {
	WOrD        int8    `json:"w_or_d"`       // 0: 代表充值，1：代表提现
	WalletId    int64   `json:"wallet_id"`    // 用户钱包 ID
	UserUuid    string  `json:"user_uuid"`    // 用户 UUid
	AssetName   string  `json:"asset_name"`   // 资产名称
	ChainName   string  `json:"chain_name"`   // 链名称 erc20 / trc20
	FromAddress string  `json:"from_address"` // 转出地址
	ToAddress   string  `json:"to_address"`   // 转入地址
	Amount      float64 `json:"amount"`       // 提币金额
	TxHash      string  `json:"tx_hash"`      // 交易 Hash
	TxFee       float64 `json:"tx_fee"`       // 链上提币手续费
	Status      int8    `json:"status"`       // 上报状态
}

func (wwdn WalletWithdrawDepositNotify) WalletWithdrawDepositNotifyParamValidate() (int, error) {
	if wwdn.WOrD != 0 && wwdn.WOrD != 1 {
		return types.InvalidWOrD, errors.New("无效的充值或者提现类型")
	}
	if wwdn.AssetName == "" {
		return types.AssetNameIsEmpty, errors.New("资产名称为空")
	}
	if wwdn.ChainName == "" {
		return types.ChainNameIsEmpty, errors.New("链名称不能为空")
	}
	if wwdn.Amount <= 0 {
		return types.AmountLessZero, errors.New("提现或者充值金额小于0")
	}
	if wwdn.TxHash == "" {
		return types.TxHashEmpty, errors.New("交易hash不能为空")
	}
	if wwdn.WOrD == 1 {
		if wwdn.TxFee <= 0 {
			return types.TxFeeLessZero, errors.New("交易手续费小于 0")
		}
		if wwdn.FromAddress != "" || wwdn.ToAddress != "" {
			return types.AddressIsEmpty, errors.New("转出或者转入地址为空")
		}
	}
	if wwdn.WOrD == 0 {
		if wwdn.ToAddress == "" {
			return types.AddressIsEmpty, errors.New("转入地址为空")
		}
	}
	return types.ReturnSuccess, nil
}

type WalletAddressReq struct {
	UserUuid  string `json:"user_uuid"`
	AssetName string `json:"asset_name"`
	WalletId  int64  `json:"wallet_id"`
	ChainName string `json:"chain_name"`
}

type WalletWithdrawReq struct {
	AssetName string `json:"asset_name"`
	ChainName string `json:"chain_name"`
	UserUuid  string `json:"user_uuid"`
	WithdrawId int64 `json:"withdraw_id"`
	FromAddr string `json:"from_addr"`
	ToAddr string `json:"to_addr"`
	Amount float64 `json:"amount"`
	TransFee float64 `json:"trans_fee"`
}


type AddressReq struct {
	UserId  int64 `json:"user_id"`
	WalletId  int64 `json:"wallet_id"`
}


type AddressRep struct {
	AssetName string `json:"asset_name"`
	ChainName string `json:"chain_name"`
	Address   string   `json:"address"`
}


type WalletAddressRep struct {
	Status bool         `json:"status"`
	Code   int64        `json:"code"`
	Msg    string       `json:"msg"`
	Data []*AddressRep  `json:"data"`
}
