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

    //上传图片
    beego.NSRouter("/setting/upload", &controllers.IndexController{}, "post:Upload"),

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

    //商品大类管理-商品大类管理
    beego.NSRouter("/goods/type/index", &controllers.GoodsTypeController{}, "get:Index"),
    //商品大类管理-添加界面
    beego.NSRouter("/goods/type/add", &controllers.GoodsTypeController{}, "get:Add"),
    //商品大类管理-添加
    beego.NSRouter("/goods/type/create", &controllers.GoodsTypeController{}, "post:Create"),
    //商品大类管理-修改界面
    beego.NSRouter("/goods/type/edit", &controllers.GoodsTypeController{}, "get:Edit"),
    //商品大类管理-修改
    beego.NSRouter("/goods/type/update", &controllers.GoodsTypeController{}, "post:Update"),
    //商品大类管理-删除
    beego.NSRouter("/goods/type/del", &controllers.GoodsTypeController{}, "post:Del"),

    //商品属性管理-商品属性管理
    beego.NSRouter("/goods/attr/index", &controllers.GoodsAttrController{}, "get:Index"),
    //商品属性管理-添加界面
    beego.NSRouter("/goods/attr/add", &controllers.GoodsAttrController{}, "get:Add"),
    //商品属性管理-添加
    beego.NSRouter("/goods/attr/create", &controllers.GoodsAttrController{}, "post:Create"),
    //商品属性管理-修改界面
    beego.NSRouter("/goods/attr/edit", &controllers.GoodsAttrController{}, "get:Edit"),
    //商品属性管理-修改
    beego.NSRouter("/goods/attr/update", &controllers.GoodsAttrController{}, "post:Update"),
    //商品属性管理-删除
    beego.NSRouter("/goods/attr/del", &controllers.GoodsAttrController{}, "post:Del"),

    //币种管理-商品属性管理
    beego.NSRouter("/asset/index", &controllers.AssetController{}, "get:Index"),
    //币种管理-添加界面
    beego.NSRouter("/asset/add", &controllers.AssetController{}, "get:Add"),
    //币种管理-添加
    beego.NSRouter("/asset/create", &controllers.AssetController{}, "post:Create"),
    //币种管理-修改界面
    beego.NSRouter("/asset/edit", &controllers.AssetController{}, "get:Edit"),
    //币种管理-修改
    beego.NSRouter("/asset/update", &controllers.AssetController{}, "post:Update"),
    //币种管理-删除
    beego.NSRouter("/asset/del", &controllers.AssetController{}, "post:Del"),

    //公告管理-商品属性管理
    beego.NSRouter("/news/index", &controllers.NewsController{}, "get:Index"),
    //公告管理-添加界面
    beego.NSRouter("/news/add", &controllers.NewsController{}, "get:Add"),
    //公告管理-添加
    beego.NSRouter("/news/create", &controllers.NewsController{}, "post:Create"),
    //公告管理-修改界面
    beego.NSRouter("/news/edit", &controllers.NewsController{}, "get:Edit"),
    //公告管理-修改
    beego.NSRouter("/news/update", &controllers.NewsController{}, "post:Update"),
    //公告管理-删除
    beego.NSRouter("/news/del", &controllers.NewsController{}, "post:Del"),


    //论坛分类管理-商品属性管理
    beego.NSRouter("/forum/category/index", &controllers.ForumCateController{}, "get:Index"),
    //论坛分类管理-添加界面
    beego.NSRouter("/forum/category/add", &controllers.ForumCateController{}, "get:Add"),
    //论坛分类管理-添加
    beego.NSRouter("/forum/category/create", &controllers.ForumCateController{}, "post:Create"),
    //论坛分类管理-修改界面
    beego.NSRouter("/forum/category/edit", &controllers.ForumCateController{}, "get:Edit"),
    //论坛分类管理-修改
    beego.NSRouter("/forum/category/update", &controllers.ForumCateController{}, "post:Update"),
    //论坛分类管理-删除
    beego.NSRouter("/forum/category/del", &controllers.ForumCateController{}, "post:Del"),
    //论坛管理-论坛列表
    beego.NSRouter("/forum/index", &controllers.ForumController{}, "get:List"),
    //论坛分类管理-添加界面
    beego.NSRouter("/forum/check", &controllers.ForumController{}, "post:CheckForum"),

    //论坛回复管理-论坛列表
    beego.NSRouter("/forum/reply/index", &controllers.ForumController{}, "get:Reply"),
    //论坛回复管理-审核
    beego.NSRouter("/forum/reply/check", &controllers.ForumController{}, "post:CheckReplyForum"),

    //区块链-用户钱包管理
    beego.NSRouter("/wallet/user/index", &controllers.WalletController{}, "get:User"),
    //区块链-钱包记录管理
    beego.NSRouter("/wallet/record/index", &controllers.WalletController{}, "get:Record"),

    //工单管理-工单列表
    beego.NSRouter("/message/index", &controllers.MessageController{}, "get:Index"),
    beego.NSRouter("/message/history", &controllers.MessageController{}, "get:History"),
    beego.NSRouter("/message/send", &controllers.MessageController{}, "post:Send"),

    //地域管理-商品属性管理
    beego.NSRouter("/origin/index", &controllers.OriginStateController{}, "get:Index"),
    //地域管理-添加界面
    beego.NSRouter("/origin/add", &controllers.OriginStateController{}, "get:Add"),
    //地域管理-添加
    beego.NSRouter("/origin/create", &controllers.OriginStateController{}, "post:Create"),
    //地域管理-修改界面
    beego.NSRouter("/origin/edit", &controllers.OriginStateController{}, "get:Edit"),
    //地域管理-修改
    beego.NSRouter("/origin/update", &controllers.OriginStateController{}, "post:Update"),
    //地域管理-删除
    beego.NSRouter("/origin/del", &controllers.OriginStateController{}, "post:Del"),

    //订单管理
    beego.NSRouter("/order/index",&controllers.OrderController{},"get:Index"),
    beego.NSRouter("/order/edit",&controllers.OrderController{},"get:Edit"),
    beego.NSRouter("/order/update",&controllers.OrderController{},"post:Update"),
    beego.NSRouter("/order/del",&controllers.OrderController{},"post:Del"),
    beego.NSRouter("/order/process",&controllers.OrderController{},"get:Process"),
    beego.NSRouter("/order/process/verify",&controllers.OrderController{},"post:Verify"),
    beego.NSRouter("/order/process/detail",&controllers.OrderController{},"get:Detail"),


    //用户管理
    beego.NSRouter("/user/index", &controllers.UserController{}, "get:Index"),
    //用户管理-添加界面
    beego.NSRouter("/user/add", &controllers.UserController{}, "get:Add"),
    //用户管理-账户资金
    //beego.NSRouter("/user/account", &controllers.UserController{}, "get:Account"),
    //用户管理-添加
    beego.NSRouter("/user/create", &controllers.UserController{}, "post:Create"),
    //用户管理-修改界面
    beego.NSRouter("/user/edit", &controllers.UserController{}, "get:Edit"),
    //用户管理-修改
    beego.NSRouter("/user/update", &controllers.UserController{}, "post:Update"),
    //用户管理-删除
    beego.NSRouter("/user/del", &controllers.UserController{}, "post:Del"),
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

    beego.NSNamespace("/goods",
      beego.NSInclude(
        &api.GoodsController{},
      ),
    ),

    beego.NSNamespace("/uwallet",
      beego.NSInclude(
        &api.UserWalletController{},
      ),
    ),

    beego.NSNamespace("/marchant",
      beego.NSInclude(
        &api.MerchantController{},
      ),
    ),

    beego.NSNamespace("/order",
      beego.NSInclude(
        &api.OrderController{},
      ),
    ),

    beego.NSNamespace("/user",
        beego.NSInclude(
            &api.UserController{},
        ),
    ),

    beego.NSNamespace("/forum",
      beego.NSInclude(
        &api.ForumController{},
      ),
    ),

    beego.NSNamespace("/message",
      beego.NSInclude(
        &api.MessageController{},
      ),
    ),

    beego.NSNamespace("/cwallet",
      beego.NSInclude(
        &api.ChainWalletController{},
      ),
    ),

    beego.NSNamespace("/comment",
      beego.NSInclude(
        &api.CommentController{},
      ),
    ),

    beego.NSNamespace("/address",
      beego.NSInclude(
        &api.UserAddressController{},
      ),
    ),

    beego.NSNamespace("/blacklist",
      beego.NSInclude(
        &api.BlackListController{},
      ),
    ),

    beego.NSNamespace("/goods_collect",
      beego.NSInclude(
        &api.GoodsCollectController{},
      ),
    ),

    beego.NSNamespace("/marchant_collect",
      beego.NSInclude(
        &api.MerchantCollectController{},
      ),
    ),

    beego.NSNamespace("/help_desk",
      beego.NSInclude(
        &api.HelpDeskController{},
      ),
    ),

    beego.NSNamespace("/question",
      beego.NSInclude(
        &api.QuestionController{},
      ),
    ),
  )
  beego.AddNamespace(api_path)
}
