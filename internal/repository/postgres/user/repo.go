package postgresuser

import (
	"goProject/internal/adapter/redis"
	"goProject/internal/repository/postgres"
)

type Repo struct {
	redis *redis.Adapter
	db    postgres.DBTX
}

func New(db postgres.DBTX, redis *redis.Adapter) *Repo {
	return &Repo{db: db, redis: redis}
}
