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
	CatId    int64   `json:"cat_id"`     // 分类ID
	IsFather int8    `json:"is_father"`  // 是不是父级，0:不是，1:是
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

type ForumTopiceLikeReq struct {
	ForumId    int64 `json:"forum_id"`
}

type CommentReplyLikeReq struct {
	CmtReplyId    int64 `json:"cmt_reply_id"`
}

