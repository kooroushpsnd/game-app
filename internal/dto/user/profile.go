package userdto

type userGetProfileRequestDto struct {
	UserID uint 
}

type userGetProfileResponseDto struct {
	User UserInfoDto `json:"user"`
}