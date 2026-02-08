package userdto

type UpdateRequestAdminDto struct {
	Email *string `json:"email"`
	Name *string `json:"name"`
	Role *string `json:"role" validate:"oneof=user admin"`
	Status *bool `json:"status"`
}

type UpdateResponseAdminDto struct {
	User   UserInfoDto `json:"user"`
}