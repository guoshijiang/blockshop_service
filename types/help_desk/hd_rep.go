package help_desk


type HdListRep struct {
	Id         int64  `json:"qs_id"`
	HdTitle    string `json:"qs_title"`
	IsHand     int8   `json:"is_hand"`  // 0:未处理；1:已处理
	HdTime     string `json:"hd_time"`
}


type HdDetailRep struct {
	Id         int64  `json:"qs_id"`
	IsHand     int8   `json:"is_hand"`  // 0:未处理；1:已处理
	HdTitle    string `json:"qs_title"`
	HdDetai    string `json:"hd_detai"`
	HdTime     string `json:"hd_time"`
}
