package emailcodedto

import "goProject/internal/entity"

type UpdateEmailCodeRequestDto struct {
	Status   *entity.EmailCodeStatus
	Attempts *int
}