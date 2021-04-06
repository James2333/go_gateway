<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [go_gateway](#go_gateway)
    - [现在开始](#%E7%8E%B0%E5%9C%A8%E5%BC%80%E5%A7%8B)
    - [文件分层](#%E6%96%87%E4%BB%B6%E5%88%86%E5%B1%82)
    - [log / redis / mysql / http.client 常用方法](#log--redis--mysql--httpclient-%E5%B8%B8%E7%94%A8%E6%96%B9%E6%B3%95)
    - [swagger文档生成](#swagger%E6%96%87%E6%A1%A3%E7%94%9F%E6%88%90)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# go_gateway
Gin best practices, gin development scaffolding, too late to explain, get on the bus.

使用gin构建了企业级脚手架，代码简洁易读，可快速进行高效web开发。
主要功能有：
1. 请求链路日志打印，涵盖mysql/redis/request
2. 支持多语言错误信息提示及自定义错误提示。
3. 支持了多配置环境
4. 封装了 log/redis/mysql/http.client 常用方法
5. 支持swagger文档生成

### 现在开始
- 安装软件依赖
go mod使用请查阅：

https://blog.csdn.net/e421083458/article/details/89762113
```
git clone https://github.com/James2333/go_gateway.git
cd go_gateway
go mod tidy
```
- 确保正确配置了 conf/mysql_map.toml、conf/redis_map.toml：

- 运行sql脚本  这边建议直接用PHPadmin快速启动数据库，然后导入sql文件即可

```
go run main.go

➜  go_gateway git:(master) ✗ go run main.go
------------------------------------------------------------------------
[INFO]  config=./conf/dev/
[INFO]  start loading resources.
[INFO]  success loading resources.
------------------------------------------------------------------------
....
 [INFO] HttpServerRun::8880
```

### 文件分层
```
.
├── conf   配置文件夹
│   └── dev
│       ├── base.toml
│       ├── mysql_map.toml
│       └── redis_map.toml
├── controller   控制层
│   ├── admin.go
│   ├── admin_login.go
│   └── service.go
├── dao    DB数据层
│   ├── admin.go
│   ├── service_access_control.go
│   ├── service.go
│   ├── service_grpc_rule.go
│   ├── service_http_rule.go
│   ├── service_info.go
│   ├── service_load_balance.go
│   └── service_tcp_rule.go
├── docs   swagger文件层
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── dto     输入输出结构层
│   ├── admin.go
│   ├── admin_login.go
│   ├── dashboard.go
│   └── service.go
├── go
├── go.mod
├── go.sum
├── logs    收集的日志文件
│   ├── gin_scaffold.inf.log
│   └── gin_scaffold.wf.log
├── main.go
├── middleware   中间件层
│   ├── ip_auth.go
│   ├── recovery.go
│   ├── request_log.go
│   ├── response.go
│   ├── session_auth.go
│   └── translation.go
├── public    公共文件
│   .....
├── README.md
├── reverse_proxy
│   └── load_balance   负载均衡算法
│        .....
└── router     路由层
    ├── httpserver.go
    └── route.go

```

### log / redis / mysql / http.client 常用方法

参考文档：https://github.com/e421083458/golang_common


### swagger文档生成

https://github.com/swaggo/swag/releases

- linux/mac 下载对应操作系统的执行文件到$GOPATH/bin下面

如下：
```
➜  go_gateway git:(master) ✗ ll -r $GOPATH/bin
total 434168
-rwxr-xr-x  1 niuyufu  staff    13M  4  3 17:38 swag
```
- windows下需要找到GOPATH下的github文件夹下的swaggo文件

先执行：
```
go get -u github.com/swaggo/swag/cmd/swag
```
然后找到$GOPATH/pkg\mod\github.com\swaggo\swag@*
这个路径下，cmd中执行 go install 即可安装swag成功
然后回到项目路径下
```
swag init && go run main.go 
```
生成接口文档并启动项目。
- 设置接口文档参考： `controller/demo.go` 的 Bind方法的注释设置

```
// ListPage godoc
// @Summary 测试数据绑定
// @Description 测试数据绑定
// @Tags 用户
// @ID /demo/bind
// @Accept  json
// @Produce  json
// @Param polygon body dto.DemoInput true "body"
// @Success 200 {object} middleware.Response{data=dto.DemoInput} "success"
// @Router /demo/bind [post]
```

- 生成接口文档：`swag init`
- 然后启动服务器：`go run main.go`，浏览地址: http://127.0.0.1:8880/swagger/index.html

