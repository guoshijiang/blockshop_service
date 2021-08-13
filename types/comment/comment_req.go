package comment

import (
	"blockshop/types"
	"github.com/pkg/errors"
)

type AddCommentReq struct {
	OrderId      int64  `json:"order_id"`
	GoodsId      int64  `json:"goods_id"`
	UserId       int64  `json:"user_id"`
  MerchantId   int64  `json:"merchant_id"`
	Title        string `json:"title"`
	QualityStar  int8   `json:"quality_star"`
	ServiceStar  int8   `json:"Service_star"`
	TradeStar    int8   `json:"trade_star"`
	Content      string `json:"content"`
	ImgOneId     int64  `json:"img_one_id"`
	ImgTwoId     int64  `json:"img_two_id"`
	ImgThreeId   int64  `json:"img_three_id"`

}

func (this AddCommentReq) ParamCheck() (int, error) {
	if this.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("商品ID不能小于0")
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID不能小于0")
	}
	if this.MerchantId <= 0 {
    return types.ParamLessZero, errors.New("商户ID不能小于0")
  }
	if this.Title == "" {
		return types.ParamEmptyError, errors.New("评论标题为空")
	}
	if this.ServiceStar <= 0 || this.TradeStar <= 0  || this.QualityStar <= 0 {
		return types.ParamEmptyError, errors.New("评论星级不能小于0")
	}
	if this.Content == "" {
		return types.ParamEmptyError, errors.New("评论内容为空")
	}
	if this.ImgOneId < 0 {
		return types.ParamEmptyError, errors.New("评论图片一ID小于0")
	}
	if this.ImgTwoId < 0 {
		return types.ParamEmptyError, errors.New("评论图片二ID小于0")
	}
	if this.ImgThreeId < 0 {
		return types.ParamEmptyError, errors.New("评论图片三ID小于0")
	}
	return types.ReturnSuccess, nil
}

type DelCommentReq struct {
	CommentId  int64 `json:"comment_id"`
	UserId     int64 `json:"user_id"`
}

func (this DelCommentReq) ParamCheck() (int, error) {
	if this.CommentId <= 0 || this.UserId <= 0 {
		return types.ParamLessZero, errors.New("评论ID和用户ID不能小于0")
	}
	return types.ReturnSuccess, nil
}

type CommentListReq struct {
	types.PageSizeData
	GoodsId  int64 `json:"goods_id"`
}

func (this CommentListReq) ParamCheck() (int, error) {
	code, err := this.ParamCheck()
	if err != nil {
		return code, err
	}
	if this.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("评论ID和用户ID不能小于0")
	}
	return types.ReturnSuccess, nil
}

