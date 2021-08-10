package models

import (
	"blockshop/common"
	"blockshop/types"
	"blockshop/types/goods"
	"blockshop/types/merchant"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Goods struct {
	BaseModel
	Id             int64     `orm:"pk;column(id);auto;size(11)" description:"商家ID" json:"id"`
	UserId         int64     `orm:"column(user_id)" description:"商家对应的用户ID" json:"user_id"`
	GoodsCatId     int64     `orm:"column(goods_cat_id)" description:"商品分类ID" json:"goods_cat_id"`
	GoodsTypeId    int64     `orm:"column(goods_type_id)" description:"商品大类ID" json:"goods_type_id"`
	GoodsMark      string    `orm:"column(goods_mark);size(512);index" description:"商品备注" json:"goods_mark"`
	Serveice       string    `orm:"column(serveice);size(512);index" description:"服务说明" json:"serveice"`
	CalcWay        int8      `orm:"column(calc_way);default(0);index" description:"计量方式" json:"calc_way"`
	MerchantId     int64     `orm:"column(merchant_id);default(0);index" description:"商品所属商家ID" json:"merchant_id"`
	Title          string    `orm:"column(title);size(512)" description:"商品标题" json:"title"`
	Logo           string    `orm:"column(logo);size(150);default(/static/upload/default/user-default-60x60.png)" description:"商品封面" json:"logo"`
	OriginStateId  int64     `orm:"column(origin_state_id);default(0)" description:"商品的产地" json:"origin_state_id"`
	TotalAmount    int64     `orm:"column(total_amount);default(150000)" description:"商品总量" json:"total_amount"`
	LeftAmount     int64     `orm:"column(left_amount);default(150000)" description:"剩余商品总量" json:"left_amount"`
	GoodsPrice     float64   `orm:"column(goods_price);default(1);digits(22);decimals(8)" description:"商品价格" json:"goods_price"`
	GoodsDisPrice  float64   `orm:"column(goods_discount_price);default(1);digits(22);decimals(8)" description:"商品折扣价格" json:"goods_discount_price"`
	GoodsName      string    `orm:"column(goods_name);size(512);index" description:"产品名称" json:"goods_name"`
	GoodsParams    string    `orm:"column(goods_params);type(text)" description:"产品参数" json:"goods_params"`
	GoodsDetail    string    `orm:"column(goods_detail);type(text)" description:"产品详细介绍" json:"goods_detail"`
	Discount       float64   `orm:"column(discount);default(0);index" description:"折扣" json:"discount"`              // 取值 0.1-9.9；0代表不打折
	Sale           int8      `orm:"column(sale);default(0);index" description:"上架下架" json:"sale"`                   // 0:上架 1:下架
	IsDisplay      int8      `orm:"column(is_display);default(0);index" description:"首页展示" json:"is_display"`       // 0:首页不展示, 1:首页展示
	SellNums       int64     `orm:"column(sell_nums);default(0);index" description:"售出数量" json:"sell_nums"`
	IsHot          int8      `orm:"column(is_hot);default(0);index" description:"爆款产品" json:"is_hot"`                // 0:非爆款产品 1:爆款产品
	IsDiscount     int8      `orm:"column(is_discount);default(0);index" description:"打折活动" json:"is_discount"`      // 0:不打折，1:打折活动产品
	LeftTime       int64     `orm:"column(left_time);default(0);index" description:"限时产品剩余时间" json:"left_time"`   // 限时产品剩余时间
	IsLimitTime    int8      `orm:"column(is_limit_time);default(0);index" description:"限时产品" json:"is_limit_time"`  // 0:不是限时产品 1:是限时
}

type Select struct {
  Id			int				`json:"id"`
  Name		string				`json:"name"`
}

func (this *Goods) TableName() string {
	return common.TableName("goods")
}

func (this *Goods) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (*Goods) SearchField() []string {
	return []string{"goods_name"}
}

func (this *Goods) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *Goods) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *Goods) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *Goods) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func GetGoodsList(req goods.GoodsListReq) ([]*Goods, int64, error) {
	offset := (req.Page - 1) * req.PageSize
	goods_list := make([]*Goods, 0)
	query_good := orm.NewOrm().QueryTable(Goods{}).Filter("IsRemoved", 0)
	if req.GoodsName != "" {
		query_good = query_good.Filter("goods_name__contains", req.GoodsName)
	}
	if req.TypeId >= 1 {
		query_good = query_good.Filter("goods_type_id", req.TypeId)
	}
	if req.CatId >= 1 {
		query_good = query_good.Filter("goods_cat_id", req.CatId)
	}
	if req.StartPrice >= 0 && req.EndPrice != 0 && req.EndPrice >= req.StartPrice {
		query_good = query_good.Filter("goods_price__gt", req.StartPrice).Filter("goods_price__lt", req.EndPrice)
	}
	// 0:时间，1:销量；2:价格; 3:商家
	if req.OrderBy >= 0 {
		if req.OrderBy == 0{
			query_good = query_good.OrderBy("-created_at")
		}else if req.OrderBy == 1{
			query_good = query_good.OrderBy("-sell_nums")
		}else if req.OrderBy == 2{
			query_good = query_good.OrderBy("-goods_price")
		} else {
			query_good = query_good.OrderBy("-merchant_id")
		}
	}
	if req.MerchantId > 0 {
		query_good = query_good.Filter("merchant_id", req.MerchantId)
	}
	if req.OriginStateId >= 1 {
		query_good = query_good.Filter("origin_state_id", req.OriginStateId)
	}
	total, _ := query_good.Count()
	_, err := query_good.Limit(req.PageSize, offset).All(&goods_list)
	if err != nil {
		return nil, 0, errors.New("查询数据库失败")
	}
	return goods_list, total, nil
}

