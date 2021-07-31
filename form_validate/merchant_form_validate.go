package form_validate

import "github.com/gookit/validate"

type MerchantForm struct {
  Id        			int64			`form:"id"`
  MerchantName  		string    		`form:"merchant_name" validate:"required"`
  MerchantIntro  		string    		`form:"merchant_intro" validate:"required"`
  MerchantDetail  	string    		`form:"merchant_detail" validate:"required"`
  ContactUser    		string    		`form:"contact_user"`
  Phone          		string    		`form:"phone"`
  WeChat         		string    		`form:"we_chat"`
  Address  			string    		`form:"address" validate:"required"`
  GoodsNum  			string    		`form:"goods_num"`
  MerchantWay  		int8    		`form:"merchant_way"`
  ShopLevel  			int8    		`form:"shop_level" validate:"required"`
  ShopServer  		int8    		`form:"shop_server" validate:"required"`
  Logo  				string    		`form:"logo"`
  UserName			string			`form:"user_name"`
  IsCreate 			int    			`form:"_create"`
}


//自定义验证返回消息
func (f MerchantForm) Messages() map[string]string {
  return validate.MS{
    "MerchantName.required":        "名称不能为空.",
    "MerchantIntro.required": 		"介绍不能为空.",
    "Address.required": 			"地址不能为空.",
    "ShopLevel.int8":          		"请填写等级",
  }
}