package collect


type BlackListReq struct {
	UserId      int64  `json:"user_id"`
	MerchantId  int64  `json:"merchant_id"`
}

type BlackListDelReq struct {
	BlackListId  int64  `json:"black_list_id"`
}

type GoodsCollectReq struct {
	UserId      int64  `json:"user_id"`
	GoodsId     int64  `json:"goods_id"`
}

type GoodsCollectDelReq struct {
	GoodsCollecId  int64 `json:"goods_collec_id"`
}

type MerchantCollectReq struct {
	UserId      int64  `json:"user_id"`
	MerchantId  int64  `json:"merchant_id"`
}

type MerchantCollectDelReq struct {
	MctCollectId  int64 `json:"mct_collect_id"`
}
