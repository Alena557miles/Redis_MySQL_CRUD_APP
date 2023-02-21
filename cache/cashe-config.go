package cache

import (
	"fmt"
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

var instance *redis.Client = nil

func GetInstance() *redis.Client {
	if instance == nil {
		fmt.Println("Creating single Redis Client instance now.")
		instance = GetClient()
		return instance
	} else {
		fmt.Println("single Redis Client instance already exist.")
		return instance
	}
}
