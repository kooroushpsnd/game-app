package userdto

type UserSignupRequestDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignupResponseDto struct {
	User UserInfoDto `json:"user"`
}
