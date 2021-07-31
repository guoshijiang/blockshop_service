package form_validate


type GoodsCateForm struct {
  Id           int64     `form:"id"`
  CatLevel     int8      `form:"cat_level"`                  // 分类级别
  FatherCatId  int64     `form:"father_cat_id"`              // 父级分类 ID
  Name         string    `form:"name"`                        // 分类标题
  Icon         string    `form:"icon"`                        // 分类Icon
  IsShow        int8      `form:"is_show"`                      // 0 不显示 1 显示
  IsCreate 	   int    	 `form:"_create"`
}

func (*GoodsCateForm) Messages() {
  //todo
}