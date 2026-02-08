package userdto

type UpdateRequestUserDto struct {
	Name *string `json:"name"`
}

type UpdateResponseUserDto struct {
	User   UserInfoDto `json:"user"`
}