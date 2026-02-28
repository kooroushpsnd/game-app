package redisuser

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	UserCacheKeyID    string        `koanf:"user_cache_key_id"`
	UserCacheKeyEmail string        `koanf:"user_cache_key_email"`
	UserCacheTTL      time.Duration `koanf:"user_cache_ttl"`
}

type UserCache struct {
	client *redis.Client
	config Config
}

func New(client *redis.Client, cfg Config) *UserCache {
	return &UserCache{client: client, config: cfg}
}