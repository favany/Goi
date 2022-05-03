package logic

import (
	"Goi/dao/mysql"
	"Goi/models"
	"Goi/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp(m *models.ParamSignUp) {
	// 判断用户存不存在
	mysql.QueryUserByUsername()
	// 生成UID
	snowflake.GenID()
	// 密码加密🔐

	// 保存进数据库
	mysql.InsertUser()
}
