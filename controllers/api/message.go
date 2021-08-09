package api

import (
	"encoding/json"
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/message"
	"github.com/astaxie/beego"
)

type MessageController struct {
	beego.Controller
}

// SendMessage @Title SendMessage
// @Description 发送消息 SendMessage
// @Success 200 status bool, data interface{}, msg string
// @router /send_msg [post]
func (this *MessageController) SendMessage() {
	var send_msg message.SendMsgReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &send_msg); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if send_msg.MsgContent != "" {
		msg := models.Message{
			SendUserId: send_msg.SendUserId,
			TargetUserId: send_msg.TargetId,
			MsgType:send_msg.MsgType,
			MsgContent: send_msg.MsgContent,
		}
		err := msg.Insert()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, err, "消息发送失败")
			this.ServeJSON()
			return
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "消息发送成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, types.MessageEmpty, nil, "发送的消息为空")
		this.ServeJSON()
		return
	}
}

// ReciveMessage @Title ReciveMessage
// @Description 接收消息 ReciveMessage
// @Success 200 status bool, data interface{}, msg string
// @router /recive_msg [post]
func (this *MessageController) ReciveMessage() {
	var recive_msg message.ReciveMsgReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &recive_msg); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	msg_list, total, err := models.GetMassageList(recive_msg.UserId, recive_msg.Page, recive_msg.PageSize)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, err.Error())
		this.ServeJSON()
		return
	}
	var massage_list []message.Message
	for _, msg := range msg_list {
		var s_user, r_user  message.MassageUser
		if msg.SendUserId != 0 {
			user_send, _:= models.GetUserById(msg.SendUserId)
			s_user = message.MassageUser{
				UserId: user_send.Id,
				UserName: user_send.UserName,
				UserPhoto: user_send.Avator,
			}
		} else {
			s_user = message.MassageUser{
				UserId: 0,
				UserName: "平台",
				UserPhoto: "",
			}
		}
		if msg.TargetUserId != 0 {
			user_target, _:= models.GetUserById(msg.TargetUserId)
			r_user = message.MassageUser{
				UserId: user_target.Id,
				UserName: user_target.UserName,
				UserPhoto: user_target.Avator,
			}
		} else {
			r_user = message.MassageUser{
				UserId: 0,
				UserName: "平台",
				UserPhoto: "",
			}
		}
		massage := message.Message {
			MsgSendUser: s_user,
			MsgTargetUser: r_user,
			SendUserId: msg.SendUserId,
			TargetUserId: msg.TargetUserId,
			MsgId: msg.Id,
			MsgType:  msg.MsgType,
			MsgSource :1,
			MsgContent: msg.MsgContent,
			SendReciveTime: msg.CreatedAt,
		}
		massage_list = append(massage_list, massage)
	}
	data_ret := map[string]interface{}{
		"total": total,
		"msg_list": massage_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data_ret, "获取消息成功")
	this.ServeJSON()
	return
}

