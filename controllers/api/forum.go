package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/forum"
	"encoding/json"
	"github.com/astaxie/beego"
)

type ForumController struct {
	beego.Controller
}

// ForumList @Title ForumList
// @Description 论坛分类列表 ForumList
// @Success 200 status bool, data interface{}, msg string
// @router /forum_list [post]
func (this *ForumController) ForumList() {
	forum_req := forum.ForumListReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &forum_req); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	var forum_cat models.ForumCat
	forum_list, total, err := forum_cat.GetCatFirstLevel(forum_req)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, err.Error())
		this.ServeJSON()
		return
	}
	var f_level_one_list []forum.ForumLevelOneRep
	for _, value := range forum_list {
		forum_level_list, err := forum_cat.GetCatByFatherId(value.Id)
		if err != nil {
			panic(err)
		}
		var f_level_list []forum.ChildFormRep
		for _, value_fl := range forum_level_list {
			fl := forum.ChildFormRep {
				Id: value_fl.Id,
				Title: value_fl.Name,
				Icon: value_fl.Icon,
			}
			f_level_list = append(f_level_list,  fl)
		}
		fllop := forum.ForumLevelOneRep {
			Id:        value.Id,
			Title:     value.Name,
			ThemeNum: int64(len(forum_level_list)),
			TopicNum:  100,
			Abstruct:  value.Introduce,
			ChildForm: f_level_list,
		}
		f_level_one_list = append(f_level_one_list, fllop)
	}
	data := map[string]interface{}{
		"total": total,
		"form_lst": f_level_one_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取论坛父板块成功")
	this.ServeJSON()
	return
}


