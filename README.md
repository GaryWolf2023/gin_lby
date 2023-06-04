# lby_go go语言学习
go学习

项目采用air工具作为热更新工具，使用方法如下：
1. 安装air工具
```
go get -u github.com/cosmtrek/air
```
2. 在项目根目录下执行air命令
```
air
```

controller: 控制器层，负责处理请求，返回响应
models: 模型层，负责数据模型定义，及相关操作
service: 服务层，负责业务逻辑处理
util: 工具层，负责各种工具类
tools: 存放一些自定义的工具方法
routers: 路由层，负责路由转发
conf: 配置层，负责配置文件的读取
middleware: 中间件层，负责中间件的实现
