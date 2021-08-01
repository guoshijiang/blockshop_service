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

    //商户管理-商户管理
    beego.NSRouter("/merchant/index", &controllers.MerchantController{}, "get:Index"),
    //商户管理-添加界面
    beego.NSRouter("/merchant/add", &controllers.MerchantController{}, "get:Add"),
    //商户管理-添加
    beego.NSRouter("/merchant/create", &controllers.MerchantController{}, "post:Create"),
    //商户管理-修改界面
    beego.NSRouter("/merchant/edit", &controllers.MerchantController{}, "get:Edit"),
    //商户管理-修改
    beego.NSRouter("/merchant/update", &controllers.MerchantController{}, "post:Update"),
    //商户管理-删除
    beego.NSRouter("/merchant/del", &controllers.MerchantController{}, "post:Del"),

    //商品管理-商品管理
    beego.NSRouter("/goods/index", &controllers.GoodsController{}, "get:Index"),
    //商品管理-添加界面
    beego.NSRouter("/goods/add", &controllers.GoodsController{}, "get:Add"),
    //商品管理-添加
    beego.NSRouter("/goods/create", &controllers.GoodsController{}, "post:Create"),
    //商品管理-修改界面
    beego.NSRouter("/goods/edit", &controllers.GoodsController{}, "get:Edit"),
    //商品管理-修改
    beego.NSRouter("/goods/update", &controllers.GoodsController{}, "post:Update"),
    //商品管理-删除
    beego.NSRouter("/goods/del", &controllers.GoodsController{}, "post:Del"),

    //商品分类管理-商品分类管理
    beego.NSRouter("/goods/category/index", &controllers.GoodsCateController{}, "get:Index"),
    //商品分类管理-添加界面
    beego.NSRouter("/goods/category/add", &controllers.GoodsCateController{}, "get:Add"),
    //商品分类管理-添加
    beego.NSRouter("/goods/category/create", &controllers.GoodsCateController{}, "post:Create"),
    //商品分类管理-修改界面
    beego.NSRouter("/goods/category/edit", &controllers.GoodsCateController{}, "get:Edit"),
    //商品分类管理-修改
    beego.NSRouter("/goods/category/update", &controllers.GoodsCateController{}, "post:Update"),
    //商品分类管理-删除
    beego.NSRouter("/goods/category/del", &controllers.GoodsCateController{}, "post:Del"),

    //商品分类管理-商品属性管理
    beego.NSRouter("/goods/type/index", &controllers.GoodsTypeController{}, "get:Index"),
    //商品属性管理-添加界面
    beego.NSRouter("/goods/type/add", &controllers.GoodsTypeController{}, "get:Add"),
    //商品属性管理-添加
    beego.NSRouter("/goods/type/create", &controllers.GoodsTypeController{}, "post:Create"),
    //商品属性管理-修改界面
    beego.NSRouter("/goods/type/edit", &controllers.GoodsTypeController{}, "get:Edit"),
    //商品属性管理-修改
    beego.NSRouter("/goods/type/update", &controllers.GoodsTypeController{}, "post:Update"),
    //商品属性管理-删除
    beego.NSRouter("/goods/type/del", &controllers.GoodsTypeController{}, "post:Del"),



    )
  beego.AddNamespace(admin)

  api_path := beego.NewNamespace("/v1",
      beego.NSNamespace("/image",
        beego.NSInclude(
          &api.ImageController{},
        ),
      ),

      beego.NSNamespace("/news",
        beego.NSInclude(
          &api.NewsController{},
        ),
      ),

      beego.NSNamespace("/user",
          beego.NSInclude(
              &api.UserController{},
          ),
      ),
  )
  beego.AddNamespace(api_path)
}
