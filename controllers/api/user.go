package api

import (
	"blockshop/common"
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/user"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
)

type UserController struct {
	beego.Controller
}


// Register @Title Register
// @Description 注册手机号 Register
// @Success 200 status bool, data interface{}, msg string
// @router /register [post]
func (this *UserController) Register() {
	user_register := user.Register{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user_register); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := user_register.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	code, msg := models.UserRegister(user_register)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, code, nil, msg)
	} else {
		this.Data["json"] = RetResource(false, code, nil, msg)
	}
	this.ServeJSON()
	return
}

// Login @Title Login
// @Description 登录 Login
// @Success 200 status bool, data interface{}, msg string
// @router /login [post]
func (this *UserController) Login() {
	user_login := user.Login{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user_login); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := user_login.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	data, code, msg := models.UserLogin(user_login)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, code, data, msg)
	} else {
		this.Data["json"] = RetResource(false, code, nil, msg)
	}
	this.ServeJSON()
	return
}

// Get2Fa @Title Get2Fa
// @Description 获取双因子字符 Get2Fa
// @Success 200 status bool, data interface{}, msg string
// @router /get_2fa [post]
func (this *UserController) Get2Fa() {
	fa2 := user.TwoFaReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &fa2); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	verify_code := common.GenValidateCode(6)
	user_info := models.GetUserInfo(fa2.UserName)
	if user_info == nil {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, nil, "用户不存在")
		this.ServeJSON()
		return
	}
	if user_info.IsOpen == 0 {
		this.Data["json"] = RetResource(false, types.NoOpenTwoFactor, nil, "没有开启双因子认证")
		this.ServeJSON()
		return
	}
	public_key := user_info.UserPublicKey
	cipher_text := common.RsaEncrypt([]byte(verify_code), []byte(public_key))
	code, msg := models.UpdateFactor(user_info.Id, verify_code)
	if code == types.ReturnSuccess {
		data := user.TwoFactorRep{
			Id: user_info.Id,
			UserName: user_info.UserName,
			CipherText: hex.EncodeToString(cipher_text),
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取登录因子成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, code, nil, msg)
		this.ServeJSON()
		return
	}
}

// OpenClose2Fa @Title OpenClose2Fa
// @Description 开启双因子验证 OpenClose2Fa
// @Success 200 status bool, data interface{}, msg string
// @router /open_close_2fa [post]
func (this *UserController) OpenClose2Fa() {
	octf := user.OpenCloseTwoFa{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &octf); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	code, msg := models.OpenCloseFactor(octf.IsOpen, octf.UserId, octf.PublicKey)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "开通或者关闭双因子认证成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, code, nil, msg)
		this.ServeJSON()
		return
	}
}

// UpdatePassword @Title UpdatePassword
// @Description 修改密码 UpdatePassword
// @Success 200 status bool, data interface{}, msg string
// @router /update_password [post]
func (this *UserController)UpdatePassword() {
	upd_pwd := user.UpdatePasswordReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &upd_pwd); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	code, msg := models.UpdatePassword(upd_pwd)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "修改密码成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, code, nil, msg)
		this.ServeJSON()
		return
	}
}

// UpdatePinCode @Title UpdatePinCode
// @Description 修改Pin码接口 UpdatePinCode
// @Success 200 status bool, data interface{}, msg string
// @router /update_pin_code [post]
func (this *UserController)UpdatePinCode() {
	upd_pin_code := user.UpdatePinCodeReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &upd_pin_code); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	code, msg := models.UpdatePinCode(upd_pin_code)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "修改Pin码成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, code, nil, msg)
		this.ServeJSON()
		return
	}
}

// ForgetPassword @Title ForgetPassword
// @Description 忘记密码 ForgetPassword
// @Success 200 status bool, data interface{}, msg string
// @router /forget_pwd [post]
func (this *UserController) ForgetPassword() {
	forget_pwd := user.ForgetPasswordReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &forget_pwd); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	code, msg := models.ForgetPassword(forget_pwd)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "找回密码成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, code, nil, msg)
		this.ServeJSON()
		return
	}
}

