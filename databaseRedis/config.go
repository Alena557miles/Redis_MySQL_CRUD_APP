package databaseRedis

import "github.com/redis/go-redis/v9"

var Opt = &redis.Options{
	Addr:     "localhost:6379",
	Password: "mycomplicatedpassword",
	DB:       0,
}
