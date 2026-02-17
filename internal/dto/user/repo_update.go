package userdto

type UserUpdatePatch struct {
	Email  *string
	Name   *string
	Role   *string
	Status *bool
}
