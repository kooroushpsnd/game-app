package emailcodedto

type SendEmailCodeDto struct {
	Email string `json:"email" validate:"required,email"`
}