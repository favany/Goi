package controller

import (
	"Goi/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id, community_name）以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不要轻易把服务器报错暴露在外面
		return
	}
	ResponseSuccess(c, data)
}
