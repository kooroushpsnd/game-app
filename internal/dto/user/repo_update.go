package userdto

type UserUpdatePatch struct {
	Email       *string
	EmailVerify *bool
	Name        *string
	Role        *string
	Status      *bool
}
