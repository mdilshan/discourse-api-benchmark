package discourse

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var namespace = "discourse-go:"

func SetCache(key string, value string) {
    redisClient.Set(ctx, namespace+key, value, 0)
}

func GetCache(key string) (string, error) {
    return redisClient.Get(ctx, namespace+key).Result()
}
