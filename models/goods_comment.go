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
	TradeStar    int8      `orm:"column(trade_star);default(5);index" description:"交易评论星级" json:"trade_star"`
	Content      string    `orm:"column(content);type(text)" description:"评论内容"  json:"content"`
	ImgOneUrl    string    `orm:"column(img_one_url);size(64)" description:"评论图片1" json:"img_one_url"`
	ImgTwoUrl    string    `orm:"column(img_two_url);size(64)" description:"评论图片2" json:"img_two_url"`
	ImgThreeUrl  string    `orm:"column(img_three_url);size(64)" description:"评论图片3" json:"img_three_url"`
}

func (this *GoodsComment) TableName() string {
	return common.TableName("goods_comment")
}

func (this *GoodsComment) SearchField() []string {
	return []string{"content"}
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

func GetGoodsCommentList(page, page_size int, goods_id, mct_id int64, cmt_status int8) ([]*GoodsComment, int64, error) {
	offset := (page - 1) * page_size
	gct_list := make([]*GoodsComment, 0)
	query := orm.NewOrm().QueryTable(GoodsComment{}).Filter("is_removed", 0).OrderBy("-id")
	if goods_id >=1 {
		query =  query.Filter("goods_id", goods_id)
	}
	if mct_id >=1 {
		query =  query.Filter("merchant_id", mct_id)
	}
	// 1:好评, 四星级以上
	if cmt_status ==1 {
		query =  query.Filter("quality_star__gte", 4)
	}
	// 2:中评 二到四星级
	if cmt_status ==2 {
		query =  query.Filter("quality_star__in", 2,3)
	}
	// 3:差评 一星级以下
	if cmt_status ==2 {
		query =  query.Filter("quality_star__lte", 1)
	}
	total, _ := query.Count()
	_, err := query.Limit(page_size, offset).All(&gct_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return gct_list, total, nil
}

//评价类型统计
func GetGoodsCommentStar(end time.Time,merchant_id int64) (bad,mid,good int64){
  cond := orm.NewCondition()
  var err error
  condExprBad := cond.And("quality_star__lt",3).Or("service_star__lt",3).Or("trade_star__lt",3).And("created_at__lt",end).And("merchant_id",merchant_id)
  bad, err = orm.NewOrm().QueryTable(GoodsComment{}).SetCond(condExprBad).Count()
  if err != nil {
    bad = 0
  }
  // data = append(data, map[string]int64{"bad":bad})
  CondExprMid := cond.And("quality_star",3).Or("service_star",3).Or("trade_star",3).And("created_at__lt",end).And("merchant_id",merchant_id)
  mid, err = orm.NewOrm().QueryTable(GoodsComment{}).SetCond(CondExprMid).Count()
  if err != nil {
    mid = 0
  }
  CondExprGood := cond.And("quality_star__gt",3).Or("service_star__gt",3).Or("trade_star__gt",3).And("created_at__lt",end).And("merchant_id",merchant_id)
  good, err = orm.NewOrm().QueryTable(GoodsComment{}).SetCond(CondExprGood).Count()
  if err != nil {
    good = 0
  }
  return bad,mid,good
}

//根据评价类型统计数量
func (this *GoodsComment)GetGoodsCommentAll(merchant_id int64) int64 {
  total,err := orm.NewOrm().QueryTable(this).Filter("merchant_id",merchant_id).Count()
  if err != nil {
    total = 0
  }
  return total
}

type StarTotal struct {
  Total           int64           `json:"total"`
}

// ty 1=>质量评价 2=>服务评价 3=>交易评价
func (this *GoodsComment)GetGoodsCommentStars(merchant_id int64,ty int8) int64 {
  var data StarTotal
  switch ty {
  case 1:
    err := orm.NewOrm().Raw("select sum(quality_star) total from goods_order where merchant_id = ?",merchant_id).QueryRow(&data)
    if err != nil {
      data = StarTotal{}
    }
  case 2:
    err := orm.NewOrm().Raw("select sum(service_star) total from goods_order where merchant_id = ?",merchant_id).QueryRow(&data)
    if err != nil {
      data = StarTotal{}
    }
  case 3:
    err := orm.NewOrm().Raw("select sum(trade_star) total from goods_order where merchant_id = ?",merchant_id).QueryRow(&data)
    if err != nil {
      data = StarTotal{}
    }
  }
  return data.Total
}

//根据情况查询列表
//ty  好评=>1 中评 =>2 差评=>3
//func (this *GoodsComment) GetListByStar(merchant_id int64,page, pageSize int,ty int8)([]*comment.CommentListJoinRep,int,error) {
//  var data []*comment.CommentListJoinRep
//  var total int64
//  switch ty {
//  case 1:
//    total,err := orm.NewOrm().Raw("select * from goods_comment as t0 inner join merchant as t1 on t1.id = t0.merchant_id inner join user as t2 on t2.id = t0.user_id inner join goods as t3 on t3.id = t0.goods_id where merchant_id = ? and quality_star > 3 or service_star > 3 or trade_star > 3 order by created_at desc limit ?,?",merchant_id,page,pageSize).QueryRows(&data)
//    if err != nil {
//      return nil, types.SystemDbErr, errors.New("查询数据库失败")
//    }
//    return data, int(total), nil
//  case 2:
//    total,err := orm.NewOrm().Raw("select * from goods_comment as t0 inner join merchant as t1 on t1.id = t0.merchant_id inner join user as t2 on t2.id = t0.user_id inner join goods as t3 on t3.id = t0.goods_id where merchant_id = ? and quality_star = 3 or service_star = 3 or trade_star = 3 order by created_at desc limit ?,?",merchant_id,page,pageSize).QueryRows(&data)
//    if err != nil {
//      return nil, types.SystemDbErr, errors.New("查询数据库失败")
//    }
//    return data, int(total), nil
//  case 3:
//    total,err := orm.NewOrm().Raw("select * from goods_comment as t0 inner join merchant as t1 on t1.id = t0.merchant_id inner join user as t2 on t2.id = t0.user_id inner join goods as t3 on t3.id = t0.goods_id where merchant_id = ? and quality_star < 3 or service_star < 3 or trade_star < 3 order by created_at desc limit ?,?",merchant_id,page,pageSize).QueryRows(&data)
//    if err != nil {
//      return nil, types.SystemDbErr, errors.New("查询数据库失败")
//    }
//    return data, int(total), nil
//  }
//  return data, int(total), nil
//}

