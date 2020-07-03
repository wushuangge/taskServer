package store

import (
	"errors"
	"fmt"
	"time"
	//"github.com/gomodule/redigo/redis"
	"github.com/go-redis/redis"
	"task/config"
)
type RedisClient struct {
	//c redis.Conn
	c *redis.Client
}

var redisdb *RedisClient

func InitRedis() error{
	redisdb = new(RedisClient)
	redisdb.c = newLoginRedisClient(100)
	if redisdb.c != nil {
		return errors.New("init redis err")
	}
	return nil
}

//func newLoginRedisClient() (redis.Conn) {
//	c ,err := redis.Dial("tcp",config.GetLoginRedisAddr())
//	if err != nil {
//		fmt.Println("connect redis error :",err)
//		return nil
//
//	}
//	if _, err := c.Do("AUTH", config.GetLoginRedisPassword()); err != nil {
//		c.Close()
//		return nil
//	}
//	return c
//}


func newLoginRedisClient(poolSize int) *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:         config.GetLoginRedisAddr(),
		Password:     config.GetLoginRedisPassword(),
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     poolSize,
		PoolTimeout:  30 * time.Second,
	})
	_, err := c.Ping().Result()
	if err != nil {
		panic(err)
	}
	return c
}

func RegisterTask(key string, value []byte) string {
	err := redisdb.c.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("error:", err)
		return err.Error()
	}
	return "ok"
}

func UpdateTask(key string, value []byte) string {
	res, err := redisdb.c.Get(key).Result()
	if len(res) == 0{
		return "id not exist"
	}

	err = redisdb.c.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("error:", err)
		return err.Error()
	}
	SelectTask(key)
	return "ok"
}

func SelectTask(key string) (string, error) {
	res, err := redisdb.c.Get(key).Result()
	fmt.Println(res)
	return res, err
}