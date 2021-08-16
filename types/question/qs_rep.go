package question


type QsListRep struct {
	QsId       int64  `json:"qs_id"`
	QsAuthor   string `json:"qs_author"`
	QsTitle    string `json:"qs_title"`
	CreateTime string `json:"create_time"`
}


type QsDetailRep struct {
	QsId       int64   `json:"qs_id"`
	QsAuthor   string  `json:"qs_author"`
	QsTitle    string  `json:"qs_title"`
	QsDetail   string  `json:"qs_detail"`
	CreateTime string `json:"create_time"`
}
