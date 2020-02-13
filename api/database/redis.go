package database

import (
	"fmt"
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

func (r *RedisHandler) GetCachedAlbum(username string) ([]string, error) {
	c := r.Pool.Get()
	defer c.Close()

	album, e := redis.Strings(c.Do("SMEMBERS", username))
	if e != nil {
		fmt.Println(e)
		return nil, e
	}

	return album, e
}

func (r *RedisHandler) CacheAlbum(username string, album []string) error {
	c := r.Pool.Get()
	defer c.Close()

	c.Send("MULTI")
	for _, key := range album {
		c.Send("SADD", username, key)
	}

	_, e := c.Do("EXEC")
	if e != nil {
		fmt.Println(e)
	}

	return e
}

func (r *RedisHandler) AddKeyToCache(username string, key string) error {
	c := r.Pool.Get()
	defer c.Close()

	_, e := c.Do("SADD", username, key)
	return e
}

func (r *RedisHandler) RemoveKeyFromCache(username string, key string) error {
	c := r.Pool.Get()
	defer c.Close()

	_, e := c.Do("SREM", username, key)
	return e
}
