package userdto

type UserLoginRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserLoginResponseDto struct {
	User   UserInfoDto `json:"user"`
	Tokens TokensDto   `json:"tokens"`
}
