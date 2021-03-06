package order

import "time"

type OrderListRet struct {
	MerchantId   int64    `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	MerchantPhone string `json:"merchant_phone"`
	OrderId      int64      `json:"order_id"`
	GoodsName    string    `json:"goods_name"`
	GoodsTitel   string    `json:"goods_titel"`
	GoodsLogo    string    `json:"goods_logo"`
	GoodsPrice   float64  `json:"goods_price"`
	RecUser      string   `json:"rec_user"`
	RecPhone     string `json:"rec_phone"`
	RecAddress   string `json:"rec_address"`
	OrderStatus  int8    `json:"order_status"`
	BuyNums      int64   `json:"buy_nums"`
	PayCnyPrice   float64 `json:"pay_cny_price"`
	PayCoinAmount float64 `json:"pay_coin_amount"`
	PayWay      int8      `json:"pay_way"`
	IsCancle    int8      `json:"is_cancle"`
	IsComment   int8     `json:"is_comment"`
	IsDiscount  int8    `json:"is_discount"`   // 0:不打折，1:打折活动产品
	IsAdmin     int8    `json:"is_admin"`
	CreateTime  string `json:"create_time"`
}

type ReturnOrderProcess struct {
	ProcessId      int64  `json:"process_id"`
	ReturnUser     string `json:"return_user"`
	ReturnPhone    string `json:"return_phone"`
	ReturnAddress  string `json:"return_address"`
	ReturnReson    string `json:"return_reson"`   // 退货原因
	ReturnAmount   float64 `json:"return_amount"`
	RejectDescribe string `json:"reject_describe"`  // 拒绝描述
	QuestionDes    string  `json:"question_des"`   // 问题描述
	AcceptRejectRs string `json:"accept_reject_rs"`
	AdjustContent  string  `json:"adjust_content"`   // 申述描述
	QsImgOne       string  `json:"qs_img_one"`
	QsImgTwo       string  `json:"qs_img_two"`
	QsImgThree     string  `json:"qs_img_three"`
	AskTime        time.Time `json:"ask_time"`
	Process        int8  `json:"process"`   	// 0:等待卖家确认; 1:卖家已同意; 2:卖家拒绝; 3:等待买家邮寄; 4:等待卖家收货; 5:卖家已经发货; 6:等待买家收货; 7:已完成
	IsRecvGoods    int8  `json:"is_recv_goods"`
	LeftTime       int64 `json:"left_time"`
}

type OrderDetailRet struct {
	OrderId    int64    `json:"order_id"`
	GoodsId   int64     `json:"goods_id"`
	RecUser    string   `json:"rec_user"`
	RecPhone   string   `json:"rec_phone"`
	RecAddress string   `json:"rec_address"`
	MerchantId int64    `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	GoodsName string    `json:"goods_name"`
	GoodsTilte string   `json:"goods_tilte"`
	GoodsPrice float64  `json:"goods_price"`
	PayIntegral float64 `json:"pay_integral"`
	SendIntegral float64 `json:"send_integral"`
	GoodsLogo string `json:"goods_logo"`
	OrderStatus int8    `json:"order_status"`
	BuyNums     int64   `json:"buy_nums"`
	PayCnyPrice float64 `json:"pay_cny_price"`
	PayCoinAmount float64 `json:"pay_coin_amount"`
	ShipFee     float64 `json:"ship_fee"`
	Logistics	string  `json:"logistics"`
	ShipNumber  string  `json:"ship_number"`
	RetShipNumber  string `json:"ret_ship_number"`
	Coupons     float64 `json:"coupons"`
	PayWay      int8    `json:"pay_way"`
	OrderNumber string  `json:"order_number"`
	PayTime     *time.Time `json:"pay_time"`
	CreateTime  time.Time `json:"create_time"`
	IsCancle    int8      `json:"is_cancle"`
	IsComment   int8     `json:"is_comment"`
	IsDiscount  int8    `json:"is_discount"`   // 0:不打折，1:打折活动产品
	IsIntegral  int8    `json:"is_integral"`
	GoodsTypes  string `json:"goods_types"`
	IsAdmin     int8 `json:"is_admin"`
	RetrurnOrder  *ReturnOrderProcess `json:"retrurn_order"`
}

type StateStatic struct {
  Total       int64            `json:"total"`
  State       int8             `json:"state"`
}

type StateStaticReq struct {
  Total       int64           `json:"total"`
  StateName   string          `json:"state_name"`
  State       int8            `json:"state"`
}
