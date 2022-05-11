package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUserID 获取当前登陆的用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getPageInfo(c *gin.Context) (pageNum, pageSize int64) {
	// 获取分页参数
	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")

	var err error

	pageNum, err = strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		pageNum = 1
	}
	pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 10
	}
	return
}
