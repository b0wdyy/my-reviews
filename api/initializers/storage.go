package initializers

import "github.com/gofiber/storage/redis/v2"

var RedisStorage *redis.Storage

func InitStorage() {
	RedisStorage = redis.New()
}
