# Goi

A forum based on Go & Vue
一个基于 Go & Vue 的论坛 

Usage:

- Run Locally:
```shell
go run main.go
```
- Build and Deploy:
```shell
go build
```

Structure:
Web
 - conf: 配置文件
 - controllers: 请求参数的获取和校验
 - dao: 数据库的一些操作
 - logger: 日志
 - logic: 业务逻辑
 - models: 模板
 - pkg: 第三方的库
 - routes: 路由
 - settings: 配置相关函数
 - main.go: 主程序