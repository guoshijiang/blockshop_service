package form_validate

import "github.com/gookit/validate"

type GoodsForm struct {
  Id             int64     `form:"id"`
  GoodsCatId     int64     `form:"goods_cat_id" validate:"required"`					// 商品所属一级分类ID
  GoodsMark      string    `form:"goods_mark" validate:"required"`    				// 商品备注
  Serveice       string    `form:"serveice" validate:"required"`      				// 服务说明
  CalcWay        int8      `form:"calc_way" validate:"int"`     					    // 0:按件计量 1:按近计量
  OriginStateId  int64     `form:"origin_state_id"`
  MerchantId     int64     `form:"merchant_id" validate:"required"`                   // 商品所属商家ID
  Title          string    `form:"title" validate:"required"`         				// 商品标题
  Logo           string    `form:"logo"`   						                    // 商品封面
  TotalAmount    int64     `form:"total_amount" validate:"required"`                 	// 商品总量
  GoodsPrice     float64   `form:"goods_price" validate:"required"`                   // 商品价格
  GoodsDisPrice  float64   `form:"goods_discount_price" validate:"required"`          // 商品折扣价格
  GoodsName      string    `form:"goods_name" validate:"required"`  					// 产品名称
  GoodsParams    string    `form:"goods_params" validate:"required"`   				// 产品参数
  GoodsDetail    string    `form:"goods_detail" validate:"required"`   				// 产品详细介绍
  Discount       float64   `form:"discount" validate:"float64"`       				// 折扣，取值 0.1-9.9；0代表不打折
  Sale           int8      `form:"sale" validate:"int"`             					// 0:上架 1:下架
  IsDisplay      int8      `form:"is_display" validate:"int"` 						// 0:首页不展示, 1:首页展示
  IsHot          int8      `form:"is_hot" validate:"int"`                       		// 0:非爆款产品 1:爆款产品
  IsDiscount     int8      `form:"is_discount" validate:"int"`                  		// 0:不打折，1:打折活动产品
  IsIntegral     int8      `form:"is_integral" validate:"int" `                  		// 0:非积分兑换产品 1:积分兑换产品
  IsLimitTime    int8      `form:"is_limit_time" validate:"int"`                	// 0:不是限时产品 1:是限时
  IsCreate 	   int    	 `form:"_create"`
}


func (g GoodsForm) Messages() map[string]string {
  return validate.MS{
    "GoodsCatId.required":        	"商品所属一级分类不能为空.",
    "GoodsMark.required": 			    "商品备注不能为空.",
    "Serveice.required": 			"服务说明不能为空.",
    "CalcWay.int":          		"单位不能为空",
    "Title.required":          		"商品标题不能为空",
    "TotalAmount.required":         "商品总量不能为空",
    "GoodsPrice.required":          "商品价格不能为空",
    "GoodsDisPrice.required":       "商品折扣不能为空",
    "GoodsParams.required":         "产品参数不能为空",
    "GoodsName.required":          	"产品名称不能为空",
    "Discount.float64":             "折扣不能为空",
    "Sale.int":                     "上下架必须选择",
    "IsDisplay.int":                "首页是否展示必须选择",
    "IsHot.int":                    "是否热销必须选择",
    "IsDiscount.int":               "是否折扣产品必须选择",
    "IsIntegral.int":               "是否积分兑换必须选择",
    "IsLimitTime.int":              "是否限时产品必须选择",
  }
}