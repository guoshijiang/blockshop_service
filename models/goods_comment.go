package models

import (
	"blockshop/common"
	"blockshop/types"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
  "time"
)

type GoodsComment struct {
	BaseModel
	Id           int64     `orm:"pk;column(id);auto;size(11)" description:"评论ID" json:"id"`
	GoodsId      int64     `orm:"column(goods_id);size(64)" description:"评论商品" json:"goods_id"`
	UserId       int64     `orm:"column(user_id);default(1);" description:"评论人" json:"user_id"`
  MerchantId   int64     `orm:"column(merchant_id);default(1);" description:"商家" json:"merchant_id"`
	QualityStar  int8      `orm:"column(quality_star);default(5);index" description:"质量评论星级" json:"quality_star"`
  ServiceStar  int8      `orm:"column(service_star);default(5);index" description:"服务评论星级" json:"service_star"`
	TradeStar    int8       `orm:"column(trade_star);default(5);index" description:"交易评论星级" json:"trade_star"`
	Content      string    `orm:"column(content);type(text)" description:"评论内容"  json:"content"`
	ImgOneId     int64     `orm:"column(img_one_id);size(64)" description:"评论图片1" json:"img_one_id"`
	ImgTwoId     int64     `orm:"column(img_two_id);size(64)" description:"评论图片2" json:"img_two_id"`
	ImgThreeId   int64     `orm:"column(img_three_id);size(64)" description:"评论图片3" json:"img_three_id"`
}

func (this *GoodsComment) TableName() string {
	return common.TableName("goods_comment")
}


func (this *GoodsComment) SearchField() []string {
	return []string{"title"}
}

func (this *GoodsComment) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsComment) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsComment) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsComment) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}

func GetGoodsCommentList(page, pageSize int, goods_id int64) ([]*GoodsComment, int64, error) {
	offset := (page - 1) * pageSize
	gct_list := make([]*GoodsComment, 0)

	cond := orm.NewCondition()
	cond.And("")

	query := orm.NewOrm().QueryTable(GoodsComment{}).Filter("GoodsId", goods_id)
	total, _ := query.Count()

	_, err := query.Limit(pageSize, offset).All(&gct_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return gct_list, total, nil
}

//评价类型统计
// ty 1表示差评 2表示中评 3表示好评
func GetGoodsCommentStar(end time.Time) []map[string]int64{
  cond := orm.NewCondition()
  data := make([]map[string]int64,0)
  condExprBad := cond.And("quality_star__lt",3).Or("service_star__lt",3).Or("trade_star__lt",3).And("created_at__lt",end)
  bad, err := orm.NewOrm().QueryTable(GoodsComment{}).SetCond(condExprBad).Count()
  if err != nil {
    bad = 0
  }
  data = append(data, map[string]int64{"bad":bad})
  CondExprMid := cond.And("quality_star",3).Or("service_star",3).Or("trade_star",3).And("created_at__lt",end)
  mid, err := orm.NewOrm().QueryTable(GoodsComment{}).SetCond(CondExprMid).Count()
  if err != nil {
    mid = 0
  }
  data = append(data, map[string]int64{"bad":mid})

  CondExprGood := cond.And("quality_star__gt",3).Or("service_star__gt",3).Or("trade_star__gt",3).And("created_at__lt",end)
  good, err := orm.NewOrm().QueryTable(GoodsComment{}).SetCond(CondExprGood).Count()
  if err != nil {
    good = 0
  }
  data = append(data, map[string]int64{"bad":good})
  return data
}