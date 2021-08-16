package help_desk

import "blockshop/types"

type CreateHelpDeskReq struct {
	Author    string `json:"author"`
	Contract  string `json:"contract"`
	HdTitle   string `json:"hd_title"`
	HdDetail  string `json:"hd_detail"`
}


type HdListCheck struct {
	types.PageSizeData
}


type HdDetailCheck struct {
	HdId int64 `json:"hd_id"`
}
