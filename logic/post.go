package logic

import (
	"Goi/dao/mysql"
	"Goi/models"
	"Goi/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成post id
	p.ID = snowflake.GenID()
	// 从 c 取到当前发请求的用户的ID

	// 2. 保存到数据库并返回
	return mysql.CreatePost(p)
}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	data = new(models.ApiPostDetail)
	// 查询并组合接口想用的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed",
			zap.Int64("pid", pid), zap.Error(err))
		return
	}

	// 根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
			zap.Int64("author", post.AuthorID),
			zap.Error(err))
		return
	}

	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailById(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById(post.CommunityID) failed",
			zap.Int64("community", post.CommunityID),
			zap.Error(err))
		return
	}

	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}

	return
}

// GetPostList 获取帖子列表
func GetPostList(offset, limit int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(offset, limit)
	fmt.Println(posts)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author", post.AuthorID),
				zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById(post.CommunityID) failed",
				zap.Int64("community", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return data, err
}
