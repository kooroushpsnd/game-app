package userdto

type UserInfoDto struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Status bool   `json:"status"`
}
