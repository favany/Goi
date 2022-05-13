package models

// 定义请求的参数结构体

// ParamSignUp 注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登陆请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamPostList 获取帖子列表 query string 参数
type ParamPostList struct {
	Page  int64  `json:"page"`
	Size  int64  `json:"size"`
	Order string `json:"order"`
}

const (
	OrderTime  = "time"
	OrderScore = "score"
)