func GetGoodsDetail(id int64) (*Goods, int, error) {
	var gds Goods
	if err := orm.NewOrm().QueryTable(Goods{}).Filter("Id", id).RelatedSel().One(&gds); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &gds, types.ReturnSuccess, nil
}

func GetMerchantGoodsNums(metchant_id int64) int64 {
	total, err := orm.NewOrm().QueryTable(Goods{}).Filter("MerchantId", metchant_id).Count()
	if err != nil {
		return 0
	}
	return total
}

// 添加商品
func CreateMerchantGoods(gds_param merchant.MerchantAddUpdGoodsReq) (int, error) {
	var gds_cat_id int64
	if gds_param.GoodsCatName != "" {
		gds_cat := GetGdsCatByName(gds_param.GoodsCatName)
		if gds_cat == nil {
			gs_c := GoodsCat {
				CatLevel: 0,
				FatherCatId: 0,
				Name: gds_param.GoodsCatName,
			}
			_, id := gs_c.Insert()
			gds_cat_id = id
		} else {
			gds_cat_id = gds_cat.Id
		}
	}
	var gds_type_id int64
	if gds_param.GoodsTypeName != "" {
		gds_type := GetGdsCatByName(gds_param.GoodsTypeName)
		if gds_type == nil {
			gs_t := GoodsType {
				Name: gds_param.GoodsTypeName,
			}
			_, id := gs_t.Insert()
			gds_type_id = id
		} else {
			gds_type_id = gds_type.Id
		}
	}
	var os_state_id int64
	if gds_param.OriginStName != "" {
		os_state := GetGdsOsByName(gds_param.GoodsTypeName)
		if os_state == nil {
			gds_os := GoodsOriginState{
				Name: gds_param.OriginStName,
			}
			_, id := gds_os.Insert()
			os_state_id = id
		} else {
			os_state_id = os_state.Id
		}
	}
	create_gds := Goods {
		UserId: gds_param.UserId,
		GoodsCatId: gds_cat_id,
		GoodsTypeId: gds_type_id,
		GoodsMark: gds_param.GoodsMark,
		Serveice: gds_param.Serveice,
		CalcWay: gds_param.CalcWay,
		MerchantId: gds_param.MerchantId,
		Title: gds_param.Title,
		Logo: gds_param.Logo,
		OriginStateId: os_state_id,
		TotalAmount: gds_param.TotalAmount,
		GoodsPrice : gds_param.GoodsPrice,
		GoodsDisPrice: gds_param.GoodsPrice * gds_param.Discount,
		GoodsName: gds_param.GoodsCatName,
		GoodsParams: gds_param.GoodsParams,
		GoodsDetail: gds_param.GoodsDetail,
		Discount: gds_param.Discount,
		Sale: gds_param.Sale,
		IsDiscount: gds_param.IsDiscount,
	}
	err, goods_id := create_gds.Insert()
	if err != nil {
		return types.SystemDbErr, errors.New("创建商品记录失败")
	}
	if gds_param.GoodsAttrValue != "nil" {
		gds_attr := GoodsAttr{
			GoodsId: goods_id,
			TypeKey: gds_param.GoodsAttrKey,
			TypeVale: gds_param.GoodsAttrValue,
		}
		err, _ :=gds_attr.Insert()
		if err != nil {
			return types.SystemDbErr, errors.New("创建商品属性失败")
		}
	}
	if gds_param.GoodsImgIds != nil {
		for _, item := range gds_param.GoodsImgIds {
			img := GetImageById(item)
			if img != nil {
				gds_img := GoodsImage{
					GoodsId: goods_id,
					Image: img.Url,
				}
				err, _ := gds_img.Insert()
				if err != nil {
					return types.SystemDbErr, errors.New("处理商品图片失败")
				}
			}
		}
	}
	return types.ReturnSuccess, errors.New("创建商品记录成功")
}

