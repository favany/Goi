package logic

import (
	"Goi/dao/mysql"
	"Goi/models"
	"Goi/pkg/jwt"
	"Goi/pkg/snowflake"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// 存放业务逻辑的代码

const secret = "Vooce.net"

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Sum([]byte(oPassword))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2. 生成UID
	userID := snowflake.GenID()
	// 构造一个 User 实例
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		// 密码加密🔐
		Password: encryptPassword(p.Password),
	}
	fmt.Println(user)
	// 3. 保存进数据库
	err = mysql.InsertUser(&user)
	return
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: encryptPassword(p.Password),
	}
	// 传递的是指针，就能拿到user.UserID
	if err = mysql.Login(user); err != nil {
		return "", err
	}
	// 生成JWT
	return jwt.GenToken(user.UserID)
}
