package goods

type GoodsListRep struct {
	GoodsId        int64   `json:"goods_id"`
	Title          string  `json:"title"`
	Logo           string  `json:"logo"`
	GoodsPrice     float64 `json:"goods_price"`
	GoodsDisPrice  float64 `json:"goods_discount_price"`
	IsDiscount     int8    `json:"is_discount"`
}
