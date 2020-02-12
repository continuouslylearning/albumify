package database

import (
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisHandler *RedisHandler

type RedisHandler struct {
	Pool *redis.Pool
}

func InitializeRedis() *RedisHandler {
	url := os.Getenv("REDIS_URL")
	redisHandler = &RedisHandler{
		Pool: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.DialURL(url)
			},
			IdleTimeout: 2000 * time.Second,
			MaxIdle:     3,
		},
	}

	return redisHandler
}

func (r *RedisHandler) GetCachedURLs(keys []string) ([]string, error) {
	conn := r.Pool.Get()
	defer conn.Close()

	conn.Send("MULTI")
	for _, key := range keys {
		conn.Send("GET", key)
	}
	reply, e := redis.Strings(conn.Do("EXEC"))
	if e != nil {
		return nil, e
	}

	return reply, nil
}

func (r *RedisHandler) CacheURLs(keys []string, presignedURLS []string) error {
	conn := r.Pool.Get()
	defer conn.Close()

	conn.Send("MULTI")
	for _, key := range keys {
		conn.Send("BITCOUNT", key)
	}
	reply, e := redis.Ints(conn.Do("EXEC"))

	conn.Send("MULTI")
	for i, count := range reply {
		if count == 0 {
			conn.Send("SET", keys[i], presignedURLS[i], "EX", 3000)
		}
	}
	_, e = conn.Do("EXEC")

	return e
}
