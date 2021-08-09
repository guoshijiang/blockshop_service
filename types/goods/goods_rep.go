package goods

type GoodsType struct {
	Id int64 `json:"id"`
	TypeName string `json:"type_name"`
}

type GoodsCat struct {
	Id int64 `json:"id"`
	CatName string `json:"cat_name"`
}

type OriginState struct {
	Id int64 `json:"id"`
	StateName string `json:"state_name"`
}

type OrderBy struct {
	Way int64 `json:"way"`
	WayName string `json:"way_name"`
}

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
