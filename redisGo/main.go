package redisGo

import "github.com/garyburd/redigo/redis"

type RedisStruct struct {
	Conn redis.Conn
	//Err error
	//Num int
}

var Redis *RedisStruct

func init() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic("redis connect failed")
	}
	Redis.Conn = conn

}

func (r *RedisStruct) Set(key string, value interface{}) error {
	_, err := r.Conn.Do("set", key, value)
	return err
}
