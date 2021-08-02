package controllers

import "blockshop/services"

type WalletController struct {
  baseController
}

//用户钱包
func (Self *WalletController) User() {
  var serv services.WalletService
  data, pagination := serv.GetUserData(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "wallet/user.html"
}
//资产记录
func (Self *WalletController) Record() {
  var serv services.WalletService
  data, pagination := serv.GetRecordData(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "wallet/record.html"
}