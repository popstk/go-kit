package rediskit

import "github.com/go-redis/redis"

func NewRedisClient(u string) *redis.Client {
	opt, err := redis.ParseURL(u)
	if err != nil {
		panic(err)
	}

	// notice: no read timeout
	opt.ReadTimeout = -1
	return redis.NewClient(opt)
}

