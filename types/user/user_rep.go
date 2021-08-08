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

type UserWalletStat struct {
	OutAmount float64  `json:"spent_amount"`  // 花费金额
	InAmount  float64  `json:"buy_amount"`    // 收入金额
	Balance   float64  `json:"balance"`       // 托管金额，也就是余额
	Address   string   `json:"address"`
}

type CoinPrice struct {
	Asset      string `json:"asset"`
	ChainName  string `json:"chain_name"`
	UsdPrice   string `json:"usd_price"`
	CnyPrice   string `json:"cny_price"`
}

type UserSecrity struct {
	AccountPct  string  `json:"account_pct"`
	IsSetKey    bool    `json:"is_set_key"`
	IsOpen2Fa   bool    `json:"is_open_2_fa"`
}

type UserInfoRep struct {
	UserId          int64   `json:"user_id"`
	Photo           string  `json:"photo"`
	UserName        string  `json:"user_name"`
	IsMerchant      int8    `json:"is_merchant"`
	JoinTime        string  `json:"join_time"`
	TrustLevel      int8    `json:"trust_level"`
	PublicKey       string  `json:"public_key"`
	BtcOrderAmount  string  `json:"btc_order_amount"`
	UsdtOrderAmount string  `json:"usdt_order_amount"`
	AdjustVictor    int64   `json:"adjust_victor"`
	AdjustFail      int64   `json:"adjust_fail"`
	UserSecrity     UserSecrity    `json:"user_secrity"`  // 账号的安全形
	BtcWtStat       UserWalletStat `json:"btc_wt_stat"`   // BTC 钱包情况
	UsdtWtStat      UserWalletStat `json:"usdt_wt_stat"`  // USDT 钱包的情况
	BtcPrice        CoinPrice      `json:"btc_price"`     // BTC 价格
	UsdtPrice       CoinPrice      `json:"usdt_price"`    // USDT 价格情况
}