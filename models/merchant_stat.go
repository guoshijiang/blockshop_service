package models


import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type MerchantStat struct {
	BaseModel
	Id             int64     `orm:"pk;column(id);auto;size(11)" description:"商品属性ID" json:"id"`
	MarchantId     int64     `orm:"column(marchant_id);size(64);index" description:"商户ID" json:"marchant_id"`
	SericeBest     int64     `orm:"column(serice_best);size(64);index" description:"服务好评" json:"serice_best"`
	ServiceGood    int64     `orm:"column(service_good);size(64);index" description:"服务中评" json:"service_good"`
	ServiceBad     int64     `orm:"column(service_good);size(64);index" description:"服务差评" json:"service_bad"`
	ServiceAvg	   float64   `orm:"column(service_avg);size(64);index" description:"服务平均星星率" json:"service_avg"`
	ServiceNum     int64     `orm:"column(service_num);size(64);index" description:"服务评价星星数量" json:"service_num"`
	TradeBest 	   int64     `orm:"column(serice_best);size(64);index" description:"交易好评" json:"trade_best"`
	TradeGood 	   int64     `orm:"column(trade_good);size(64);index" description:"交易中评" json:"trade_good"`
	TradeBad 	   int64     `orm:"column(trade_bad);size(64);index" description:"交易差评" json:"trade_bad"`
	TradeAvg       float64   `orm:"column(trade_avg);size(64);index" description:"交易平均星星率" json:"trade_avg"`
	TradeNum       int64     `orm:"column(trade_num);size(64);index" description:"交易评价星星数量" json:"trade_num"`
	QualityBest    int64     `orm:"column(quality_best);size(64);index" description:"交易差评" json:"quality_best"`
	QualityGood    int64     `orm:"column(quality_good);size(64);index" description:"交易差评" json:"quality_good"`
	QualityBad     int64     `orm:"column(quality_bad);size(64);index" description:"交易差评" json:"quality_bad"`
	QualityAvg     float64   `orm:"column(quality_avg);size(64);index" description:"质量平均星星率" json:"quality_avg"`
	QualityNum     int64     `orm:"column(quality_num);size(64);index" description:"质量评价星星数量" json:"quality_num"`
	TotalCmtNum    int64     `orm:"column(total_cmt_num);size(64);index" description:"总评价星星数量" json:"total_cmt_num"`
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

