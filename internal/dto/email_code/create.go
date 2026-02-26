package emailcodedto

type CreateEmailCodeDto struct {
	Email          string
	HashCode       string
	UserID         uint
}
