package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/form_validate"
  "blockshop/models"
  "github.com/astaxie/beego/orm"
  "net/url"
)

type MerchantService struct {
  BaseService
}

func (self *MerchantService) GetPaginateData(listRows int, params url.Values) ([]*models.Merchant, beego_pagination.Pagination) {
  //搜索、查询字段赋值
  self.SearchField = append(self.SearchField, new(models.Merchant).SearchField()...)

  var merchant []*models.Merchant
  o := orm.NewOrm().QueryTable(new(models.Merchant))
  _, err := self.PaginateAndScopeWhere(o, listRows, params).All(&merchant)
  if err != nil {
    return nil, self.Pagination
  } else {
    return merchant, self.Pagination
  }
}

func (*MerchantService) Create(form *form_validate.MerchantForm) int {
  merchant := models.Merchant{
    MerchantName: form.MerchantName,
    MerchantIntro: form.MerchantIntro,
    MerchantDetail: form.MerchantDetail,
    ContactUser: form.ContactUser,
    Phone: form.Phone,
    WeChat: form.WeChat,
    Address: form.Address,
    ShopLevel: form.ShopLevel,
    ShopServer: form.ShopServer,
    MerchantWay: form.MerchantWay,
    Logo:form.Logo,
  }
  id, err := orm.NewOrm().Insert(&merchant)

  if err == nil {
    return int(id)
  } else {
    return 0
  }
}

func (*MerchantService) GetMerchantById(id int64) *models.Merchant {
  o := orm.NewOrm()
  merchant := models.Merchant{Id: id}
  err := o.Read(&merchant)
  if err != nil {
    return nil
  }
  return &merchant
}

func (*MerchantService) IsExistName(merchant_name string, id int64) bool {
  if id == 0 {
    return orm.NewOrm().QueryTable(new(models.Merchant)).Filter("merchant_name", merchant_name).Exist()
  } else {
    return orm.NewOrm().QueryTable(new(models.Merchant)).Filter("merchant_name", merchant_name).Exclude("id", id).Exist()
  }
}

func (*MerchantService) Update(form *form_validate.MerchantForm) int{
  o := orm.NewOrm()
  merchant := models.Merchant{Id: form.Id}
  if o.Read(&merchant) == nil {
    merchant.MerchantName = form.MerchantName
    merchant.MerchantIntro = form.MerchantIntro
    merchant.MerchantDetail = form.MerchantDetail
    merchant.ContactUser = form.ContactUser
    merchant.Phone = form.Phone
    merchant.WeChat = form.WeChat
    merchant.Address = form.Address
    merchant.MerchantWay = form.MerchantWay
    if len(form.Logo) > 0 {
      merchant.Logo = form.Logo
    }
    num, err := o.Update(&merchant)
    if err == nil {
      return int(num)
    } else {
      return 0
    }
  }
  return 0
}

func (*MerchantService) Del(ids []int) int{
  count, err := orm.NewOrm().QueryTable(new(models.Merchant)).Filter("id__in", ids).Delete()
  if err == nil {
    return int(count)
  } else {
    return 0
  }
}

func (*MerchantService) GetMerchants() []*models.Merchant {
  var merchants []*models.Merchant
  _,err := orm.NewOrm().QueryTable(new(models.Merchant)).All(&merchants)
  if err != nil {
    return nil
  }
  return merchants
}
