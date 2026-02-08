package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	userdto "goProject/internal/dto/user"
	"goProject/internal/entity"
)

type Repository interface {
	Register(ctx context.Context ,u entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context ,email string) (entity.User, error)
	GetUserByID(ctx context.Context, userID uint) (entity.User, error)
	GetAllUsers(ctx context.Context, dto userdto.GetAllRequestUserDto) ([]entity.User, error)
	IsEmailUnique(ctx context.Context ,email string) (bool ,error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

func New(authGenerator AuthGenerator, repo Repository) *Service {
	return &Service{auth: authGenerator, repo: repo}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}