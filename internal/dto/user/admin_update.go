package userdto

type UpdateRequestAdminDto struct {
	Email       *string `json:"email" validate:"omitempty,email"`
	EmailVerify *bool   `json:"email_verify" validate:"omitempty,boolean"`
	Name        *string `json:"name"`
	Role        *string `json:"role" validate:"omitempty,oneof=user admin"`
	Status      *bool   `json:"status"`
}
