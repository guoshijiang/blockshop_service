package forum

type ChildFormRep struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
}

type LastestForumRep struct {
	UserId      int64   `json:"user_id"`
	FormId      int64   `json:"form_id"`
	UserName    string  `json:"user_name"`
	UserPhoto   string  `json:"user_photo"`
	LstComment  string  `json:"lst_comment"`
	DataTime    string  `json:"data_time"`
}

type ForumLevelOneRep struct {
	Id          int64           `json:"id"`
	Title       string          `json:"title"`
	Abstruct    string          `json:"abstruct"`
	ThemeNum    int64           `json:"theme_num"`   // 主题
	TopicNum    int64           `json:"topic_num"`   // 帖子
	ChildForm   []ChildFormRep  `json:"child_form"`
	LastestForm *LastestForumRep `json:"lastest_form"`
}

type ForumLevelChildRep struct {
	Id          int64           `json:"id"`
	Title       string          `json:"title"`
	TopicNum    int64           `json:"topic_num"`   // 帖子
	ReplyNum    int64           `json:"reply_num"`
	LastestForm *LastestForumRep `json:"lastest_form"`
}
