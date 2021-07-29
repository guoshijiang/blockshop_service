//自定义模板函数
package template

import (
	"github.com/astaxie/beego"
	"time"
)

func init() {
	beego.AddFuncMap("TimeForFormat", TimeForFormat)
	beego.AddFuncMap("DateForFormat", DateForFormat)
	beego.AddFuncMap("WalletRecordType", WalletRecordType)
	beego.AddFuncMap("WalletRecordIsHandle", WalletRecordIsHandle)
	beego.AddFuncMap("WalletRecordSource", WalletRecordSource)
	beego.AddFuncMap("WalletRecordStatus", WalletRecordStatus)
	beego.AddFuncMap("ProcessStatus", ProcessStatus)
	beego.AddFuncMap("ProcessIsRecvGoods", ProcessIsRecvGoods)
	beego.AddFuncMap("ProcessFundRet", ProcessFundRet)
	beego.AddFuncMap("UnixTimeForFormat", UnixTimeForFormat)
	beego.AddFuncMap("OrderStatus", OrderStatus)
	beego.AddFuncMap("CancelStatus", CancelStatus)
	beego.AddFuncMap("PayWay", PayWay)
	beego.AddFuncMap("IntegralType", IntegralType)
	beego.AddFuncMap("IntegralRecord", IntegralRecord)
	beego.AddFuncMap("SettleStatus", SettleStatus)
}

//时间轴转时间字符串
func UnixTimeForFormat(timeUnix int) string {
	//转化所需模板
	timeLayout := "2006-01-02 15:04:05"
	return time.Unix(int64(timeUnix), 0).Format(timeLayout)
}

func TimeForFormat(t interface{}) string {
	timeLayout := "2006-01-02 15:04:05"
	if _,ok := t.(time.Time);ok {
		return t.(time.Time).Format(timeLayout)
	}

	if _,ok := t.(*time.Time);ok {
		if t.(*time.Time) == nil{
			return ""
		}
		return t.(*time.Time).Format(timeLayout)
	}
	return  ""
}

func DateForFormat(t interface{}) string {
	timeLayout := "2006-01-02"
	if _,ok := t.(time.Time);ok {
		return t.(time.Time).Format(timeLayout)
	}

	if _,ok := t.(*time.Time);ok {
		if t.(*time.Time) == nil{
			return ""
		}
		return t.(*time.Time).Format(timeLayout)
	}
	return  ""
}

//状态  0:交易中；1: 交易成功 2: 交易失败
func IntegralRecord(t int8) string {
	switch t {
	case 0:
		return "交易中"
	case 1:
		return "交易成功"
	case 2:
		return "交易失败"
	default:
		return "未知"
	}
}

//积分类型 // 1:邀请积分; 2:购买积分; 3: 管理奖励;  4:积分消费
func IntegralType(t int8) string {
	switch t {
	case 1:
		return "邀请积分"
	case 2:
		return "购买积分"
	case 3:
		return "管理奖励"
	case 4:
		return "积分消费"
	default:
		return "未知"
	}
}

//支付方式 0:积分兑换，1:账户余额支付，2:微信支付；3:支付宝支付; 4:未知支付方式
func PayWay(t int8) string {
	switch t {
	case 0:
		return "积分兑换"
	case 1:
		return "账户余额支付"
	case 2:
		return "微信支付"
	case 3:
		return "支付宝支付"
	default:
		return "未知"
	}
}

//订单状态 0: 未支付，1: 支付中，2：支付成功；3：支付失败 4：已发货；5：已完成
func OrderStatus(t int8) string {
	switch t {
	case 0:
		return "未支付"
	case 1:
		return "支付中"
	case 2:
		return "支付成功"
	case 3:
		return "支付失败"
	case 4:
		return "已发货"
	case 5:
		return "已完成"
	default:
		return "未知"
	}
}

//取消订单状态 0 正常；1.退货; 2:换货; 3:退货成功; 4:换货成功
func CancelStatus(t int8) string {
	switch t {
	case 0:
		return "正常"
	case 1:
		return "退货"
	case 2:
		return "换货"
	case 3:
		return "退货成功"
	case 4:
		return "换货成功"
	default:
		return "未知"
	}
}

//资金类型
func WalletRecordType(t int8) string{
	switch t {
	case 0:
		return "充值"
	case 1:
		return "提现"
	case 2:
		return "积分兑换"
	case 3:
		return "消费"
	default:
		return "未知"
	}
}

//资金处理状态
func WalletRecordIsHandle(t int8) string{
	switch t {
	case 0:
		return "审核中"
	case 1:
		return "审核通过"
	case 2:
		return "已打款"
	case 3:
		return "审核拒绝"
	default:
		return "未知"
	}
}

//来源平台 0：支付宝 1:微信; 2:积分兑换
func WalletRecordSource(t int8) string{
	switch t {
	case 0:
		return "支付宝"
	case 1:
		return "微信"
	case 2:
		return "积分兑换"
	default:
		return "未知"
	}
}

//来源平台 0:入账中; 1:入账成功; 2:入账失败
func WalletRecordStatus(t int8) string{
	switch t {
	case 0:
		return "入账中"
	case 1:
		return "入账成功"
	case 2:
		return "入账失败"
	default:
		return "未知"
	}
}

//退款订单状态
func ProcessStatus(t int8) string {
	switch t {
	case 0:
		return "等待卖家确认"
	case 1:
		return "卖家已同意"
	case 2:
		return "卖家拒绝"
	case 3:
		return "等待买家邮寄"
	case 4:
		return "等待卖家收货"
	case 5:
		return "卖家已经发货"
	case 6:
		return "等待买家收货"
	case 7:
		return "已完成"
	default:
		return "未知"
	}
}

func ProcessIsRecvGoods(t int8) string {
	switch t {
	case 0:
		return "未收到货物"
	case 1:
		return "已收到货物"
	default:
		return "未知"
	}
}

func ProcessFundRet(t int8) string {
	switch t {
	case 0:
		return "返回到平台钱包"
	case 1:
		return "原路返回"
	default:
		return "未知"
	}
}

// 0:商家已确认； 1:平台已确认； 2：已付款
func SettleStatus(t int8) string{
	switch t {
	case 0:
		return "商家已确认"
	case 1:
		return "平台已确认"
	case 2:
		return "已付款"
	default:
		return "未知"
	}
}
