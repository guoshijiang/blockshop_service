package question

import "blockshop/types"

type ContractWayCheck struct {
	QueryWay int8 `json:"query_way"`  // 1:申请入住  2:客户服务
}

type QsListCheck struct {
	types.PageSizeData
}


type QsDetailCheck struct {
	QsId int64 `json:"qs_id"`
}
