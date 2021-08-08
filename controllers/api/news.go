package api

import (
	"blockshop/models"
	"blockshop/types"
	news2 "blockshop/types/news"
	"encoding/json"
	"github.com/astaxie/beego"
)

type NewsController struct {
	beego.Controller
}

// GetNewsList @Title GetNewsList
// @Description 获取资产列表 PostSendPhoneCode
// @Success 200 status bool, data interface{}, msg string
// @router /get_news_list [post]
func (nc *NewsController) GetNewsList() {
	var pageP types.PageSizeData
	if err := json.Unmarshal(nc.Ctx.Input.RequestBody, &pageP); err != nil {
		nc.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		nc.ServeJSON()
		return
	}
	news_list, total, msg := models.GetNewsList(int64(pageP.Page), int64(pageP.PageSize))
	news_img_url := beego.AppConfig.String("news_img_path")
	var news_lists []news2.News
	for _, value := range news_list {
		news_r := news2.News{
			Id:        value.Id,
			Title:     value.Title,
			Abstract:  value.Abstract,
			Image:     news_img_url + value.Image,
			Author:    value.Author,
			Views:     value.Views,
			Likes:     value.Likes,
			CreatedAt: value.CreatedAt,
		}
		news_lists = append(news_lists, news_r)
	}
	data := map[string]interface{}{
		"total":     total,
		"gds_lst":   news_lists,
	}
	nc.Data["json"] = RetResource(true, types.ReturnSuccess, data, msg)
	nc.ServeJSON()
	return
}

// GetNewsDetail @Title GetNewsDetail
// @Description 获取新闻详情 GetNewsDetail
// @Success 200 status bool, data interface{}, msg string
// @router /news_detail [post]
func (nc *NewsController) GetNewsDetail() {
	var news news2.NewsDatail
	if err := json.Unmarshal(nc.Ctx.Input.RequestBody, &news); err != nil {
		nc.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		nc.ServeJSON()
		return
	}
	var news_s models.News
	news_s.Id = news.NewsId
	news_ret, code, msg := news_s.GetNewsById()
	if code != types.ReturnSuccess {
		nc.Data["json"] = RetResource(false, code, nil, msg)
		nc.ServeJSON()
		return
	}
	nc.Data["json"] = RetResource(true, types.ReturnSuccess, news_ret, msg)
	nc.ServeJSON()
	return
}
