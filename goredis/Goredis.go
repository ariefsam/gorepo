package goredis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Goredis struct {
	Host     string
	Port     string
	Password string
	Database int
}

func (g Goredis) client() (redisClient *redis.Client) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     g.Host + ":" + g.Port,
		Password: g.Password,
		DB:       g.Database,
	})
	return
}

func (g Goredis) GetInt(key string) (data int, err error) {
	data, err = g.client().Get(key).Int()
	return
}

func (g Goredis) GetString(key string) (data string, err error) {
	data, err = g.client().Get(key).Result()
	return
}

func (g Goredis) Set(key string, data interface{}) (err error) {
	g.client().Set(key, data, 0)
	return
}

func (g Goredis) SetEx(key string, data interface{}, expireSecond int) (err error) {
	duration, _ := time.ParseDuration(fmt.Sprint(expireSecond) + "s")
	g.client().Set(key, data, duration)
	return
}

func (g Goredis) PFAdd(key string, data ...interface{}) (result int64, err error) {
	res := g.client().PFAdd(key, data...)
	result = res.Val()
	err = res.Err()
	return
}

func (g Goredis) PFCount(key ...string) (result int64, err error) {
	result, err = g.client().PFCount(key...).Result()
	return
}
