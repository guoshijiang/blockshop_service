package collect


type BlackListRep struct {
	BlId        int64  `json:"bl_id"`
	MerchantId  int64  `json:"merchant_id"`
	MerchanName string `json:"merchan_name"`
	DateTime    string `json:"date_time"`
}

