package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

// 推荐阅读
// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm
// 投票功能：
// 1. 用户投票的数据

// 本项目使用简化版的投票分数
// 投一票就加432分 86400/200  --> 200张赞成票就可以给你的帖子续一天

/*
投票的几种情况：
direction = 1 时，有两种情况：       新旧差值的绝对值
	1. 之前没有投过票，现在投赞成票    1 (+432)
	2. 之前投反对票，现在投赞成票      2 (+432*2)
direction = 0 时，有两种情况：
	1. 之前投赞成票，现在要取消投票    1 (-432)
	2. 之前投反对票，现在要取消投票    1 (432)
direction = -1 时，有两种情况：
	1. 之前没有投过票，现在投反对票    1 (-432)
	2. 之前投过赞成票，现在投反对票    2 (-432*2)

投票的限制🚫：
每个帖子自发表之日起，一个星期内允许用户投票，超过一个星期就不允许再投票了。
	1. 到期之后将redis中保存的赞成票数及反对票数存储到MySQL表中
	2. 到期后删除 KeyPostVotedZSetPrefix

投票之后 --> 更新分数和投票记录
*/

const (
	VoteExpireTimeInSeconds = 7 * 24 * 3600 // 投票过期时间 -- 按秒记的一周
	scorePerVote            = 432           // 每票占的分数
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func VoteForPost(userID, postID string, value float64) error {
	// 1、判断投票的限制
	// 去redis取帖子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostScoreZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > VoteExpireTimeInSeconds {
		return ErrVoteTimeExpire
	}
	// 2、更新帖子的分数
	// 先查当前用户给当前帖子的投票记录
	preScore := rdb.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()

	var direction float64
	diff := math.Abs(preScore - value)
	if value > preScore {
		direction = 1
	} else {
		direction = -1
	}
	// ?                                      (value - preScore)*scorePerVote
	_, err := rdb.ZIncrBy(getRedisKey(KeyPostScoreZSet), direction*diff*scorePerVote, postID).Result()
	if ErrVoteTimeExpire != nil {
		return err
	}
	// 3、记录用户为该帖子投票的数据
	if value == 0 {
		rdb.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), postID).Result()
	}
	rdb.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
		Score:  value, // 赞成票还是反对票
		Member: userID,
	}).Result()
	return err
}
