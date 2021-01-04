package redis

import (
	"context"
	// 	"fmt"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

var ctx = context.Background()
var rdb *redis.Client

func ClientInit() {
	redis_url := os.Getenv("REDIS_URL")
	fmt.Println(redis_url)
	if redis_url == "" {
		redis_url = "joseph-test.coe3c4.ng.0001.use1.cache.amazonaws.com:6379"
	}
	fmt.Println(redis_url)
	rdb = redis.NewClient(&redis.Options{
		Addr:     redis_url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func Set(key, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}

func Incr(key string) (int64, error) {
	return rdb.Incr(ctx, key).Result()
}
