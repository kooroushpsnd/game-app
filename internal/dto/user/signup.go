package userdto

type SignupRequestDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponseDto struct {
	User UserInfoDto `json:"user"`
}
