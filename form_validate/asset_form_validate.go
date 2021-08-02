package form_validate

import "github.com/gookit/validate"

type AssetForm struct {
  Id        			int64			    `form:"id"`
  Name  		      string    		`form:"name" validate:"required"`
  Unit  		      int64    		  `form:"unit" validate:"required"`
  IsCreate 			  int    			  `form:"_create"`
}


//自定义验证返回消息
func (f AssetForm) Messages() map[string]string {
  return validate.MS{
    "name.required":        "名称不能为空.",
    "unit.required": 		    "单位不能为空.",
  }
}