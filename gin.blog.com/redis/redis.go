package redis

import (
	"gin.blog.com/config"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

var client *redis.Client

func InitRedis() error {
	conf := config.GetConfig()
	client = redis.NewClient(&redis.Options{
		Addr:     conf.RedisConfig.Addr,
		Password: conf.RedisConfig.Password,
		DB:       conf.RedisConfig.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		err = errors.Wrap(err, "InitRedis")
		return err
	}
	return nil
}

//初始化调用全局
func GetRedis() *redis.Client {
	return client
}
