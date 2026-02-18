package emailservice

import (
	"context"
	emailcodedto "goProject/internal/dto/email_code"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
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
	if(expiredEmail || err != nil){
		return  richerror.New(op).WithErr(err)
	}

	OTPCode ,err := s.mailerSvc.GenerateOTP(ctx ,3)
	s.mailerSvc.SendEmail(ctx ,req.Email ,"EMAIL" ,OTPCode)
	return nil
}

func (s *Service) VerifyEmailCode(ctx context.Context ,req emailcodedto.VerifyEmailCodeDto) (bool ,error){

}