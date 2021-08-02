package form_validate


type UserForm struct {
  Id             int64         `form:"id"`
  UserName       string        `form:"user_name"`
  Avator         string        `form:"avator"`
  Password       string        `form:"password"`
  PinCode        string        `form:"pin_code"`
  FundPassword   string        `form:"fund_password"`
  LoginCount     int64         `form:"login_count"`
  Token          string        `form:"token"`
  IsMerchant     int8          `form:"is_merchant"` // ：0 不是，1: 是
  MemberLevel    int8          `form:"member_level"`  // 0:v0:普通会员 1:V1:白银会员，2:V2:白金会员，3:V3:黄金会员; 4:V4:砖石会有; 5:V5:皇冠会员
  Factor         string        `form:"factor"`
  IsOpen         int8          `form:"is_open"` // ：0 不是，1: 是
  TimeRange	     string        `form:"time_range"`
  IsCreate 	     int    		    `form:"_create"`
}