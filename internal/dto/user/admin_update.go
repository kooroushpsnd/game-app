package userdto

type UpdateRequestAdminDto struct {
	Email *string `json:"email" validate:"omitempty,email"`
	Name *string `json:"name"`
	Role *string `json:"role" validate:"omitempty,oneof=user admin"`
	Status *bool `json:"status"`
}