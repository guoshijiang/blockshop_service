package forum

type ChildFormRep struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
}


type ForumLevelOneRep struct {
	Id          int64          `json:"id"`
	Title       string         `json:"title"`
	Abstruct    string         `json:"abstruct"`
	ThemeNum    int64          `json:"theme_num"`   // 主题
	TopicNum    int64          `json:"topic_num"`   // 帖子
	ChildForm   []ChildFormRep `json:"child_form"`
}
