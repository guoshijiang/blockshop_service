package types

import "github.com/pkg/errors"

// 错误码定义
const (
	ReturnSuccess                 = 2000  // 成功返回
	SystemDbErr                   = 3000  // 数据库错误
	InvalidFormatError            = 3001  // 无效的参数格式
	ParamEmptyError               = 3002  // 传入参数为空
	UserToKenCheckError           = 3003  // 用户 Token 校验失败
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
)

type PageSizeData struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (this PageSizeData) PageSizeDataParamValidate() (int, error) {
	if this.Page == 0 {
		return PageIsZero, errors.New("page 不能为 0")
	}
	if this.PageSize == 0 {
		return PageSizeIsZero, errors.New("pageSize 不能为 0")
	}
	return ReturnSuccess, nil
}