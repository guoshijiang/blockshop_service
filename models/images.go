package models

import (
	"blockshop/common"
	"blockshop/types"
	"github.com/astaxie/beego/orm"
)

type ImageFile struct {
	BaseModel
	Id        int64     `orm:"pk;column(id);auto;size(11)" description:"图片ID" json:"id"`
	Url       string    `orm:"column(url);size(256);index" description:"图片Url"  json:"url"`
	ImgType   int8      `orm:"column(img_type)" description:"图片类别" json:"img_type"`  // 0:用户头像  1:商品评论图片
}

func (this *ImageFile) TableName() string {
	return common.TableName("user_image")
}

func (this *ImageFile) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this);
	if err != nil {
		return err, 0
	}
	return nil, id
}

func (this *ImageFile) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *ImageFile) GetImageById(id int64) (*ImageFile, int, error) {
	var image ImageFile
	err := image.Query().Filter("Id", id).One(&image)
	if err != nil {
		return nil, 100, err
	}
	return &image, types.ReturnSuccess, nil
}
