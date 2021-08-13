package collect


type BlackListReq struct {
	UserId      int64  `json:"user_id"`
	MerchantId  int64  `json:"merchant_id"`
}

type GoodsCollectReq struct {
	UserId      int64  `json:"user_id"`
	GoodsId     int64  `json:"goods_id"`
}

type MerchantCollectReq struct {
	UserId      int64  `json:"user_id"`
	MctId     	int64 `json:"mct_id"`
}
