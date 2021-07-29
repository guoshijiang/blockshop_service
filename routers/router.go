package routers

import (
  controllers "blockshop/controllers/admin"
  "blockshop/controllers/api"
  "blockshop/middleware"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/context"
  "github.com/dchest/captcha"
  "net/http"
)

func init() {
  middleware.AuthMiddle()
  //admin
  beego.Get("/", func(ctx *context.Context) {
    ctx.Redirect(http.StatusFound, "/admin/index/index")
  })

  admin := beego.NewNamespace("/admin",
      beego.NSRouter("/index/index",&controllers.IndexController{}, "get:Index"),
    beego.NSRouter("/admin_log/index", &controllers.AdminLogController{}, "get:Index"),
    //登录页
    beego.NSRouter("/auth/login", &controllers.AuthController{}, "get:Login"),
    //退出登录
    beego.NSRouter("/auth/logout", &controllers.AuthController{}, "get:Logout"),
    //二维码图片输出
    beego.NSHandler("/auth/captcha/*.png", captcha.Server(240, 80)),
    //登录认证
    beego.NSRouter("/auth/check_login", &controllers.AuthController{}, "post:CheckLogin"),
    //刷新验证码
    beego.NSRouter("/auth/refresh_captcha", &controllers.AuthController{}, "post:RefreshCaptcha"),

    //首页
    beego.NSRouter("/index/index", &controllers.IndexController{}, "get:Index"),

    beego.NSRouter("/admin_user/index", &controllers.AdminUserController{}, "get:Index"),

    //菜单管理
    beego.NSRouter("/admin_menu/index", &controllers.AdminMenuController{}, "get:Index"),
    //菜单管理-添加菜单-界面
    beego.NSRouter("/admin_menu/add", &controllers.AdminMenuController{}, "get:Add"),
    //菜单管理-添加菜单-创建
    beego.NSRouter("/admin_menu/create", &controllers.AdminMenuController{}, "post:Create"),
    //菜单管理-修改菜单-界面
    beego.NSRouter("/admin_menu/edit", &controllers.AdminMenuController{}, "get:Edit"),
    //菜单管理-更新菜单
    beego.NSRouter("/admin_menu/update", &controllers.AdminMenuController{}, "post:Update"),
    //菜单管理-删除菜单
    beego.NSRouter("/admin_menu/del", &controllers.AdminMenuController{}, "post:Del"),

    //系统管理-个人资料
    beego.NSRouter("/admin_user/profile", &controllers.AdminUserController{}, "get:Profile"),
    //系统管理-个人资料-修改昵称
    beego.NSRouter("/admin_user/update_nickname", &controllers.AdminUserController{}, "post:UpdateNickName"),
    //系统管理-个人资料-修改密码
    beego.NSRouter("/admin_user/update_password", &controllers.AdminUserController{}, "post:UpdatePassword"),
    //系统管理-个人资料-修改头像
    beego.NSRouter("/admin_user/update_avatar", &controllers.AdminUserController{}, "post:UpdateAvatar"),
    //系统管理-用户管理-添加界面
    beego.NSRouter("/admin_user/add", &controllers.AdminUserController{}, "get:Add"),
    //系统管理-用户管理-添加
    beego.NSRouter("/admin_user/create", &controllers.AdminUserController{}, "post:Create"),
    //系统管理-用户管理-修改界面
    beego.NSRouter("/admin_user/edit", &controllers.AdminUserController{}, "get:Edit"),
    //系统管理-用户管理-修改
    beego.NSRouter("/admin_user/update", &controllers.AdminUserController{}, "post:Update"),
    //系统管理-用户管理-启用
    beego.NSRouter("/admin_user/enable", &controllers.AdminUserController{}, "post:Enable"),
    //系统管理-用户管理-禁用
    beego.NSRouter("/admin_user/disable", &controllers.AdminUserController{}, "post:Disable"),
    //系统管理-用户管理-删除
    beego.NSRouter("/admin_user/del", &controllers.AdminUserController{}, "post:Del"),
    //系统管理-角色管理
    beego.NSRouter("/admin_role/index", &controllers.AdminRoleController{}, "get:Index"),
    //系统管理-角色管理-添加界面
    beego.NSRouter("/admin_role/add", &controllers.AdminRoleController{}, "get:Add"),
    //系统管理-角色管理-添加
    beego.NSRouter("/admin_role/create", &controllers.AdminRoleController{}, "post:Create"),
    //菜单管理-角色管理-修改界面
    beego.NSRouter("/admin_role/edit", &controllers.AdminRoleController{}, "get:Edit"),
    //菜单管理-角色管理-修改
    beego.NSRouter("/admin_role/update", &controllers.AdminRoleController{}, "post:Update"),
    //菜单管理-角色管理-删除
    beego.NSRouter("/admin_role/del", &controllers.AdminRoleController{}, "post:Del"),
    //菜单管理-角色管理-启用角色
    beego.NSRouter("/admin_role/enable", &controllers.AdminRoleController{}, "post:Enable"),
    //菜单管理-角色管理-禁用角色
    beego.NSRouter("/admin_role/disable", &controllers.AdminRoleController{}, "post:Disable"),
    //菜单管理-角色管理-角色授权界面
    beego.NSRouter("/admin_role/access", &controllers.AdminRoleController{}, "get:Access"),
    //菜单管理-角色管理-角色授权
    beego.NSRouter("/admin_role/access_operate", &controllers.AdminRoleController{}, "post:AccessOperate"),
    )
  beego.AddNamespace(admin)
  //api
	api_path := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&api.UserController{},
			),
		),
	)
	beego.AddNamespace(api_path)
}
