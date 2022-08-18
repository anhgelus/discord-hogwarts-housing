package database

import (
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

var RedisAddr = os.Args[3]

func GetRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 60 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", RedisAddr+":6379") },
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func RedisSet(pool *redis.Pool, k string, v string) (string, error) {
	c := pool.Get()
	defer c.Close()
	return redis.String(c.Do("SET", k, v))
}

func RedisGet(pool *redis.Pool, k string) (string, error) {
	c := pool.Get()
	defer c.Close()
	return redis.String(c.Do("GET", k))
}

func RedisHset(pool *redis.Pool, k string, v interface{}) (string, error) {
	c := pool.Get()
	defer c.Close()
	return redis.String(c.Do("HSET", redis.Args{}.Add(k).AddFlat(v)...))
}

func RedisHget(pool *redis.Pool, k string, f string) (string, error) {
	c := pool.Get()
	defer c.Close()
	return redis.String(c.Do("HGET", k, f))
}

func RedisHgetAll(pool *redis.Pool, k string) ([]interface{}, error) {
	c := pool.Get()
	defer c.Close()
	return redis.Values(c.Do("HGETALL", k))
}
