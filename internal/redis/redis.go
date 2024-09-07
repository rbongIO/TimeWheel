package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var DB redis.Conn
var Pool *redis.Pool

func init() {
	var err error

	Pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "172.29.1.113:6371", redis.DialPassword("redis_123456"))
		},
		MaxIdle:   5,
		MaxActive: 20,
	}
	DB, err = redis.Dial("tcp", "172.29.1.113:6371",
		redis.DialClientName("test"), redis.DialPassword("redis_123456"),
		redis.DialDatabase(1), redis.DialConnectTimeout(10*time.Second))
	if err != nil {
		panic(fmt.Sprintln("Dial redis failed,err:", err))
	}

}
