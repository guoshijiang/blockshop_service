package services

import (
  "blockshop/common"
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/form_validate"
  "blockshop/models"
  "blockshop/types"
  user2 "blockshop/types/user"
  "errors"
  "github.com/astaxie/beego/orm"
  "net/url"
  "strconv"
  "strings"
)

type UserService struct {
  BaseService
}

func (self *UserService) GetPaginateData(listRows int, params url.Values) ([]*models.User, beego_pagination.Pagination) {
  var err error
  //搜索、查询字段赋值
  self.SearchField = append(self.SearchField, new(models.User).SearchField()...)

  var data []*models.User
  o := orm.NewOrm().QueryTable(new(models.User))
  vStr := params.Get("verify")
  vInt,_ := strconv.Atoi(vStr)
  if vInt == 1 {
    _, err = self.PaginateAndScopeWhere(o, listRows, params).Filter("is_auth__gt", 0).All(&data)
  } else {
    _, err = self.PaginateAndScopeWhere(o, listRows, params).All(&data)
  }
  if err != nil {
    return nil, self.Pagination
  } else {
    return data, self.Pagination
  }
}

//func (Self *UserService) GetPaginateDataAccount(listRows int, params url.Values) ([]*models.UserIncomeAccountData, beego_pagination.Pagination) {
//  var data []*models.UserIncomeAccountData
//  var total int64
//  om := orm.NewOrm()
//  inner := "from  user_income_account as t0 inner join user as t1 on t1.id = t0.user_id  inner join asset as t2 on t2.id = t0.asset_id where t0.id > 0 "
//  sql := "select t0.*,t1.user_name,t2.asset_name " + inner
//  sql1 := "select count(*) total " + inner
//
//  //搜索、查询字段赋值
//  Self.SearchField = append(Self.SearchField, new(models.UserIncomeAccount).SearchField()...)
//  where,param := Self.ScopeWhereRaw(params)
//  Self.PaginateRaw(listRows,params)
//
//  if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  Self.Pagination.Total = int(total)
//  param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
//  if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  return data,Self.Pagination
//}

//func (Self *UserService) GetPaginateDataWallet(listRows int, params url.Values)  ([]*models.UserWalletData, beego_pagination.Pagination)  {
//  var data []*models.UserWalletData
//  var total int64
//  om := orm.NewOrm()
//  inner := "from  user_wallet as t0 inner join user as t1 on t1.user_uuid = t0.user_uuid inner join asset t2 on t2.id = t0.asset_id where t0.id > 0 "
//  sql := "select t0.*,t1.user_name,t2.asset_name " + inner
//  sql1 := "select count(*) total " + inner
//
//  if len(params.Get("t1.id"))  > 0 {
//    Self.WhereField = append(Self.WhereField,"t1.id")
//  }
//  //搜索、查询字段赋值
//  where,param := Self.ScopeWhereRaw(params)
//  Self.PaginateRaw(listRows,params)
//
//  if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  Self.Pagination.Total = int(total)
//  param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
//  if _,err := om.Raw(sql+where+" order by t0.created_at desc limit ?,?",param).QueryRows(&data);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  return data,Self.Pagination
//}

//func (Self *UserService) GetPaginateDataPosition(listRows int, params url.Values)  ([]*models.UserPosition, beego_pagination.Pagination)  {
//
//  //搜索、查询字段赋值
//  Self.SearchField = append(Self.SearchField, new(models.UserPosition).SearchField()...)
//
//  Self.WhereField = append(Self.WhereField,[]string{"user_id","is_open"}...)
//  var data []*models.UserPosition
//  o := orm.NewOrm().QueryTable(new(models.UserPosition))
//  _, err := Self.PaginateAndScopeWhere(o, listRows, params).All(&data)
//  if err != nil {
//    return nil, Self.Pagination
//  } else {
//    return data, Self.Pagination
//  }
//
//  //var data []*models.UserPositionData
//  //var total int64
//  //om := orm.NewOrm()
//  //inner := "from  user_position as t0 inner join btc_fund_base as t1 on t1.id = t0.fund_id where  "
//  //sql := "select t0.*,t1.fund_name " + inner
//  //sql1 := "select count(*) total " + inner
//  //
//  //if len(params.Get("user_id"))  > 0 {
//  //	Self.WhereField = append(Self.WhereField,"user_id")
//  //}
//  ////搜索、查询字段赋值
//  //where,param := Self.ScopeWhereRaw(params)
//  //Self.PaginateRaw(listRows,params)
//  //
//  //if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
//  //	return nil,beego_pagination.Pagination{}
//  //}
//  //Self.Pagination.Total = int(total)
//  //param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
//  //if _,err := om.Raw(sql+where+" order by t0.created_at desc limit ?,?",param).QueryRows(&data);err != nil {
//  //	return nil,beego_pagination.Pagination{}
//  //}
//  //return data,Self.Pagination
//}

