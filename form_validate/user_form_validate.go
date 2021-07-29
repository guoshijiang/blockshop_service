package form_validate


type UserForm struct {
	Id			   int			 `form:"id"`
	Phone          string        `form:"phone"`
	UserName       string        `form:"user_name"`
	Avator         string        `form:"avator"`
	Password       string        `form:"password"`
	Email          string        `form:"email"`
	MemberLevel    int8          `form:"member_level"`
	IsCreate 	   int    		 `form:"_create"`
}