// UpdateUserInfo @Title UpdateUserInfo
// @Description 修改用户信息 UpdateUserInfo
// @Success 200 status bool, data interface{}, msg string
// @router /update_user_info [post]
func (this *UserController) UpdateUserInfo() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	user_info := user.UpdateUserInfoReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user_info); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := user_info.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	code, msg := models.UpdateUser(user_info)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "修改用户信息成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, code, nil, msg)
		this.ServeJSON()
		return
	}
}

// GetUserInfo @Title GetUserInfo
// @Description 获取用户信息 GetUserInfo
// @Success 200 status bool, data interface{}, msg string
// @router /get_user_info [post]
func (this *UserController) GetUserInfo() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user_if, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var btc_c_price user.CoinPrice
	btc_price := models.GetAssetByName("BTC")
	if btc_price != nil {
		btc_c_price = user.CoinPrice {
			Asset: "BTC",
			ChainName: "Bitcoin",
			UsdPrice: btc_price.UsdPrice,
			CnyPrice: btc_price.CnyPrice,
		}
	} else {
		btc_c_price = user.CoinPrice {
			Asset: "BTC",
			ChainName: "Bitcoin",
			UsdPrice: "0",
			CnyPrice: "0",
		}
	}
	var usdt_price user.CoinPrice
	usdt_price_md := models.GetAssetByName("USDT")
	if usdt_price_md != nil {
		usdt_price = user.CoinPrice {
			Asset: "USDT",
			ChainName: "Trc20",
			UsdPrice: usdt_price_md.UsdPrice,
			CnyPrice: usdt_price_md.CnyPrice,
		}
	} else {
		usdt_price = user.CoinPrice {
			Asset: "USDT",
			ChainName: "Trc20",
			UsdPrice: "0",
			CnyPrice: "0",
		}
	}
	u_w_btc, _ := models.GetUserWalletByUserId(user_if.Id, btc_price.Id)
	var btc_uws user.UserWalletStat
	if u_w_btc != nil {
		btc_uws = user.UserWalletStat {
			OutAmount: u_w_btc.OutAmount,
			InAmount: u_w_btc.InAmount,
			Balance: u_w_btc.Balance,
			Address: u_w_btc.Address,
		}
	}
	u_w_usdt, _ := models.GetUserWalletByUserId(user_if.Id, usdt_price_md.Id)
	var usdt_uws user.UserWalletStat
	if u_w_usdt != nil {
		usdt_uws = user.UserWalletStat {
			OutAmount: u_w_usdt.OutAmount,
			InAmount: u_w_usdt.InAmount,
			Balance: u_w_usdt.Balance,
			Address: u_w_usdt.Address,
		}
	}
	// 判断用户是否开通商家
	var merchant_id int64
	mct, _, _ := models.GetMerchantByUserId(user_if.Id)
	if mct != nil {
		merchant_id = mct.Id
	} else {
		merchant_id = 0
	}
	var IsOpen2Fa, IsSetKey bool
	act_pct := "65%"
	if user_if.IsOpen == 0 {
		IsOpen2Fa =  false
		IsSetKey = false
		act_pct = "99%"
	} else {
		IsOpen2Fa = true
		IsSetKey = true
		act_pct = "99%"
	}
	usy := user.UserSecrity {
		AccountPct: act_pct,
		IsSetKey: IsSetKey,
		IsOpen2Fa: IsOpen2Fa,
	}
	t_btc, t_usdt, t_buy :=models.UserOderStat(user_if.Id)
	data := user.UserInfoRep{
		UserId: user_if.Id,
		MctId: merchant_id,
		Photo: user_if.Avator,
		UserName: user_if.UserName,
		IsMerchant: user_if.IsMerchant,
		JoinTime: user_if.CreatedAt.Format("2006-01-02 15:04:05"),
		TrustLevel: user_if.MemberLevel,
		PublicKey: user_if.UserPublicKey,
		BtcOrderAmount: t_btc,
		UsdtOrderAmount: t_usdt,
		TotalBuy: t_buy,
		AdjustVictor: user_if.AdjustVictor,
		AdjustFail: user_if.AdjustFail,
		UserSecrity: usy,
		BtcWtStat: btc_uws,
		UsdtWtStat: usdt_uws,
		BtcPrice: btc_c_price,
		UsdtPrice: usdt_price,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取用户信息成功")
	this.ServeJSON()
	return
}
