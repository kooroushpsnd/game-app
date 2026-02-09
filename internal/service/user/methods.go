package userservice

import (
	"context"
	userdto "goProject/internal/dto/user"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/mapper"
	"goProject/internal/pkg/richerror"
)

func (s *Service) Login(ctx context.Context, req userdto.LoginRequestDto) (userdto.LoginResponseDto, error) {
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
		User: mapper.ToUserInfoDto(user),
		Tokens: userdto.TokensDto{
			AccessToken:  accessToken,
			RefreshToken: refreshToken},
	}, nil
}

func (s *Service) Register(ctx context.Context, req userdto.SignupRequestDto) (userdto.SignupResponseDto, error) {
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
	return userdto.SignupResponseDto{User: mapper.ToUserInfoDto(createdUser)}, nil
}

func (s *Service) GetProfile(ctx context.Context ,req userdto.GetProfileRequestDto) (userdto.GetProfileResponseDto ,error){
	const op = "userservice.Profile"

	user ,err := s.repo.GetUserByID(ctx ,req.UserID)
	if err != nil {
		return userdto.GetProfileResponseDto{} ,richerror.New(op).
		WithErr(err).
		WithKind(richerror.KindNotFound).
		WithMessage(errmsg.ErrorMsg_UserNotFound)
	}

	return userdto.GetProfileResponseDto{User: mapper.ToUserInfoDto(user)} ,nil
}

func (s *Service) GetAllUsers(ctx context.Context ,req userdto.GetAllRequestUserDto) (userdto.GetAllResponseUserDto ,error){
	const op = "userservice.GetAllUsers"

	result, err := s.repo.GetAllUsers(ctx, req)
	if err != nil {
		return userdto.GetAllResponseUserDto{}, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsg_SomethingWentWrong)
	}

	usersDto := make([]userdto.UserInfoDto, 0, len(result))
	for _, u := range result {
		usersDto = append(usersDto, mapper.ToUserInfoDto(u))
	}

	return userdto.GetAllResponseUserDto{
		Users:                 usersDto,
		PaginationResponseDto: mapper.ToPaginationDto(req.PaginationDto, len(usersDto)),
	}, nil
}

func (s *Service) Update(ctx context.Context ,userID int ,req userdto.UpdateRequestAdminDto) (userdto.UpdateResponseAdminDto ,error){
	
}