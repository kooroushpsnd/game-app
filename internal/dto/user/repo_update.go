package userdto

import "goProject/internal/entity"

type UserUpdatePatch struct {
	Email  *string
	Name   *string
	Role   *entity.Role
	Status *bool
}
