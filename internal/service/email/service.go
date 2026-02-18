package emailservice

import (
	"context"
	emailcodedto "goProject/internal/dto/email_code"
	"goProject/internal/entity"
	userservice "goProject/internal/service/user"
)

type MailerAdaptor interface {
	GenerateOTP(ctx context.Context, n int) (string, error)
	SendEmail(ctx context.Context, to, subject, body string) error
}

type Repository interface {
	CheckEmailCodeReadyToResend(ctx context.Context, email string) (bool, error)
	CheckEmailCodeExpiration(ctx context.Context, email string) (bool, error)
	GetLatestEmailCode(ctx context.Context, email string) (entity.EmailCode, error)
	CreateEmailCode(ctx context.Context, req emailcodedto.CreateEmailCodeDto) error
	UpdateEmailCode(ctx context.Context, email string, req emailcodedto.UpdateEmailCodeRequestDto) (entity.EmailCode, error)
}

type Service struct {
	mailerSvc MailerAdaptor
	userSvc   *userservice.Service
	repo      Repository
}

func New(userService *userservice.Service, repo Repository, mailerService MailerAdaptor) *Service {
	return &Service{mailerSvc: mailerService, userSvc: userService, repo: repo}
}
