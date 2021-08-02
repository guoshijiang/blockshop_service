package form_validate

type MessageForm struct {
  Id              int64      `form:"id"`
  SendUserId      int64      `form:"send_user_id"`     //0:为平台发送;其他数字对应用户ID
  TargetUserId    int64  	   `form:"target_user_id"`  // 0:为发送给平台;其他数字对应用户ID
  MsgType         int8       `form:"msg_type"`        // 0:文字消息；1:图片消息
  MsgText         string     `form:"msg_text"`                // 消息内容
  MsgImg          string     `form:"msg_img"`                 // 消息内容
  MsgContent      string     `form:"content"`
  IsCreate 	 	    int        `form:"_create"`
}