package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

type Options struct {
	Addrs        []string
	DB           int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	Username     string
	Password     string
}

func defaultOptions() *Options {
	return &Options{
		Addrs:       []string{"127.0.0.1:6379"},
		DB:          0,
		DialTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
		PoolSize:    50,
	}
}

func NewRedisClient(opt *Options) redis.UniversalClient {
	if opt == nil {
		opt = defaultOptions()
	}
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:           opt.Addrs,
		Password:        opt.Password,
		DB:              opt.DB,
		DialTimeout:     opt.DialTimeout,
		ReadTimeout:     opt.ReadTimeout,
		WriteTimeout:    opt.WriteTimeout,
		PoolSize:        opt.PoolSize, // go-redis 连接池大小为 runtime.GOMAXPROCS * 10，在大多数情况下默认值已经足够使用，
		Username:        opt.Username,
		MaxRetries:      0,                      //命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
	})
	return client
}

var (
	once        sync.Once
	redisClient redis.UniversalClient
)

func GetRedisOr(opt *Options) redis.UniversalClient {
	if opt == nil && redisClient == nil {
		panic("redis client not initialized")
	}

	once.Do(func() {
		redisClient = NewRedisClient(opt)
	})

	if err := redisClient.Ping(context.TODO()).Err(); err != nil {
		log.Fatalf("failed to connect redis: %v", err)
	}

	return redisClient
}
