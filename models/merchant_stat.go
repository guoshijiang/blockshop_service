package models

import (
  "blockshop/common"
  "fmt"
  "github.com/astaxie/beego/orm"
  "strconv"
)

type MerchantStat struct {
	BaseModel
	Id             int64     `orm:"pk;column(id);auto;size(11)" description:"商品属性ID" json:"id"`
	MerchantId     int64     `orm:"column(merchant_id);size(64);index" description:"商户ID" json:"merchant_id"`
	ServiceBest     int64     `orm:"column(service_best);size(64)" description:"服务好评" json:"serice_best"`
	ServiceGood    int64     `orm:"column(service_good);size(64)" description:"服务中评" json:"service_good"`
	ServiceBad     int64     `orm:"column(service_bad);size(64)" description:"服务差评" json:"service_bad"`
	ServiceAvg	   float64   `orm:"column(service_avg);size(64)" description:"服务平均星星率" json:"service_avg"`
	ServiceNum     int64     `orm:"column(service_num);size(64)" description:"服务评价星星数量" json:"service_num"`
	TradeBest 	   int64     `orm:"column(trade_best);size(64)" description:"交易好评" json:"trade_best"`
	TradeGood 	   int64     `orm:"column(trade_good);size(64)" description:"交易中评" json:"trade_good"`
	TradeBad 	   int64     `orm:"column(trade_bad);size(64)" description:"交易差评" json:"trade_bad"`
	TradeAvg       float64   `orm:"column(trade_avg);size(64)" description:"交易平均星星率" json:"trade_avg"`
	TradeNum       int64     `orm:"column(trade_num);size(64)" description:"交易评价星星数量" json:"trade_num"`
	QualityBest    int64     `orm:"column(quality_best);size(64)" description:"交易差评" json:"quality_best"`
	QualityGood    int64     `orm:"column(quality_good);size(64)" description:"交易差评" json:"quality_good"`
	QualityBad     int64     `orm:"column(quality_bad);size(64)" description:"交易差评" json:"quality_bad"`
	QualityAvg     float64   `orm:"column(quality_avg);size(64)" description:"质量平均星星率" json:"quality_avg"`
	QualityNum     int64     `orm:"column(quality_num);size(64)" description:"质量评价星星数量" json:"quality_num"`
	TotalCmtNum    int64     `orm:"column(total_cmt_num);size(64)" description:"总评价星星数量" json:"total_cmt_num"`
}

type MerchantStateCountRaw struct {
  MerchantId   int64  `json:"merchant_id"`
  QualityStar  int8   `json:"quality_star"`
  ServiceStar  int8   `json:"Service_star"`
  TradeStar    int8   `json:"trade_star"`
}


func (this *MerchantStat) TableName() string {
	return common.TableName("merchant_stat")
}

func (this *MerchantStat) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func(this *MerchantStat) Update(fields ...string) (err error,id int64) {
  if id,err = orm.NewOrm().Update(this,fields...);err != nil {
    return err,0
  }
  return nil,id
}

func(this *MerchantStat) QueryByMerchantCount() (total int64,err error) {
  total,err = orm.NewOrm().QueryTable(this).Filter("merchant_id",this.MerchantId).Count()
  return
}

func (this *MerchantStat) QueryByMerchant() (MerchantStat,error) {
  var data MerchantStat
  if err := orm.NewOrm().QueryTable(this).Filter("merchant_id",this.MerchantId).One(&data);err != nil {
    return MerchantStat{},err
  }
  return  data,nil
}



func (this *MerchantStat) UpdateByMerchant(addComment MerchantStateCountRaw)(int64,error) {
 this.MerchantId = addComment.MerchantId
 stateTotal,err := new(GoodsComment).QueryByMerchantCount(this.MerchantId)             //总评论数
 fmt.Println("---",stateTotal)
 if err != nil {
  return 0, err
 }
 stateRow,err := this.QueryByMerchant()
 if err != nil {
  return 0, err
 }

 stateRow.TotalCmtNum = int64(addComment.QualityStar+ addComment.ServiceStar+ addComment.TradeStar) //总星数
 if addComment.QualityStar < 3 {                                                                    //质量差评
  stateRow.QualityBad++
 } else if addComment.QualityStar == 3 { //质量中评
  stateRow.QualityGood++
 } else {                                    //质量差评
  stateRow.QualityBest++
 }
 stateRow.QualityNum += int64(addComment.QualityStar)            //质量总星数量
 stateRow.QualityAvg,_ = strconv.ParseFloat(fmt.Sprintf("%.4f",float64(stateRow.QualityNum)/ float64(stateTotal)),64) //质量平均星率

 if addComment.ServiceStar < 3 { //服务差评
  stateRow.ServiceBad++
 } else if addComment.QualityStar == 3 { //服务中评
  stateRow.ServiceGood++
 } else {                                    //服务差评
  stateRow.ServiceBest++
 }
 stateRow.ServiceNum += int64(addComment.QualityStar)            //服务总星数量
 stateRow.ServiceAvg,_ = strconv.ParseFloat(fmt.Sprintf("%.4f",float64(stateRow.ServiceNum)/ float64(stateTotal)),64) //服务平均星率


 if addComment.TradeStar < 3 { //交易差评
  stateRow.TradeBad++
 } else if addComment.QualityStar == 3 { //交易中评
  stateRow.TradeGood++
 } else {                                    //交易差评
  stateRow.TradeBest++
 }
 stateRow.TradeNum += int64(addComment.TradeStar)            //交易总星数量

 stateRow.TradeAvg,_ = strconv.ParseFloat(fmt.Sprintf("%.4f",float64(stateRow.TradeNum)/ float64(stateTotal)),64) //交易平均星率

 err,rows := stateRow.Update()
 if err != nil {
  return 0,err
 }
 return rows,nil
}
