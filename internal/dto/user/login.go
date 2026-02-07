package userdto

type LoginRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponseDto struct {
	User   UserInfoDto `json:"user"`
	Tokens TokensDto   `json:"tokens"`
}
