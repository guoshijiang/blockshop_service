### 1.项目描述

本项目是一个基于 beego 开发的商城系统后台，包含管理台和 api 接口，前端的 Vue 代码也已经开源呆本账号下，如果需要，请自取。

项目后端登陆用户可以是 admin, 密码 asd..123 ：


### 2.项目开发部署

第一步：克隆代码

```bigquery
git clone git@github.com:guoshijiang/blockshop_service.git
cd blockshop_service
```

第二步：数据库 migrate

去 models 下面的 base.go 里面打开 `RunSyncdb` 函数，自动生成数据库，注意配置 false 和 true 项

第三步：运行开发
```bigquery
bee run 即可运行代码进行开发
```

第三步：代码部署

```bigquery
go build 完成之后，使用进程管理管理启动服务即可
```

如果您使用这套代码，开发搭建过程中有任何问题，可以去问我学院（www.wenwoha.com） 上面找联系方式联系我们，也可以直接加我的微信：LGZAXE


### 4.使用该代码做二开的条件

使用本套代码做二次开发，需要把所有的产品加友链。具体的友链请加微信聊

