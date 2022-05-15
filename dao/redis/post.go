package redis

import (
	"Goi/models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func getIDsFromKey(key string, page, size int64) ([]string, error) {
	// 1. 确定查询的索引起始列
	start := (page - 1) * size
	end := start + size - 1

	// 2. ZRevRange 按分数从大到小查询指定的分数
	return rdb.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从 redis 获取 id
	// 1. 根据用户请求中携带的 order 参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	// 2. 确定查询的索引起始点
	return getIDsFromKey(key, p.Page, p.Size)
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

// GetCommunityPostIDsInOrder 按社区根据Ids查询每篇帖子的投赞成票的情况
func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	// 使用 zinterscore 把分区的帖子 set 与帖子分数的 zset 生成一个新的 zset
	// 针对新的 zset 按之前的逻辑取数据

	// 社区的key
	communityKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(p.CommunityID)))

	// 利用缓存 key 减少 zinterscore 执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if rdb.Exists(orderKey).Val() < 1 {
		// 不存在，需要计算
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, communityKey) // zinterscore 计算
		pipeline.Expire(key, 60*time.Second) // 设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}

	}

	return getIDsFromKey(key, p.Page, p.Size)
}
