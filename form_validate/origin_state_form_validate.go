package form_validate

type OriginStateForm struct {
  Id           int64     `form:"id"`
  Name         string    `form:"name"`
  IsShow       int8      `form:"is_show"`   // 0 显示 1 不显示
  IsCreate 	   int     `form:"_create"`
}