//func (Self *UserService) GetPaginateDataSource(listRows int, params url.Values) ([]*models.UserIncomeSourceData, beego_pagination.Pagination) {
//  var data []*models.UserIncomeSourceData
//  var total int64
//  om := orm.NewOrm()
//  inner := "from  user_income_source as t0 inner join user as t1 on t1.id = t0.user_id  where t0.id > 0 "
//  if len(params.Get("user_id")) > 0 {
//    inner += " and user_id = " + params.Get("user_id")
//  }
//  sql := "select t0.*,t1.user_name " + inner
//  sql1 := "select count(*) total " + inner
//
//  //搜索、查询字段赋值
//  Self.SearchField = append(Self.SearchField, new(models.UserIncomeSource).SearchField()...)
//
//  where,param := Self.ScopeWhereRaw(params)
//  Self.PaginateRaw(listRows,params)
//
//  if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  Self.Pagination.Total = int(total)
//  param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
//  if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  return data,Self.Pagination
//}


func (*UserService) IsExistName(user_name string, id int64) bool {
  if id == 0 {
    return orm.NewOrm().QueryTable(new(models.User)).Filter("user_name", user_name).Exist()
  } else {
    return orm.NewOrm().QueryTable(new(models.User)).Filter("user_name", user_name).Exclude("id", id).Exist()
  }
}


// 新增用户
func (*UserService) Create(form *form_validate.UserForm) int {
 registerParam := user2.Register{UserName:form.UserName,Password: form.Password,ConfirmPassword: form.Password,PinCode: form.PinCode}
 ok,_:= models.UserRegister(registerParam)
 if ok ==  types.ReturnSuccess {
   return 1
 }
 return 0
}

//根据id获取一条user数据
func (*UserService) GetRowById(id int64) *models.User {
  o := orm.NewOrm()
  user := models.User{Id: id}
  err := o.Read(&user)
  if err != nil {
    return nil
  }
  return &user
}

//func (*UserService) GetUserWalletById(id int64) *models.UserWallet {
//  o := orm.NewOrm()
//  wallet := models.UserWallet{Id: id}
//  err := o.Read(&wallet)
//  if err != nil {
//    return nil
//  }
//  return &wallet
//}


//更新钱包
//func (*UserService) UpdateWallet(form *form_validate.UserWalletForm) int {
//  o := orm.NewOrm()
//  wallet := models.UserWallet{Id: form.Id}
//  if o.Read(&wallet) == nil {
//    wallet.TotalAmount = form.TotalAmount
//    num, err := o.Update(&wallet)
//    if err == nil {
//      return int(num)
//    } else {
//      return 0
//    }
//  }
//  return 0
//}

//查看密码是否相等
func (*UserService) CheckPassword(form *form_validate.UserForm,ty int) error {
  o := orm.NewOrm()
  user := models.User{Id: form.Id}
  if o.Read(&user) != nil {
    return errors.New("用户不存在")
  }
  if ty == 1 {
    if user.FundPassword == common.ShaOne(form.FundPassword) {
      return errors.New("新密码和原密码一样")
    }
  } else {
    if user.Password == common.ShaOne(form.Password) {
      return errors.New("新密码和原密码一样")
    }
  }
  return nil
}


//更新用户
func (*UserService) Update(form *form_validate.UserForm) int {
  o := orm.NewOrm()
  user := models.User{Id: form.Id}
  if o.Read(&user) == nil {
    if len(form.Password) > 0 {
      user.Password = common.ShaOne(form.Password)
    }
    if len(form.FundPassword) > 0 {
      user.FundPassword = common.ShaOne(form.FundPassword)
    }
    dateTimeStr := strings.Split(form.TimeRange,"-")
    if len(dateTimeStr) > 1 {
      //startTime,_ := time.ParseInLocation("2006-01-02",dateTimeStr[0],time.Local)
      //endTime,_ := time.ParseInLocation("2006-01-02",dateTimeStr[1],time.Local)
      //user.StartTime = startTime.String()
      //user.EndTime = endTime.String()
    }
    user.UserName = form.UserName
    //user.Phone = form.Phone
    //user.Email = form.Email
    //if form.IsAuth > 0 {
    //  user.IsAuth  = form.IsAuth
    //}
    //if form.IsAuth > 1 {
    //  user.Reason = form.Reason
    //}
    num, err := o.Update(&user)
    if err == nil {
      return int(num)
    } else {
      return 0
    }
  }
  return 0
}


