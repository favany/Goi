package mysql

import (
	"Goi/models"
	"database/sql"
	"errors"
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
		return errors.New("用户已存在")
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
		return errors.New("用户不存在")
	} else if err != nil {
		return err
	}
	// 判断密码是否正确
	if oPassword != user.Password {
		return errors.New("密码错误")
	}
	return
}
