package user

type LoginRep struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`   // 用户名
	Token    string `json:"token"`       // token
}

type TwoFactorRep struct {
	Id          int64  `json:"id"`
	UserName    string `json:"user_name"`    // 用户名
	CipherText  string `json:"cipher_text"`  // cipher_text
}

type UserInfoRep struct {
	UserId          int64   `json:"user_id"`
	Photo           string  `json:"photo"`
	UserName        string  `json:"user_name"`
	IsMerchant      int8    `json:"is_merchant"`
	JoinTime        string  `json:"join_time"`
	TrustLevel      int8    `json:"trust_level"`
	BtcOrderAmount  string  `json:"btc_order_amount"`
	UsdtOrderAmount string  `json:"usdt_order_amount"`
	AdjustVictor    int64   `json:"adjust_victor"`
	AdjustFail      int64   `json:"adjust_fail"`
	BtcBalance      string  `json:"btc_balance"`
	UsdtBalance     string  `json:"usdt_balance"`
	BtcSpend        string  `json:"btc_spend"`
	UsdtSpend       string  `json:"usdt_spend"`
	BtdAddress      string  `json:"btd_address"`
	UsdtAddress     string  `json:"usdt_address"`
	PublicKey       string  `json:"public_key"`
}