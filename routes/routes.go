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

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("app.version"))
	})

	return r
}
