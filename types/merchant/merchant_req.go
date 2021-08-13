package merchant

import (
	"blockshop/types"
	"github.com/pkg/errors"
)

type MerchantListReq struct {
	types.PageSizeData
	IsShow  int8 `json:"is_show"`
}

func (this MerchantListReq) ParamCheck() (int, error) {
	code, err := this.SizeParamCheck()
	if err != nil {
		return code, err
	}
	return types.ReturnSuccess, nil
}

type MerchantDetailReq struct {
	MerchantId int64 `json:"merchant_id"`
}

func (this MerchantDetailReq) ParamCheck() (int, error) {
	if this.MerchantId <= 0 {
		return types.ParamEmptyError, errors.New("MerchantId 不能小于 0")
	}
	return types.ReturnSuccess, nil
}

type MerchantAddUpdGoodsReq struct {
	UserId         int64    `json:"user_id"`                // 用户ID
	GoodsCatName   string   `json:"goods_cat_name"`         // 商品分类
	GoodsTypeName  string   `json:"goods_type_name"`        // 商品类别
	GoodsAttrKey   string   `json:"goods_attr_key"`
	GoodsAttrValue string   `json:"goods_attr_value"`        // 商品属性
	GoodsImgIds    []int64  `json:"goods_img_ids"`          // 商品图片
	GoodsMark      string   `json:"goods_mark"`      		// 商品备注
	Serveice       string   `json:"serveice"`        		// 服务说明
	CalcWay        int8     `json:"calc_way"`        		// 计量方式
	MerchantId     int64    `json:"merchant_id"`     		// 商品所属商家ID
	Title          string   `json:"title"`           		// 商品标题
	Logo           string   `json:"logo"`            		// 商品封面
	OriginStName   string   `json:"origin_st_name"` 		// 商品的产地
	TotalAmount    int64    `json:"total_amount"`    		// 商品总量
	GoodsPrice     float64  `json:"goods_price"`     		// 商品价格
	GoodsName      string   `json:"goods_name"`             // 产品名称
	GoodsParams    string 	`json:"goods_params"`           // 产品参数
	GoodsDetail    string   `json:"goods_detail"`           // 产品详细介绍
	Discount       float64  `json:"discount"`               // 折扣 取值 0.1-9.9；0代表不打折
	Sale           int8     `json:"sale"`                   // 上架下架: 0:上架 1:下架
	IsDiscount     int8     `json:"is_discount"`      		// 打折活动: 0:不打折，1:打折活动产品
}

func (this MerchantAddUpdGoodsReq) GoodsAddParamCheck() (int, error) {
	if this.UserId <= 0 {
		return types.UserIsNotExist, errors.New("该用户不存在")
	}
	if this.MerchantId <= 0 {
		return types.ParamEmptyError, errors.New("该商户不存在")
	}
	return types.ReturnSuccess, nil
}

type UpdateGoodsReq struct {
	GoodsId   int64  `json:"goods_id"`
	MerchantAddUpdGoodsReq
}

func (this UpdateGoodsReq) GoodsUpdParamCheck() (int, error) {
	if this.GoodsId <= 0 {
		return types.UserIsNotExist, errors.New("该商品不存在")
	}
	if this.UserId <= 0 {
		return types.UserIsNotExist, errors.New("该用户不存在")
	}
	if this.MerchantId <= 0 {
		return types.ParamEmptyError, errors.New("该商户不存在")
	}
	return types.ReturnSuccess, nil
}

type DeleteGoodsReq struct {
	GoodsId   int64  `json:"goods_id"`
}

type StaticDetailReq struct {
  MerchantId  int64       `json:"merchant_id`
}

func (this DeleteGoodsReq) ParamCheck() (int, error) {
	if this.GoodsId <= 0 {
		return types.UserIsNotExist, errors.New("该商品不存在")
	}
	return types.ReturnSuccess, nil
}

func (this StaticDetailReq) ParamCheck() (int, error) {
  if this.MerchantId <= 0 {
    return types.UserIsNotExist, errors.New("参数错误")
  }
  return types.ReturnSuccess, nil
}

type OpenMerchantReq struct {
	UserId      int64  `json:"user_id"`
	PayWay      int8   `json:"pay_way"` //0:BTC，1:USDT
	MctName     string `json:"mct_name"`
	MctAbstruct string `json:"mct_abstruct"`
	MctDetail   string `json:"mct_detail"`
	MctLogo     string `json:"mct_logo"`
	MctService  string `json:"mct_service"`
	MctCrtName  string `json:"mct_crt_name"`
	MctCrtPhone string `json:"mct_crt_phone"`
}
