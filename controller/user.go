package controller

import (
	"Goi/logic"
	"Goi/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	if err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Sign up with invalid param", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
				"msg":    err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
				"msg":    removeTopStruct(errs.Translate(trans)), // 翻译错误
			})
			return
		}
		return
	}

	// 2. 业务逻辑处理
	err = logic.SignUp(p)

	// 3. 返回响应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    "注册失败",
			"info":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"info":   "注册成功！",
		})
	}

}

func LoginHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamLogin)
	err := c.ShouldBindJSON(&p)

	// 手动对请求参数进行详细的业务规则校验
	if err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Sign up with invalid param", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
				"msg":    err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
				"msg":    removeTopStruct(errs.Translate(trans)), // 翻译错误
			})
			return
		}
		return
	}
	// 2. 业务逻辑处理
	err = logic.Login(p)

	// 3. 返回响应
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    "用户名或密码错误",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"msg":    "登陆成功",
		})
	}

}
