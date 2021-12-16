package app

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

// AppCache struct
type AppCache struct {
	cache      *cache.Cache
	defaultTTL time.Duration
}

// CreateAppCache
func CreateAppCache(cfg *AppConfig) (ac *AppCache) {
	ac = new(AppCache)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisServer,
		Password: cfg.RedisPassword,
	})

	ac.cache = cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(cfg.LocalCacheSize, cfg.LocalCacheTTL),
	})

	ac.defaultTTL = cfg.RedisTTL

	return
}

// Get function
func (ac *AppCache) Get(ctx context.Context, key string, value interface{}) (err error) {
	err = ac.cache.Get(ctx, key, value)
	return
}

// Set function
func (ac *AppCache) Set(ctx context.Context, key string, value interface{}) {
	ac.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   ac.defaultTTL,
	})
}
