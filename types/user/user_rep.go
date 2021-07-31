package user

type LoginRep struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`   // 用户名
	Token    string `json:"token"`       // token
}
