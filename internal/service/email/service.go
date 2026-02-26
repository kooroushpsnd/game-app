package emailservice

import (
	"context"
	"goProject/internal/config"
	emailcodedto "goProject/internal/dto/email_code"
	"goProject/internal/entity"
	userservice "goProject/internal/service/user"
)

type MailerAdaptor interface {
	HashOTP(email, otp string) (string, error)
	GenerateAndHashOTP(ctx context.Context, email string, digits int) (string, string, error)
	SendEmail(ctx context.Context, to, subject, body string) error
}

type Repository interface {
	CheckEmailCodeReadyToResend(ctx context.Context, email string) (bool, error)
	CheckEmailCodeExpiration(ctx context.Context, email string) (bool, error)
	GetLatestEmailCode(ctx context.Context, email string) (entity.EmailCode, error)
	CreateEmailCode(ctx context.Context, req emailcodedto.CreateEmailCodeDto) error
	UpdateEmailCode(ctx context.Context, email string, req emailcodedto.UpdateEmailCodeRequestDto) (entity.EmailCode, error)
	IncrementEmailCodeAttempts(ctx context.Context, email string) (entity.EmailCode, error)
}

type Service struct {
	mailerSvc MailerAdaptor
	userSvc   *userservice.Service
	repo      Repository
	config    config.Application
}

func New(userService *userservice.Service, repo Repository, mailerService MailerAdaptor, cfg config.Application) *Service {
	return &Service{mailerSvc: mailerService, userSvc: userService, repo: repo, config: cfg}
}
