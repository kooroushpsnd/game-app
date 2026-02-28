package redisadaptor

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	DialTimeout    time.Duration   `koanf:"redis_dial_timeout"`
	ReadTimeout    time.Duration   `koanf:"redis_read_timeout"`
	WriteTimeout   time.Duration   `koanf:"redis_write_timeout"`
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
		DialTimeout: config.DialTimeout,
		ReadTimeout: config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	})

	log.Println("Redis Started Successfully")
	return &Adapter{client: rdb ,config: config}
}

func (a *Adapter) Client() *redis.Client {
	return a.client
}