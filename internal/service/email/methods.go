package emailservice

import (
	"context"
	emailcodedto "goProject/internal/dto/email_code"
	userdto "goProject/internal/dto/user"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
	"time"
)

func (s *Service) SendEmail(ctx context.Context ,req emailcodedto.SendEmailCodeDto) error{
	const op = "emailservice.SendEmail"
	user ,err := s.userSvc.GetUserByEmail(ctx ,req.Email)
	if err != nil {
		return  err
	}
	if(user.EmailVerify == true){
		return  richerror.New(op).WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsg_UserAlreadyVerified)
	}
	expiredEmail ,err := s.repo.CheckEmailCodeReadyToResend(ctx ,req.Email)
	if(!expiredEmail || err != nil){
		return  richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsg_TooManyAttempts)
	}

	otp, hash, err := s.mailerSvc.GenerateAndHashOTP(ctx, req.Email, 6)
	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	if err := s.repo.CreateEmailCode(ctx ,emailcodedto.CreateEmailCodeDto{
		Email:          req.Email,
		HashCode:       hash,
		UserID:         user.ID,
	}); err != nil {
		return richerror.New(op).WithErr(err)
	}

	if err := s.mailerSvc.SendEmail(ctx, req.Email, "EMAIL", otp); err != nil {
		return richerror.New(op).WithErr(err)
	}
	return nil
}

func (s *Service) VerifyEmailCode(ctx context.Context, req emailcodedto.VerifyEmailCodeDto) (bool, error) {
	const op = "emailservice.VerifyEmailCode"

	user ,err := s.userSvc.GetUserByEmail(ctx ,req.Email)
	if err != nil {
		return false ,err
	}
	if(user.EmailVerify == true){
		return false, richerror.New(op).WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsg_UserAlreadyVerified)
	}

	emailCode, err := s.repo.GetLatestEmailCode(ctx, req.Email)
	if err != nil {
		return false, richerror.New(op).WithErr(err)
	}

	if emailCode.Status != entity.EmailCodeStatusActive {
		return false, richerror.New(op).
			WithKind(richerror.KindInvalid).
			WithMessage(errmsg.ErrorMsg_EmailCodeNotActive)
	}

	if time.Now().After(emailCode.ExpirationDate) {
		return false, richerror.New(op).
			WithKind(richerror.KindInvalid).
			WithMessage(errmsg.ErrorMsg_EmailCodeNotActive)
	}

	if emailCode.Attempts >= s.config.EmailCodeAttemptsMax {
		return false, richerror.New(op).
			WithKind(richerror.KindInvalid).
			WithMessage(errmsg.ErrorMsg_TooManyAttempts)
	}

	inputHash ,err := s.mailerSvc.HashOTP(req.Email, req.Code)
	if err != nil {
		return false, richerror.New(op).WithErr(err)
	}

	if inputHash != emailCode.HashCode {
		_, err := s.repo.IncrementEmailCodeAttempts(ctx, emailCode.Email)
		if err != nil {
			return false, richerror.New(op).WithErr(err)
		}
	}

	status := entity.EmailCodeStatusVerified
	if _ ,err := s.repo.UpdateEmailCode(ctx, emailCode.Email ,emailcodedto.UpdateEmailCodeRequestDto{
		Status: &status,
	}); err != nil {
		return false, richerror.New(op).WithErr(err)
	}

	emailCerify := true
	if _ ,err := s.userSvc.Update(ctx, user.ID ,userdto.UserUpdatePatch{
		EmailVerify: &emailCerify,
	}); err != nil {
		return false, richerror.New(op).WithErr(err)
	}

	return true, nil
}
