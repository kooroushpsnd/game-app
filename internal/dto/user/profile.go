package userdto

type GetProfileRequestDto struct {
	UserID uint 
}

type GetProfileResponseDto struct {
	User UserInfoDto `json:"user"`
}