package redis

// 定义redis key.注意使用命名空间的方式方便查询和拆分
const (
	Prefix             = "bluebell"
	KeyPostTimeZSet    = "post:time"  //zset;帖子及发帖时间
	KeyPostScoreZSet   = "post:score" //zset;帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted" //zest;记录用户及投票类型；参数是post id
	KeyPostVoteSumPF   = "post:votesum:" //string;记录帖子的投票总数；参数是post id
	KeyCommunitySetPF  = "community"  //set;保存每个分区下帖子的id
)

// getRedisKey 给redis key加前缀
func getRedisKey(key string) string {
	return Prefix + key
}
