package services

import (
	"blockshop/common/utils"
	beego_pagination "blockshop/common/utils/beego-pagination"
	"blockshop/models"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strconv"
	"strings"
	"time"
)

 var AdminUserVal	*models.AdminUser

type BaseService struct {
	//可搜索字段
	SearchField []string
	//可作为条件的字段
	WhereField []string
	//可做为时间范围查询的字段
	TimeField []string
	//禁止删除的数据id,在model中声明就可以了，可不用在此处声明
	//NoDeletionId []int
	//分页
	Pagination beego_pagination.Pagination
}

func (this *BaseService) SetAdmin(admin *models.AdminUser){
	AdminUserVal = admin
}

//分页处理
func (this *BaseService) Paginate(seter orm.QuerySeter, listRows int, parameters url.Values) orm.QuerySeter {
	var pagination beego_pagination.Pagination
	qs := pagination.Paginate(seter, listRows, parameters)
	this.Pagination = pagination
	return qs
}

func (this *BaseService) PaginateMul(listRows int, parameters url.Values) {
	var pagination = &beego_pagination.Pagination{Total: this.Pagination.Total}
	this.Pagination = *pagination.PaginateMul(listRows, parameters)
}


//查询处理
func (this *BaseService) ScopeWhere(seter orm.QuerySeter, parameters url.Values) orm.QuerySeter {
	//关键词like搜索
	keywords := parameters.Get("_keywords")
	cond := orm.NewCondition()
	if keywords != "" && len(this.SearchField) > 0 {
		for _, v := range this.SearchField {
			cond = cond.Or(v+"__icontains", keywords)
		}
	}

	merchantId,_ := strconv.Atoi(parameters.Get("_merchant_id"))
	if merchantId > 0 {
		cond = cond.And("merchant_id",merchantId)
	}
	UserId,_ := strconv.Atoi(parameters.Get("_user_id"))
	if UserId > 0 {
		cond = cond.And("user_id",UserId)
	}
	//字段条件查询
	if len(this.WhereField) > 0 && len(parameters) > 0 {
		for k, v := range parameters {
			if v[0] != "" && utils.InArrayForString(this.WhereField, k) {
				cond = cond.And(k, v[0])
			}
		}
	}

	//时间范围查询
	if len(this.TimeField) > 0 && len(parameters) > 0 {
		for key, value := range parameters {
			if value[0] != "" && utils.InArrayForString(this.TimeField, key) {
				timeRange := strings.Split(value[0], " - ")
				startTimeStr := timeRange[0]
				endTimeStr := timeRange[1]

				loc, _ := time.LoadLocation("Local")
				startTime, err := time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, loc)

				if err == nil {
					unixStartTime := startTime.Unix()
					if len(endTimeStr) == 10 {
						endTimeStr += "23:59:59"
					}

					endTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, loc)
					if err == nil {
						unixEndTime := endTime.Unix()
						cond = cond.And(key+"__gte", unixStartTime).And(key+"__lte", unixEndTime)
					}
				}
			}
		}
	}

	//将条件语句拼装到主语句中
	seter = seter.SetCond(cond)

	//排序
	order := parameters.Get("_order")
	by := parameters.Get("_by")
	if order == "" {
		order = "id"
	}

	if by == "" {
		by = "-"
	} else {
		if by == "asc" {
			by = ""
		} else {
			by = "-"
		}
	}

	//排序
	seter = seter.OrderBy(by + order)

	return seter
}


//查询处理
func (this *BaseService) ScopeWhereRaw(parameters url.Values) (string,[]interface{}) {
	//关键词like搜索
	keywords := parameters.Get("_keywords")
	conditionRaw := ""
	condition := make([]interface{},0)
	if keywords != "" && len(this.SearchField) > 0 {
		for _,v := range this.SearchField {
			conditionRaw += " and "+v+" like ? "
			condition = append(condition,keywords)
		}
	}
	merchantId,_ := strconv.Atoi(parameters.Get("_merchant_id"))
	if merchantId > 0 {
		conditionRaw += " and merchant_id = ? "
		condition = append(condition,merchantId)
	}
	//字段条件查询
	if len(this.WhereField) > 0 && len(parameters) > 0 {
		for k, v := range parameters {
			if v[0] != "" && utils.InArrayForString(this.WhereField, k) {
				if strings.Contains(k,";") {
					kstr := strings.Split(k,";")
					if kstr[1] == "gte" {
						conditionRaw += " and "+kstr[0]+" >= ? "
					} else if kstr[1] == "lte" {
						conditionRaw += " and "+kstr[0]+" <= ? "
					} else {
						conditionRaw += " and "+kstr[0]+" = ? "
					}
				} else {
					conditionRaw += " and "+k+" like ? "
				}
				condition = append(condition,v[0])
			}
		}
	}
	//时间范围查询
	if len(this.TimeField) > 0 && len(parameters) > 0 {
		for key, value := range parameters {
			if value[0] != "" && utils.InArrayForString(this.TimeField, key) {
				timeRange := strings.Split(value[0], " - ")
				startTimeStr := timeRange[0]
				endTimeStr := timeRange[1]

				loc, _ := time.LoadLocation("Local")
				startTime, err := time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, loc)

				if err == nil {
					unixStartTime := startTime.Unix()
					if len(endTimeStr) == 10 {
						endTimeStr += "23:59:59"
					}

					endTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, loc)
					if err == nil {
						unixEndTime := endTime.Unix()
						conditionRaw += " and "+key+" >= ? and "+ key+ " <= ? "
						condition = append(condition,unixStartTime)
						condition = append(condition,unixEndTime)
					}
				}
			}
		}
	}
	return conditionRaw,condition
}


//分页和查询合并，多用于首页列表展示、搜索
func (this *BaseService) PaginateAndScopeWhere(seter orm.QuerySeter, listRows int, parameters url.Values) orm.QuerySeter {
	return this.Paginate(this.ScopeWhere(seter, parameters), listRows, parameters)
}

func (this *BaseService) PaginateRaw(listRows int, parameters url.Values) {
	this.PaginateMul(listRows, parameters)
}
