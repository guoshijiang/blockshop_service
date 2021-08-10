package wallet

import "time"

type AssetListRep struct {
	Id            int64     `json:"id"`
	AssetName     string    `json:"asset_name"`
	ChainName     string `json:"chain_name"`
	Fee           string    `json:"fee"`
	MinWithAmount string    `json:"min_with_amount"`
}

type RetDeposit struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	AssetName string `json:"asset_name"`
	ChainName string `json:"chain_name"`
	Address   string `json:"address"`
}

type RetWithdrawAmount struct {
	Id        int64   `json:"id"`
	UserId    int64   `json:"user_id"`
	AssetName string  `json:"asset_name"`
	ChainName string  `json:"chain_name"`
	Amount    float64 `json:"amount"`
}

type UserWalletFundAsset struct {
	AssetId     int64   `json:"asset_id"`
	AssetName   string  `json:"asset_name"`
	TotalAmount float64 `json:"total_amount"`
	UsdPrice    float64 `json:"usd_price"`
	CnyPrice    float64 `json:"cny_price"`
}

type AccountDetailsListRep struct {
	Total  []AccountDetails `json:"total"`
	TotalPage int64 `json:"total_page"`
}

type AccountDetails struct {
	Amount  string `json:"amount"`
	Way     string  `json:"way"`
	PayTime  string `json:"pay_time"`
	AssetName string `json:"asset_name"`
}

type WithdrawRecordRep struct {
	Id          int64     `json:"id"`
	AssetName   string    `json:"asset_name"`
	ChainName   string    `json:"chain_name"`
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	Amount      float64   `json:"amount"`
	TxHash      string    `json:"tx_hash"`
	TxFee       float64   `json:"tx_fee"`
	TransFee    float64   `json:"trans_fee"`
	Comment     string    `json:"comment"`
	WOrD        int8      `json:"w_or_d"`
	Status      int8      `json:"status"`     // 0:审核中；1:提币中 2: 提币成功 3: 提币种失败
	IsRemoved   int8      `json:"is_removed"` // 0: 正常，1: 删除
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DepositRecordData struct {
	Id          int64     `json:"id"`
	AssetName   string    `json:"asset_name"`
	ChainName   string    `json:"chain_name"`
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	Amount      float64   `json:"amount"`
	TxHash      string    `json:"tx_hash"`
	TxFee       float64   `json:"tx_fee"`
	Status      int8      `json:"status"`     // 0:入账成功；2: 入账失败
	IsRemoved   int8      `json:"is_removed"` // 0: 正常，1: 删除
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
