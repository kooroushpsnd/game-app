package emailservice

import (
	"context"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (s *Service) SendEmail(ctx context.Context ,email string) error{
	const op = "emailservice.SendEmail"
	user ,err := s.userSvc.GetUserByEmail(ctx ,email)
	if err != nil {
		return  err
	}
	if(user.EmailVerify == true){
		return  richerror.New(op).WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsg_UserAlreadyVerified)
	}
	expiredEmail ,err := s.repo.CheckEmailCodeReadyToResend(ctx ,email)
	if(expiredEmail || err != nil){
		return  richerror.New(op).WithErr(err)
	}

	OTPCode ,err := s.emailSvc.GenerateOTP(ctx ,3)
	s.emailSvc.SendEmail(ctx ,email ,"EMAIL" ,OTPCode)

}