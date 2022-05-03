package controller

import (
	"Goi/logic"
	"Goi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// 参数校验

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	err := c.ShouldBindJSON(&p)
	// 手动对请求参数进行详细的业务规则校验
	if err != nil || len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		// 请求参数有误，直接返回响应
		zap.L().Error("Sign up with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	fmt.Println(p)

	// 2. 业务处理
	logic.SignUp(p)
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
