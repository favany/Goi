package logic

import (
	"Goi/dao/mysql"
	"Goi/dao/redis"
	"Goi/models"
	"Goi/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成post id
	p.ID = snowflake.GenID()

	// 2. 保存到数据库并返回
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)

	return
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

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 去 redis 查询 id 列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return no data")
	}
	fmt.Println(ids)
	// 根据id 去MySQL 数据库查询帖子的详细信息
	// 返回的数据要根据给定的id顺序返回
	posts, err := mysql.GetPostListByIdList(ids)
	if err != nil {
		return
	}
	// 提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	fmt.Println("voteData", voteData)
	if err != nil {
		return
	}

	// 将帖子的作者及分区信息查询出来，填充到帖子中
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for index, post := range posts {
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
			LikeCount:       voteData[index],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return data, err
}

func GetCommunityPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 去 redis 查询 id 列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return no data")
	}
	fmt.Println(ids)
	// 根据id 去MySQL 数据库查询帖子的详细信息
	// 返回的数据要根据给定的id顺序返回
	posts, err := mysql.GetPostListByIdList(ids)
	if err != nil {
		return
	}
	// 提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	fmt.Println("voteData", voteData)
	if err != nil {
		return
	}

	// 将帖子的作者及分区信息查询出来，填充到帖子中
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for index, post := range posts {
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
			LikeCount:       voteData[index],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return data, err
}

// GetPostListNew 将两个查询接口合二为一的函数
func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	if p.CommunityID == 0 {
		// 获取数据
		data, err = GetPostList2(p)
	} else {
		// 获取数据
		data, err = GetCommunityPostList(p)
	}

	if err != nil {
		zap.L().Error("GetPostListNew failed", zap.Error(err))
	}

	return
}
