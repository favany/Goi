package controller

import (
	"Goi/logic"
	"Goi/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子的处理函数
func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	// 1. 获取参数及参数的校验
	// ShouldBindJSON validator --> binding tag
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数（从 url 中获取帖子的 id）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 根据 ID获取帖子的数据（查数据库）
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {
	pageNum, pageSize := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
	return
}
