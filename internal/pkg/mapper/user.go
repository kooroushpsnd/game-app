package mapper

import (
	userdto "goProject/internal/dto/user"
	"goProject/internal/entity"
)

func ToUserInfoDto(u entity.User) userdto.UserInfoDto {
	return userdto.UserInfoDto{
		ID:     u.ID,
		Email:  u.Email,
		Name:   u.Name,
		Role:   u.Role.String(),
		Status: u.Status,
	}
}

func ToUserInfoDtos(users []entity.User) []userdto.UserInfoDto {
	out := make([]userdto.UserInfoDto, 0, len(users))
	for _, u := range users {
		out = append(out, ToUserInfoDto(u))
	}
	return out
}

func AdminDtoToPatch(req userdto.UpdateRequestAdminDto) userdto.UserUpdatePatch {
	patch := userdto.UserUpdatePatch{
		Email:  req.Email,
		Name:   req.Name,
		Status: req.Status,
		Role:   req.Role,
	}

	return patch
}

func UserDtoToPatch(req userdto.UpdateRequestUserDto) userdto.UserUpdatePatch {
	return userdto.UserUpdatePatch{
		Name: req.Name,
	}
}
