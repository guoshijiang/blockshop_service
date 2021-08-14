package merchant

type MerchantListRep struct {
	MctId          int64  `json:"mct_id"`
	MctName        string `json:"mct_name"`
	MctIntroduce   string `json:"mct_introduce"`
	MerchantDetail string `json:"merchant_detail"`
	MctLogo        string `json:"mct_logo"`
	MctWay         int8   `json:"mct_way"`      // 0:自营商家； 1:认证商家  2:普通商家
	ShopLevel      int8   `json:"shop_level"`   // 店铺等级
	ShopServer     int8   `json:"shop_server"`  // 店铺服务
}


type OrderDataStat struct {
	WaidPayOrderNum    int64 `json:"waid_pay_order_num"`
	WaitSendOrderNum   int64 `json:"wait_send_order_num"`
	WaitReturnOrderNum int64 `json:"wait_return_order_num"`
	SendOrderNum       int64 `json:"send_order_num"`
}

type GoodsDataStat struct {
	OnSaleNum    int64 `json:"on_sale_num"`
	SoldOutNum   int64 `json:"sold_out_num"`
	OffShelfNum  int64 `json:"off_shelf_num"`
}

type CommentDataStat struct {
	SericeBest     int64   `json:"serice_best"`
	ServiceGood    int64   `json:"service_good"`
	ServiceBad     int64   `json:"service_bad"`
	ServiceAvg	   float64 `json:"service_avg"`
	TradeBest 	   int64   `json:"trade_best"`
	TradeGood 	   int64   `json:"trade_good"`
	TradeBad 	   int64   `json:"trade_bad"`
	TradeAvg       float64 `json:"trade_avg"`
	QualityBest    int64   `json:"quality_best"`
	QualityGood    int64   `json:"quality_good"`
	QualityBad     int64   `json:"quality_bad"`
	QualityAvg     float64 `json:"quality_avg"`
}

type MerchantDetailRep struct {
	MctId             int64            `json:"id"`
	MctLogo           string           `json:"logo"`
	MctName           string           `json:"merchant_name"`
	MctIntroduce      string           `json:"merchant_intro"`
	MerchantDetail    string           `json:"merchant_detail"`
	Address           string           `json:"address"`
	GoodsNum          int64            `json:"goods_num"`
	MctWay            int8             `json:"merchant_way"`
	ShopLevel      	  int8             `json:"shop_level"`
	ShopServer        int8             `json:"shop_server"`
	OrderStat         *OrderDataStat   `json:"order_stat"`
	GoodsStat         *GoodsDataStat   `json:"goods_stat"`
	CommentStat       *CommentDataStat `json:"comment_stat"`
	ShopScore         int64            `json:"shop_score"`
	MonthSellNum      int64            `json:"month_sell_num"`
	MonthSellAmount   float64          `json:"month_sell_amount"`
	TotalSellNum      int64            `json:"total_sell_num"`
	TotalSellAmount   float64          `json:"total_sell_amount"`
	AdjustVictor      int64            `json:"adjust_victor"`
	AdjustFail        int64            `json:"adjust_fail"`
	JoinTime          string           `json:"join_time"`
	LstLoginTime	  string           `json:"lst_login_time"`
}
