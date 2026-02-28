package postgresuser

import (
	"context"
	"goProject/internal/entity"
	"goProject/internal/repository/postgres"
)

type UserCache interface {
	GetByID(ctx context.Context, userID uint) (entity.User, bool, error)
	GetByEmail(ctx context.Context, email string) (entity.User, bool, error)
	Set(ctx context.Context, user entity.User) error
	DeleteByID(ctx context.Context, userID uint) error
	DeleteByEmail(ctx context.Context, email string) error
}

type Repo struct {
	userCache UserCache
	db        postgres.DBTX
}

func New(db postgres.DBTX, userCache UserCache) *Repo {
	return &Repo{db: db, userCache: userCache}
}
