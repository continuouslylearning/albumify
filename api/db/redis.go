package db

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

type RedisHandler struct {
	Pool *redis.Pool
}

func InitializeRedis(r *gin.Engine) *RedisHandler {
	url := os.Getenv("REDIS_URL")
	h := &RedisHandler{
		Pool: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.DialURL(url)
			},
			IdleTimeout: 2000 * time.Second,
			MaxIdle:     3,
		},
	}
	r.Use(AddRedisHandler(h))

	return h
}

func AddRedisHandler(h *RedisHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redis", h)
		c.Next()
	}
}

func (r *RedisHandler) GetCachedURLs(keys []string) ([]string, error) {
	conn := r.Pool.Get()
	defer conn.Close()

	conn.Send("MULTI")
	for _, key := range keys {
		conn.Send("GET", key)
	}
	URLs, e := redis.Strings(conn.Do("EXEC"))
	if e != nil {
		return nil, e
	}

	return URLs, nil
}

func (r *RedisHandler) CacheURLs(keys []string, presignedURLS []string) error {
	c := r.Pool.Get()
	defer c.Close()

	c.Send("MULTI")
	for _, key := range keys {
		c.Send("BITCOUNT", key)
	}
	reply, e := redis.Ints(c.Do("EXEC"))

	c.Send("MULTI")
	for i, bitcount := range reply {
		if bitcount == 0 {
			c.Send("SET", keys[i], presignedURLS[i], "EX", 3000)
		}
	}
	_, e = c.Do("EXEC")

	return e
}

func (r *RedisHandler) GetCachedAlbum(username string) ([]string, bool, error) {
	c := r.Pool.Get()
	defer c.Close()

	c.Send("MULTI")
	c.Send("EXISTS", username)
	c.Send("SMEMBERS", username)
	c.Send("EXPIRE", username, 3000)
	replies, e := redis.Values(c.Do("EXEC"))
	if e != nil {
		return nil, false, e
	}

	exists := replies[0].(int64)
	if exists == 0 {
		return nil, false, nil
	}

	album, _ := redis.Strings(replies[1], nil)
	return album, true, nil
}

func (r *RedisHandler) CacheAlbum(username string, album []string) error {
	c := r.Pool.Get()
	defer c.Close()

	c.Send("MULTI")
	for _, key := range album {
		c.Send("SADD", username, key)
	}
	c.Send("EXPIRE", username, 3000)
	_, e := c.Do("EXEC")

	return e
}

func (r *RedisHandler) AddKeys(username string, keys []string) error {
	c := r.Pool.Get()
	defer c.Close()

	ok, e := redis.Int64(c.Do("EXISTS", username))
	if e != nil {
		return e
	}

	if ok == 1 {
		c.Send("MULTI")
		for _, key := range keys {
			c.Send("SADD", username, fmt.Sprintf("%s/%s", username, key))
		}
		c.Send("EXPIRE", username, 3000)
		_, e = c.Do("EXEC")
		if e != nil {
			return e
		}
	}

	return nil
}

func (r *RedisHandler) RemoveKey(username string, key string) error {
	c := r.Pool.Get()
	defer c.Close()

	c.Send("MULTI")
	c.Send("SREM", username, fmt.Sprintf("%s/%s", username, key))
	_, e := c.Do("EXEC")

	return e
}
