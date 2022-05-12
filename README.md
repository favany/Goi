# Goi

Go WEB backend system
Go Web 后台系统

Feature 特色:
- 🚀 Swift Hot updates 疾速热更新
- 🗄 Clear Project Structure 项目结构清晰

Usage:
- Install dependencies locally:
```shell
go mod tidy
```

- Run Locally:
```shell
# hot update
air
```
or
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
 - air: 使用Air实现Go程序实时热重载