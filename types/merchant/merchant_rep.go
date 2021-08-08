package merchant

import "time"

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

type MerchantDetailRep struct {
	MctId             int64     `json:"id"`
	MctLogo           string    `json:"logo"`
	MctName           string    `json:"merchant_name"`
	MctIntroduce      string    `json:"merchant_intro"`
	MerchantDetail    string    `json:"merchant_detail"`
	Address           string    `json:"address"`
	GoodsNum          int64     `json:"goods_num"`
	MctWay            int8      `json:"merchant_way"`
	ShopLevel      	  int8      `json:"shop_level"`
	ShopServer        int8      `json:"shop_server"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt	      time.Time `json:"updated_at"`
}
