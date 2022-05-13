package redis

import (
	"Goi/models"
	"github.com/go-redis/redis"
)

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从 redis 获取 id
	// 1. 根据用户请求中携带的 order 参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	// 2. 确定查询的索引起始列
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	// 3. ZRevRange 按分数从大到小查询指定的分数
	return rdb.ZRevRange(key, start, end).Result()
}

// GetPostVoteData 根据ids查询每篇帖子的投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//keys := make([]string, 0, len(ids))

	// 使用pipeline一次发送多条命令，减少RTT
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		// 查找 key 中分数是 1 元素的数量 -> 统计每篇帖子的赞成票的数量
		pipeline.ZCount(getRedisKey(KeyPostVotedZSetPrefix+id), "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))

	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
