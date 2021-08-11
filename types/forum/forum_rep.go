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

type ForumTopicListRep struct {
	UserId      int64   `json:"user_id"`
	FormId      int64   `json:"form_id"`
	UserName    string  `json:"user_name"`
	UserPhoto   string  `json:"user_photo"`
	Title       string  `json:"title"`
	DataTime    string  `json:"data_time"`
	Views       int64   `json:"views"`
	Likes       int64   `json:"likes"`
	Answers     int64   `json:"answers"`
}

type ForumReply struct {
	Id          int64   `json:"id"`
	UserName    string  `json:"user_name"`
	UserPhoto   string  `json:"user_photo"`
	Reply       string  `json:"reply"`
	Datetime    string  `json:"datetime"`
}

type ForumCommentListRep struct {
	Id          int64        `json:"id"`
	UserName    string       `json:"user_name"`
	UserPhoto   string       `json:"user_photo"`
	Comment     string       `json:"comment"`
	Datetime    string       `json:"datetime"`
	Reply       []ForumReply `json:"reply"`
}


type FmTopicCatListRep struct {
	TcId  int64 `json:"tc_id"`
	TcName string `json:"tc_name"`
}