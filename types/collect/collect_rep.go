package collect


type BlackListRep struct {
	BlId        int64  `json:"bl_id"`
	MerchantId  int64  `json:"merchant_id"`
	MerchanName string `json:"merchan_name"`
	DateTime    string `json:"date_time"`
}

type GoodsListRep struct {
	GdsCollectId    int64  `json:"gds_collect_id"`
	GoodsId         int64  `json:"goods_id"`
	GoodsTitle      string `json:"goods_title"`
	GoodsName       string `json:"goods_name"`
	Views           int64  `json:"views"`
	SellNum         int64  `json:"sell_num"`
	LeftNum         int64  `json:"left_num"`
	GoodsPrice      float64 `json:"goods_price"`
	GoodsBtcAmount  float64 `json:"goods_btc_amount"`
	GoodsUsdtAmount float64 `json:"goods_usdt_amount"`
	IsAdmin         int8 `json:"is_admin"`
}

type MerchantListRep struct {
	MctCollectId        int64 `json:"mct_collect_id"`
	MerchantId          int64  `json:"merchant_id"`
	MerchanName         string `json:"merchan_name"`
	DateTime            string `json:"date_time"`
}

