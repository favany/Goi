package routes

import (
	"Goi/controller"
	"Goi/logger"
	"Goi/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// 连通测试
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	// 登陆业务路由
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("app.version"))
	})

	{
		v1.GET("/community", controller.CommunityHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/get_post_list/", controller.GetPostListHandler)
	}

	return r
}
