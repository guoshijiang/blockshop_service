package form_validate

type NewsForm struct {
  Id        int64     `form:"id"`
  Title     string    `form:"title"`
  Abstract  string    `form:"abstract"`
  Content   string    `form:"content"`
  Image     string    `form:"image"`
  Author    string    `form:"author"`
  Views     int64     `form:"views"`
  Likes     int64     `form:"likes"`
  IsCreate 	int    		`form:"_create"`
}