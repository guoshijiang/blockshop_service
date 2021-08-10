package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type Asset struct {
	BaseModel
	Id        int64    `orm:"pk;column(id);auto;size(11)" description:"币种ID" json:"id"`
	Name      string   `orm:"column(name);unique;size(11);index" description:"币种名称" json:"name"`
	ChainName string   `orm:"column(chain_name);unique;size(11);index" description:"链名称" json:"chain_name"`
	UsdPrice  string   `orm:"column(usd_price)" description:"币种的美元价格" json:"usd_price"`
	CnyPrice  string   `orm:"column(usd_price)" description:"币种的人民币价格" json:"cny_price"`
	Unit      int64    `orm:"column(unit);default(8)" description:"币种精度" json:"unit"`
}

func (this *Asset) TableName() string {
	return common.TableName("asset")
}

func (this *Asset) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *Asset) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *Asset) SearchField() []string {
  return []string{"name"}
}


func (a *Asset) GetAsset() (*Asset, error) {
	var asset Asset
	err := asset.Query().Filter("name", a.Name).One(&asset)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func (a *Asset) GetAssetById() (*Asset, error) {
	var asset Asset
	err := asset.Query().Filter("id", a.Id).One(&asset)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func (a *Asset) GetAssetList() ([]*Asset, error) {
	var asset_list []*Asset
	_, err := orm.NewOrm().QueryTable(a.TableName()).RelatedSel().All(&asset_list)
	return asset_list, err
}

func (a *Asset) GetAssetByName() (*Asset, error) {
	var asset Asset
	err := asset.Query().Filter("name", a.Name).One(&asset)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}


func (a *Asset) PageList(page,pageSize int,condition *orm.Condition) ([]*Asset,int64) {
	offset := (page - 1) * pageSize
	list := make([]*Asset, 0)
	query := orm.NewOrm().QueryTable(a.TableName())
	query = query.SetCond(condition)
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize,offset).RelatedSel().All(&list)
	return list, total
}

func (a *Asset) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

type AssetValues struct {
	Total int
	Found int
}

func (a *Asset) AssetList() (orm.Params,error) {
	var data  orm.Params
	_,err := orm.NewOrm().Raw("select name,id from asset").RowsToMap(&data,"name","id")
	return data,err
}