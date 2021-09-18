package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"sync"
)



var instance *redis.Client
var once sync.Once


func GetInstance() *redis.Client {
	once.Do(func(){
		client := redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis_queue.host"),
			Password: viper.GetString("redis_queue.password"), // no password set
			DB:       0,  // use default DB
		})
		instance = client
	})
	return instance
}
