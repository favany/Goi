package mysql

import (
	"Goi/models"
	"database/sql"
	"errors"
)

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

// 把每一步数据库操作封装成函数，待 logic 层根据业务需求调用

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库插入一条用户数据
func InsertUser(user *models.User) (err error) {
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

// Login 登陆
func Login(user *models.User) (err error) {
	oPassword := user.Password // 用户登陆的密码
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	} else if err != nil {
		return err
	}
	// 判断密码是否正确
	if oPassword != user.Password {
		return ErrorInvalidPassword
	}
	return
}

// GetUserById 根据用户id 获取用户名称
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
