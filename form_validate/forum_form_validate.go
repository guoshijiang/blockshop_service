package form_validate

type ForumForm struct {
  Id             int64         `form:"id" validate:"required"`
  IsCheck        int8          `form:"is_check"`  // 0:未审核 1:已审核
  IsCreate 	     int    	     `form:"_create"`
}