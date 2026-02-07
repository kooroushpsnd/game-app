package userservice

import (
	"context"
	userdto "goProject/internal/dto/user"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (s Service) Login(ctx context.Context, req userdto.LoginRequestDto) (userdto.LoginResponseDto, error) {
	const op = "userservice.login"

	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return userdto.LoginResponseDto{}, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindNotFound).
			WithMessage(errmsg.ErrorMsg_UserNotFound).
			WithMeta(map[string]interface{}{
				"Email": req.Email,
			})
	}

	if user.Password != getMD5Hash(req.Password) {
		return userdto.LoginResponseDto{}, richerror.New(op).
			WithErr(err).WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsg_WrongPassword)
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return userdto.LoginResponseDto{}, richerror.New(op).
			WithErr(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsg_SomethingWentWrong)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return userdto.LoginResponseDto{}, richerror.New(op).
			WithErr(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsg_SomethingWentWrong)
	}

	return userdto.LoginResponseDto{
		User: userdto.UserInfoDto{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
		Tokens: userdto.TokensDto{
			AccessToken:  accessToken,
			RefreshToken: refreshToken},
	}, nil
}

func (s Service) Register(ctx context.Context, req userdto.SignupRequestDto) (userdto.SignupResponseDto, error) {
	const op = "userservice.register"

	value , err := s.repo.IsEmailUnique(ctx, req.Email)
	if value == false {
		return userdto.SignupResponseDto{}, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindConflict).
			WithMessage(errmsg.ErrorMsg_UserDuplicated)
	}

	user := entity.User{
		ID:       0,
		Email:    req.Email,
		Name:     req.Name,
		Password: getMD5Hash(req.Password),
		Role:     entity.UserRole,
		Status:   true,
	}

	// create new user in storage
	createdUser, err := s.repo.Register(ctx, user)
	if err != nil {
		return userdto.SignupResponseDto{}, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsg_SomethingWentWrong)
	}

	// return created user
	return userdto.SignupResponseDto{User: userdto.UserInfoDto{
		ID:    		createdUser.ID,
		Email: 		createdUser.Email,
		Name:  		createdUser.Name,
	}}, nil
}
