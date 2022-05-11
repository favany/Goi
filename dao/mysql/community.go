package mysql

import (
	"Goi/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

//func GetCommunityDetailById(idStr int64) (community *models.Community, err error) {
//	community = new(models.Community)
//	sqlStr := `select community_id, community_name
//	from community
//	where community_id = ?`
//	err = db.Get(community, sqlStr, idStr)
//	if err == sql.ErrNoRows {
//		//err = ErrorInvalidID
//		return
//	}
//	if err != nil {
//		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
//		//err = ErrorQueryFailed
//		return
//	}
//	return
//}

func GetCommunityDetailById(idStr int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
	from community
	where community_id = ?`
	err = db.Get(community, sqlStr, idStr)
	if err == sql.ErrNoRows {
		//err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		//err = ErrorQueryFailed
		return
	}
	return
}
