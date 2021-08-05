package models

import (
	"blockshop/common"
	"blockshop/types"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type GoodsComment struct {
	BaseModel
	Id           int64     `orm:"pk;column(id);auto;size(11)" description:"评论ID" json:"id"`
	GoodsId      int64     `orm:"column(goods_id);size(64)" description:"评论商品" json:"goods_id"`
	UserId       int64     `orm:"column(user_id);default(1);" description:"评论人" json:"user_id"`
	Star         int8      `orm:"column(star);default(5);index" description:"评论星级" json:"star"`
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
	query := orm.NewOrm().QueryTable(GoodsComment{}).Filter("GoodsId", goods_id)
	total, _ := query.Count()
	_, err := query.Limit(pageSize, offset).All(&gct_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return gct_list, total, nil
}