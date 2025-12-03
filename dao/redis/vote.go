package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 //每一票的分数
)

var (
	ErrVoteTimeExpire = errors.New("超出帖子投票时间限制")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

// CreatePost
func CreatePost(postID,communityID int64) error {
	pipeline := client.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 初始化帖子的投票总数
	pipeline.Set(getRedisKey(KeyPostVoteSumPF+strconv.FormatInt(postID, 10)), 0, 0)
	// 更新：把帖子id加到社区的set
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断投票限制
	// 去redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2和3需要放到一个pipeline事务中操作

	// 2. 更新贴子的分数
	// 先查当前用户给当前帖子的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()

	// 更新：如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	// 更新帖子的投票总数
	pipeline.IncrBy(getRedisKey(KeyPostVoteSumPF+postID), int64(value-ov))

	// 3. 记录用户为该贴子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}

func GetUserVotesForPosts(userID string, postIDs []string) (data []int8, err error) {
	pipeline := client.Pipeline()
	for _, postID := range postIDs {
		key := getRedisKey(KeyPostVotedZSetPF + postID)
		pipeline.ZScore(key, userID)
	}
	cmders, err := pipeline.Exec()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	data = make([]int8, len(cmders))
	for i, cmder := range cmders {
		res, err := cmder.(*redis.FloatCmd).Result()
		if err == redis.Nil {
			// No vote recorded for this post
			data[i] = 0
		} else if err != nil {
			return nil, err
		} else {
			data[i] = int8(res)
		}
	}
	return data, nil
}