//更新用户
//func (*UserService) UpdateAuth(form *form_validate.VerifyForm) int {
//  o := orm.NewOrm()
//  user := models.User{Id: form.Id}
//  if o.Read(&user) == nil {
//    if form.Reject > 0 {
//      user.IsAuth = 2
//    } else {
//      user.IsAuth = 3
//      user.Reason = form.Reason
//    }
//    num, err := o.Update(&user)
//    if err == nil {
//      return int(num)
//    } else {
//      return 0
//    }
//  }
//  return 0
//}
//仓位解禁
//func (*UserService) Release(id int64) int{
//  o := orm.NewOrm()
//  p := models.UserPosition{Id: id}
//  if o.Read(&p) == nil {
//    p.IsForbiden = 1
//    num, err := o.Update(&p)
//    if err == nil {
//      return int(num)
//    } else {
//      return 0
//    }
//  }
//  return 0
//}

//删除
func (*UserService) Del(ids []int) int {
  count, err := orm.NewOrm().QueryTable(new(models.User)).Filter("id__in", ids).Delete()
  if err == nil {
    return int(count)
  } else {
    return 0
  }
}

//获得直推用户
//func (*UserService) DirectList(ids ...int64) []int64{
//  var (
//    list orm.ParamsList
//    data []int64
//  )
//  data = make([]int64,0)
//  if len(ids) > 0 {
//    _, err := orm.NewOrm().QueryTable(new(models.User)).Filter("invite_me_user_id__in", ids).ValuesFlat(&list, "id")
//    if err != nil {
//      fmt.Println("err----",err)
//    } else {
//      for _,v := range list {
//        data = append(data,reflect.ValueOf(v).Interface().(int64))
//      }
//    }
//  }
//  return data
//}

//代理列表
//func (Self *UserService) AgentList(listRows int, params url.Values) ([]*models.User, beego_pagination.Pagination) {
//  //搜索、查询字段赋值
//  Self.SearchField = append(Self.SearchField, new(models.User).SearchField()...)
//  if len(params.Get("id__in")) > 0 {
//    Self.WhereField = append(Self.WhereField, "id__in")
//  }
//
//  var data []*models.User
//  o := orm.NewOrm().QueryTable(new(models.User))
//  _, err := Self.PaginateAndScopeWhere(o, listRows, params).All(&data)
//  if err != nil {
//    return nil, Self.Pagination
//  } else {
//    return data, Self.Pagination
//  }
//}

//代理入金列表
//func (Self *UserService) GetPaginateDataDeposit(listRows int, params url.Values) ([]*models.WalletDepositData, beego_pagination.Pagination) {
//  var data []*models.WalletDepositData
//  var total int64
//  om := orm.NewOrm()
//  inner := "from  wallet_deposit as t0 inner join user as t1 on t1.user_uuid = t0.user_uuid  inner join asset as t2 on t2.id = t0.asset_id where t0.id > 0 "
//  sql := "select t0.*,t1.user_name,t1.phone,t1.email,t2.asset_name " + inner
//  sql1 := "select count(*) total " + inner
//
//  //搜索、查询字段赋值
//  Self.SearchField = append(Self.SearchField, new(models.WalletDeposit).SearchField()...)
//  Self.WhereField = append(Self.WhereField,"t1.id")
//  where,param := Self.ScopeWhereRaw(params)
//  Self.PaginateRaw(listRows,params)
//
//  if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  Self.Pagination.Total = int(total)
//  param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
//  if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  return data,Self.Pagination
//}

//用户钱包地址
//func (Self *UserService) GetPaginateDataWalletAddress(listRows int, params url.Values)  ([]*models.WalletAddressData, beego_pagination.Pagination)  {
//  var data []*models.WalletAddressData
//  var total int64
//  om := orm.NewOrm()
//  inner := "from  wallet_address as t0 inner join user as t1 on t1.user_uuid = t0.user_uuid inner join asset t2 on t2.id = t0.asset_id where t0.id > 0 "
//  sql := "select t0.*,t1.user_name,t2.asset_name " + inner
//  sql1 := "select count(*) total " + inner
//
//  if len(params.Get("t1.id"))  > 0 {
//    Self.WhereField = append(Self.WhereField,"t1.id")
//  }
//  //搜索、查询字段赋值
//  where,param := Self.ScopeWhereRaw(params)
//  Self.PaginateRaw(listRows,params)
//
//  if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  Self.Pagination.Total = int(total)
//  param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
//  if _,err := om.Raw(sql+where+" order by t0.created_at desc limit ?,?",param).QueryRows(&data);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  return data,Self.Pagination
//}

//func (*UserService) GetUserWalletAddressById(id int64) *models.WalletAddress {
//  o := orm.NewOrm()
//  addr := models.WalletAddress{Id: id}
//  err := o.Read(&addr)
//  if err != nil {
//    return nil
//  }
//  return &addr
//}


//更新用户
//func (*UserService) UpdateWalletAddress(form *form_validate.WalletAddressForm) int {
//  o := orm.NewOrm()
//  wallet := models.WalletAddress{Id: form.Id}
//  if o.Read(&wallet) == nil {
//    wallet.Address = form.Address
//    num, err := o.Update(&wallet)
//    if err == nil {
//      return int(num)
//    } else {
//      return 0
//    }
//  }
//  return 0
//}