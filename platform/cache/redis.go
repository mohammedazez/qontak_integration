package cache

import (
	"qontak_integration/pkg/configs"
	"qontak_integration/pkg/utils"
	"runtime"

	"github.com/gofiber/storage/redis/v3"
	"golang.org/x/net/context"
)

var fiberRedisConn *redis.Storage

// RedisConnection func for get connect to Redis server.
func RedisConnection() (*redis.Storage, error) {
	if fiberRedisConn == nil {
		return connection()
	}
	_, err := fiberRedisConn.Conn().Ping(context.Background()).Result()
	if err != nil {
		return connection()
	}
	return fiberRedisConn, nil

}

// connection func for connect to Redis server.
func connection() (*redis.Storage, error) {
	// Define Redis database number.
	dbNumber := configs.Config.Redis.Database
	// Build Redis connection URL.
	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		return nil, err
	}
	store := redis.New(redis.Config{
		Addrs:    []string{redisConnURL},
		Password: configs.Config.Redis.Password,
		Database: dbNumber,
		PoolSize: 10 * runtime.GOMAXPROCS(0),
	})

	fiberRedisConn = store
	_, err = fiberRedisConn.Conn().Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return fiberRedisConn, nil
}
