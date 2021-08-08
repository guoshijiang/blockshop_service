package forum

import "blockshop/types"

type ForumListReq struct {
	types.PageSizeData
}

type ForumChildListReq struct {
	types.PageSizeData
	CatId int64 `json:"cat_id"`
}

type ForumTopicListReq struct {
	types.PageSizeData
	LevelCatId int64 `json:"level_cat_id"`
}

type ForumTopicDetailReq struct {
	types.PageSizeData
	ForumId int64 `json:"forum_id"`
}

type ForumTopicCatReq struct {
	TcName   string  `json:"tc_name"`
	IsTc     int     `json:"is_tc"` //1:分类，2:话题
}

type CreateForumReq struct {
	Title    string  `json:"title"`
	Abstract string  `json:"abstract"`
	CatName  string  `json:"cat_name"`
	TopName  string  `json:"top_name"`
	Content  string  `json:"content"`
}

type CreateCmtReplyReq struct {
	ForumId        int64  `json:"forum_id"`
	FatherCmtId    int64  `json:"father_cmt_id"`
	CtmReply       string `json:"ctm_reply"`
}


