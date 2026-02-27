package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	UserCacheKeyID    string        `koanf:"user_cache_key_id"`
	UserCacheKeyEmail string        `koanf:"user_cache_key_email"`
	UserCacheTTL      time.Duration `koanf:"user_cache_ttl"`
}

type Adapter struct {
	config Config
	client *redis.Client
}

func New(config Config) *Adapter {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return &Adapter{client: rdb ,config: config}
}

func (a *Adapter) Client() *redis.Client {
	return a.client
}

func (a *Adapter) Config() Config {
	return a.config
}