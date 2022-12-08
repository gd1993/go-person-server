package config

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

func RedisConnPool() (pool *redis.Pool) {
	pool = &redis.Pool{
		//最大空闲连接数
		MaxIdle: 30,
		//在给定时间内，允许分配的最大连接数（当为零时，没有限制）
		MaxActive: 30,
		//在给定时间内，保持空闲状态的时间，若到达时间限制则关闭连接（当为零时，没有限制）
		IdleTimeout: 200,
		//创建和配置链接
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:49156")
			if err != nil {
				return nil, err
			}
			//如果redis设置了用户密码，使用auth指令
			if _, err := c.Do("AUTH", "redispw"); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return pool

}
