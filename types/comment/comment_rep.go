package comment


import "time"

type CommentListRep struct {
	CommentId   int64  `json:"comment_id"`
	UserName    string `json:"user_name"`
	UserPho     string `json:"user_pho"`
	GoodsId     int64  `json:"goods_id"`
	UserId      int64  `json:"user_id"`
	Title       string `json:"title"`
  QualityStar  int8   `json:"quality_star"`
  ServiceStar  int8   `json:"Service_star"`
  TradeStar    int8   `json:"trade_star"`
	Content     string `json:"content"`
	ImgOne      string `json:"img_one_id"`
	ImgTwo      string `json:"img_two_id"`
	ImgThree    string `json:"img_three_id"`
	CreateTime  time.Time `json:"create_time"`
}