// 修改商品信息
func UpdateMerchantGoods(gds_param merchant.UpdateGoodsReq) (int, error) {
	var gds Goods
	if err := orm.NewOrm().QueryTable(Goods{}).Filter("id", gds_param.GoodsId).RelatedSel().One(&gds); err != nil {
		return types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	if gds_param.GoodsCatName != "" {
		var gds_cat_id int64
		if gds_param.GoodsCatName != "" {
			gds_cat := GetGdsCatByName(gds_param.GoodsCatName)
			if gds_cat == nil {
				gs_c := GoodsCat {
					CatLevel: 0,
					FatherCatId: 0,
					Name: gds_param.GoodsCatName,
				}
				_, id := gs_c.Insert()
				gds_cat_id = id
			} else {
				gds_cat_id = gds_cat.Id
			}
		}
		gds.GoodsCatId = gds_cat_id
	}
	var gds_type_id int64
	if gds_param.GoodsTypeName != "" {
		gds_type := GetGdsCatByName(gds_param.GoodsTypeName)
		if gds_type == nil {
			gs_t := GoodsType {
				Name: gds_param.GoodsTypeName,
			}
			_, id := gs_t.Insert()
			gds_type_id = id
		} else {
			gds_type_id = gds_type.Id
		}
		gds.GoodsTypeId = gds_type_id
	}
	var os_state_id int64
	if gds_param.OriginStName != "" {
		os_state := GetGdsOsByName(gds_param.GoodsTypeName)
		if os_state == nil {
			gds_os := GoodsOriginState{
				Name: gds_param.OriginStName,
			}
			_, id := gds_os.Insert()
			os_state_id = id
		} else {
			os_state_id = os_state.Id
		}
		gds.OriginStateId = os_state_id
	}
	if gds_param.GoodsMark != "" {
		gds.GoodsMark = gds_param.GoodsMark
	}
	if gds_param.Serveice != "" {
		gds.Serveice = gds_param.Serveice
	}
	if gds_param.CalcWay >= 0 {
		gds.CalcWay = gds_param.CalcWay
	}
	if gds_param.Title != "" {
		gds.Title = gds_param.Title
	}
	if  gds_param.Logo != "" {
		gds.Logo = gds_param.Logo
	}
	if  gds_param.TotalAmount > 0 {
		gds.TotalAmount = gds_param.TotalAmount
	}
	if gds_param.GoodsPrice > 0 {
		gds.GoodsPrice = gds_param.GoodsPrice
		if gds_param.Discount > 0 {
			gds.GoodsDisPrice= gds_param.GoodsPrice * gds_param.Discount
		}
	}
	if gds_param.GoodsCatName != "" {
		gds.GoodsName = gds_param.GoodsCatName
	}
	if gds_param.GoodsParams != "" {
		gds.GoodsParams = gds_param.GoodsParams
	}
	if gds_param.GoodsDetail != "" {
		gds.GoodsDetail = gds_param.GoodsDetail
	}
	if gds_param.Discount > 0 {
		gds.Discount = gds_param.Discount
	}
	if gds_param.Sale > 0 {
		gds.Sale = gds_param.Sale
	}
	if gds.IsDiscount  > 0 {
		gds.IsDiscount = gds_param.IsDiscount
	}
	if gds_param.GoodsAttrValue != "nil" {
		gds_attr := GoodsAttr{
			GoodsId: gds_param.GoodsId,
			TypeKey: gds_param.GoodsAttrKey,
			TypeVale: gds_param.GoodsAttrValue,
		}
		err, _ :=gds_attr.Insert()
		if err != nil {
			return types.SystemDbErr, errors.New("创建商品属性失败")
		}
	}
	if gds_param.GoodsImgIds != nil {
		for _, item := range gds_param.GoodsImgIds {
			img := GetImageById(item)
			if img != nil {
				gds_img := GoodsImage{
					GoodsId: gds_param.GoodsId,
					Image: img.Url,
				}
				err, _ := gds_img.Insert()
				if err != nil {
					return types.SystemDbErr, errors.New("处理商品图片失败")
				}
			}
		}
	}
	err := gds.Update()
	if err != nil {
		return types.SystemDbErr, errors.New("更新商品信息失败")
	}
	return types.ReturnSuccess, nil
}

func DeleteGoodsById(id int64) (int, error) {
	var gds Goods
	if err := orm.NewOrm().QueryTable(Goods{}).Filter("Id", id).RelatedSel().One(&gds); err != nil {
		return types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	gds.IsRemoved = 1
	err := gds.Update()
	if err != nil {
		return types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return types.ReturnSuccess, nil
}