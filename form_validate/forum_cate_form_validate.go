package form_validate

type ForumCateForm struct {
  Id           int64     `form:"id"`
  FatherCatId  int64     `form:"father_cat_id"`
  Name         string    `form:"name"`
  Introduce    string    `form:"introduce"`
  Icon         string    `form:"icon"`
  IsShow       int8      `form:"is_show"`   // 0 显示 1 不显示
  IsCreate 	   int    	 `form:"_create"`

}