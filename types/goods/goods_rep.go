package goods

type GoodsListRep struct {
	GoodsId        int64   `json:"goods_id"`
	Title          string  `json:"title"`
	Logo           string  `json:"logo"`
	GoodsPrice     float64 `json:"goods_price"`
	GoodsDisPrice  float64 `json:"goods_discount_price"`
	IsDiscount     int8    `json:"is_discount"`
}

type GoodsImagesRet struct {
	GoodsImgId  int64  `json:"goods_img_id"`
	ImageUrl    string `json:"image_url"`
}

type GoodsAttrRet struct {
	GdsAttrKey   string   `json:"gds_attr_key"`
	GdsAttrValue []string `json:"gds_attr_value"`
}
