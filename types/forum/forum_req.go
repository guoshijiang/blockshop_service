package forum

import "blockshop/types"

type ForumListReq struct {
	types.PageSizeData
	FormLevel   int8 `json:"form_level"`
}
