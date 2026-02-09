package userdto

import "goProject/internal/dto"

type GetAllRequestUserDto struct {
	Email  *string `query:"email"    validate:"omitempty,email"`
	Status *bool   `query:"status"`
	Role   *string `query:"role"    validate:"omitempty,oneof=user admin"`
	dto.PaginationDto
}

type GetAllResponseUserDto struct {
	Users []UserInfoDto `json:"users"`
	dto.PaginationResponseDto
}
