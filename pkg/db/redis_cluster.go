package db

import "github.com/redis/go-redis/v9"

// redis Cluster
func NewRedisClusterClient(options *Options) *redis.ClusterClient {

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:           options.Addrs,
		Password:        options.Password,
		DialTimeout:     options.DialTimeout,
		ReadTimeout:     options.ReadTimeout,
		WriteTimeout:    options.WriteTimeout,
		PoolSize:        options.PoolSize, // go-redis 连接池大小为 runtime.GOMAXPROCS * 10，在大多数情况下默认值已经足够使用，
		Username:        options.Username,
		MaxRetries:      0,                         //命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * options.DialTimeout,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * options.DialTimeout, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
	})

	return client
}
