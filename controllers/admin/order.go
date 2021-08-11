package controllers

import (
  "blockshop/global"
  "blockshop/global/response"
  "blockshop/models"
  "blockshop/services"
  "strconv"
)

type OrderController struct {
  baseController
}

//订单列表
func (Self *OrderController) Index() {
  var srv services.OrderService
  data, pagination := srv.GetPaginateDataRaw(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "order/index.html"
}

//订单编辑
func (Self *OrderController) Edit() {
  id, _ := Self.GetInt64("id", -1)
  if id <= 0 {
    response.ErrorWithMessage("Param is error.", Self.Ctx)
  }
  var srv services.OrderService
  var addr *models.UserAddress

  data := srv.GetOrderById(id)
  if data == nil {
    response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
  }
  //用户信息

  if data != nil && data.AddressId > 0 {
    addr,_,_ = (&models.UserAddress{Id: data.AddressId}).GetAddressById()
  }
  if addr == nil {
    addr = new(models.UserAddress)
  }
  Self.Data["data"] = data
  Self.Data["addr"] = addr
  Self.Layout = "public/base.html"
  Self.TplName = "order/detail.html"
}



func (Self *OrderController) Update() {
  id,_ := Self.GetInt64("id",-1)
  ship_number := Self.GetString("ship_number","")
  if id < 0 {
    response.ErrorWithMessage("订单不存在",Self.Ctx)
  }
  order := models.GoodsOrder{Id: id,ShipNumber: ship_number}
  if new(services.OrderService).UpdateShipNumber(&order) > 0 {
    response.Success(Self.Ctx)
  }
  response.Error(Self.Ctx)
}


func (Self *OrderController) Del() {
  idStr := Self.GetString("id")
  ids := make([]int, 0)
  var idArr []int

  if idStr == "" {
    Self.Ctx.Input.Bind(&ids, "id")
  } else {
    id, _ := strconv.Atoi(idStr)
    idArr = append(idArr, id)
  }

  if len(ids) > 0 {
    idArr = ids
  }
  if len(idArr) == 0 {
    response.ErrorWithMessage("参数id错误.", Self.Ctx)
  }
  var orderService services.OrderService
  if orderService.Del(idArr) > 0 {
    response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

//退货列表
func (Self *OrderController) Process() {
  var orderService services.OrderService
  data, pagination := orderService.GetPaginateProcessDataRaw(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "order_process/index.html"
}

//退货单详情
func (Self *OrderController) Detail() {
  id,_ := Self.GetInt64("id",0)
  data,_,_ := models.GetOrderProcessDetailById(id)
  addr,_,_ := (&models.UserAddress{Id: data.AddressId}).GetAddressById()
  merchant,_,_ := models.GetMerchantDetail(data.MerchantId)
  order,_,_ := models.GetGoodsOrderDetail(data.OrderId)
  goods,_,_ := models.GetGoodsDetail(data.GoodsId)

  Self.Data["order"] = order
  Self.Data["goods"] = goods
  Self.Data["merchant"] = merchant
  Self.Data["addr"] = addr
  Self.Data["data"] = data
  Self.Layout = "public/base.html"
  Self.TplName = "order_process/detail.html"
}

//退货审核
func (Self *OrderController) Verify() {
  id,_ := Self.GetInt64("id")
  raft,_ := Self.GetInt("raft")
  var mc *models.OrderProcess
  if raft == 0 {
    ret_pay_rs := Self.GetString("reason")
    mc = &models.OrderProcess{Id: id,Process: 2,RetPayRs: ret_pay_rs,MerchantId: int64(loginUser.Id)}
  } else {
    mc = &models.OrderProcess{Id: id,Process: 1,MerchantId: int64(loginUser.Id)}
  }
  url := global.URL_RELOAD
  if err := mc.Update([]string{"process","ret_pay_rs","merchant_id"}...);err != nil {
    response.ErrorWithMessage("审核失败", Self.Ctx)
  }
  response.SuccessWithMessageAndUrl("审核成功", url, Self.Ctx)
}