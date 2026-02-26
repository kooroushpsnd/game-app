package userdto

type UserInfoDto struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	EmailVerify bool `json:"email_verify"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Status bool   `json:"status"`
}
