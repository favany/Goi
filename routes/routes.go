package routes

import (
	"Goi/controller"
	"Goi/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 连通测试
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	// 登陆业务路由
	r.POST("/login", controller.LoginHandler)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("app.version"))
	})

	r.GET("/ping", func(c *gin.Context) {
		// 如果是登陆的用户
		if isLogin {
			c.String(http.StatusOK, "pong")
		} else {
			// 否则就直接返回请登录
			c.String(http.StatusOK, "请登录")
		}
	})

	return r
}
