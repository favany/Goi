package mysql

import (
	"Goi/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id, title, content, author_id, community_id)
	values (?,?,?,?,?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostById 根据 id 查询单个帖子数据
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select 
				post_id, title, content, author_id, community_id, create_time
				from post
				where post_id = ?
				`
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList 查询帖子列表函数
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select
				post_id, title, content, author_id, community_id, create_time
				from post
				ORDER BY create_time DESC
				limit ?,?
				`
	posts = make([]*models.Post, 0, 2) // 长度 容量
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

func GetPostListByIdList(IdList []string) (postList []*models.Post, err error) {
	sqlStr := `SELECT post_id, title, content, author_id, community_id, create_time
				FROM post
				WHERE post_id in (?)
				ORDER BY FIND_IN_SET(post_id, ?)
				`

	// https://www.liwenzhou.com/posts/Go/sqlx/
	query, args, err := sqlx.In(sqlStr, IdList, strings.Join(IdList, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	err = db.Select(&postList, query, args...) // !!!!!
	return
}
