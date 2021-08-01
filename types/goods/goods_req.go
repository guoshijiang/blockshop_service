package goods

import (
	"blockshop/types"
)

type GoodsListReq struct {
	types.PageSizeData
	GoodsName     string   `json:"goods_name"`     // 商品名称
	TypeId     	  int64    `json:"type_id"`        // 类别ID
	CatId         int64    `json:"cat_id"`         // 类别ID
	OriginCountry string   `json:"origin_country"` // 产地
	StartPrice    float64  `json:"start_price"`    // 起始价格
	EndPrice      float64  `json:"end_price"`      // 结束价格
	OrderBy       int8     `json:"order_by"`       // 0:时间，1:销量；2:价格; 3:商家
	PayWay        string   `json:"pay_way"`        // 支付方式；BTC/USDT
}

func (this GoodsListReq) ParamCheck() (int, error) {
	code, err := this.ParamCheck()
	if err != nil {
		return code, err
	}
	return types.ReturnSuccess, nil
}

type GoodsDetailReq struct {
	UserId   int64   `json:"user_id"`
	GoodsId  int64   `json:"goods_id"`
}

func (this GoodsDetailReq) ParamCheck() (int, error) {
	return types.ReturnSuccess, nil
}

