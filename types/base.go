package types

import "github.com/pkg/errors"

// 错误码定义
const (
	ReturnSuccess                 = 2000  // 成功返回
	SystemDbErr                   = 3000  // 数据库错误
	InvalidFormatError            = 3001  // 无效的参数格式
	ParamEmptyError               = 3002  // 传入参数为空
	UserToKenCheckError           = 3003  // 用户 Token 校验失败
	ParamLessZero                 = 3004  // 参数小于 0
	PageIsZero                    = 4000  // 页码 0
	PageSizeIsZero                = 4001  // 每页数量 0
	PasswordNotEqual              = 4002  // 两次输入的密码不一样
	UserExist                     = 4003  // 用户已经存在
	UserNoExist                   = 4004  // 没有这个用户
	GetImagesFileFail             = 4005  // 获取文件失败
	FileFormatError               = 4006  // 文件格式不符合规范
	FileIsBig                     = 4007  // 文件太大了
	CreateFilePathError           = 4008  // 创建文件路径失败
	FileAlreadUpload              = 4009  // 该图片已经上传过了
	QueryNewsFail                 = 4010  // 查询新闻失败
	NoOpenTwoFactor               = 4011  // 没有开启双因子验证
	FactorIsNotSame               = 4012  // 双因子不正确
	PasswordError                 = 4013  // 两次输入的密码不一样
	GetGoodsListFail              = 4014  // 获取商品列表失败
	GetMerchantListFail           = 4015  // 获取商家列表失败
	InvalidVerifyWay              = 4016  // 无效的付款方式
	InvalidGoodsPirce             = 4017  // 无效的商品价格
	UserIsNotExist                = 4018  // 用户不存在
	AlreadyCancleOrder            = 4019  // 订单已经取消
	AddressIsEmpty                = 4020  // 地址为空
	AssetNameIsEmpty              = 4021  // 资产名称为空
	ChainNameIsEmpty              = 4022  // 链的名称为空
	NothisWallet                  = 4023  // 没有这个钱包
	UserIdEmptyError              = 4024  // 用户ID为空
	CreateWalletFail              = 4025  // 创建钱包失败
	OrderPayException             = 4026  // 订单支付异常
	OrderAlreadyPay               = 4027  // 订单已经支付
	VerifyPayAmount               = 4028  // 验证支付金额失败
	AccountAmountNotEnough        = 4029  // 账户余额不足
	PayOrderError                 = 4030  // 订单支付错误
	QueryMessageFail              = 4031  // 查询消息失败
	MessageEmpty                  = 4032  // 消息是空的
	GetAssetListError             = 4033  // 获取资产失败
	WithdrawToAddressEmpty        = 4034  // 提现转入地址为空
	TxFeeNotEnough                = 4035  // 提币手续费太低
	WithdrawAmountLessMin         = 4036  // 提币数量太小
	WalletBalanceNotEnough        = 4037  // 钱包余额不够
	QueryWalletRecodFail          = 4038  // 获取钱包记录失败
	CreateGoodsFail               = 4039  // 添加商品失败
	GetBtcRateFail                = 4040  // 获取 BTC 的费率失败
	GetUsdtRateFail               = 4041  // 获取 USDT 的费率失败
	PhoneEmptyError               = 4042  // 手机号码不能为空
	PhoneFormatError              = 4043  // 手机号码格式错误
	AddressIdLessEqError          = 4044  // 地址不存在
	CreateAddressFail             = 4045  // 创建地址失败
	UpdateAddressFail             = 4046  // 修改地址失败
	StaticDataFail                = 4047  // 统计失败
	OpenMerchantFail              = 4048  // 开通商家失败
	BlackListExist                = 4049  // 已经加入过黑名单


	InvalidWOrD                   = 5000  // 无效的充值或者提现
	AmountLessZero                = 5001  // 充值或提现小于 0
	TxHashEmpty                   = 5002  // 交易 Hash 为空
	TxFeeLessZero                 = 5003  // 交易手续小于 0
)

type PageSizeData struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (this PageSizeData) SizeParamCheck() (int, error) {
	if this.Page == 0 {
		return PageIsZero, errors.New("page 不能为 0")
	}
	if this.PageSize == 0 {
		return PageSizeIsZero, errors.New("pageSize 不能为 0")
	}
	return ReturnSuccess, nil
}