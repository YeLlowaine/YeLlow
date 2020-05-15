package gredis

import (
	"github.com/Hallelujah1025/Stroke-Survivors/pkg/setting"
	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:   setting.MaxIdle,
		MaxActive: setting.MaxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "")
			if err != nil {
				return nil, err
			}
			if setting.Password != 0 {
				if _, err := c.Do("AUTH", setting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
	return nil
}
