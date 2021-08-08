package forum

import "blockshop/types"

type ForumListReq struct {
	types.PageSizeData
}


type ForumChildListReq struct {
	types.PageSizeData
	CatId int64 `json:"cat_id"`
}

