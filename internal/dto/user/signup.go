package userdto

type SignupRequestDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SignupResponseDto struct {
	User UserInfoDto `json:"user"`
}
