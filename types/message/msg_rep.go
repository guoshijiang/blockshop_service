package message

import "time"

type MassageUser struct {
	UserId  int64 `json:"user_id"`
	UserName string `json:"user_name"`
	UserPhoto string `json:"user_photo"`
}

type Message struct {
	MsgSendUser      MassageUser `json:"msg_send_user"`
	MsgTargetUser    MassageUser `json:"msg_target_user"`
	SendUserId   int64  `json:"send_user_id"`
	TargetUserId int64  `json:"target_user_id"`
	MsgId        int64  `json:"msg_id"`
	MsgType      int8   `json:"msg_type"`
	MsgSource    int8   `json:"msg_source"`
	MsgContent   string `json:"msg_content"`
	SendReciveTime time.Time `json:"send_recive_time"`
}
