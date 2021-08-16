package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/question"
	"encoding/json"
	"github.com/astaxie/beego"
)


type QuestionController struct {
	beego.Controller
}

// QsList @Title QsList
// @Description 常见问题列表 QsList
// @Success 200 status bool, data interface{}, msg string
// @router /qs_list [post]
func (this *QuestionController) QsList() {
	qs_lp := question.QsListCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &qs_lp); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	var qs_list_ret []question.QsListRep
	qs_list, total, err := models.GetQuestionsList(qs_lp.Page, qs_lp.PageSize)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "数据库错误")
		this.ServeJSON()
		return
	}
	for _, v := range qs_list {
		qs_ := question.QsListRep{
			QsId:v.Id,
			QsAuthor: v.Author,
			QsTitle: v.QsTitle,
			CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		qs_list_ret = append(qs_list_ret, qs_)
	}
	data := map[string]interface{}{
		"total":     total,
		"qs_list": qs_list_ret,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取常见问题列表成功")
	this.ServeJSON()
	return
}


// QsDetail @Title QsDetail
// @Description 常见问题列表 QsDetail
// @Success 200 status bool, data interface{}, msg string
// @router /qs_detail [post]
func (this *QuestionController) QsDetail() {
	qs_dtl_p := question.QsDetailCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &qs_dtl_p); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	qs, code, err := models.GetQuestionsDetail(qs_dtl_p.QsId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, "数据库错误")
		this.ServeJSON()
		return
	}
	qs_detail := question.QsDetailRep{
		QsId:qs.Id,
		QsAuthor: qs.Author,
		QsTitle: qs.QsTitle,
		QsDetail: qs.QsDetail,
		CreateTime: qs.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, qs_detail, "获取常见问题列表成功")
	this.ServeJSON()
	return
}
