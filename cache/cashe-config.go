package cache

import (
	"github.com/redis/go-redis/v9"
)

var Opt = &redis.Options{
	Addr:     "localhost:6379",
	Password: "mycomplicatedpassword",
	DB:       0,
}

func GetClient() *redis.Client {
	return redis.NewClient(Opt)
}
