package message


import "blockshop/types"

type SendMsgReq struct {
	SendUserId   int64  `json:"send_user_id"`
	TargetId     int64  `json:"target_id"` // 默认传 0
	MsgType      int8   `json:"msg_type"`  // 0:文字消息，1:图片消息
	MsgContent   string `json:"msg_content"`
}


type ReciveMsgReq struct {
	types.PageSizeData
	UserId        int64  `json:"user_id"`
}
