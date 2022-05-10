package logic

import (
	"Goi/dao/mysql"
	"Goi/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}
