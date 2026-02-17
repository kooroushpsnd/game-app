package postgresemailcode

import (
	"context"
	emailcodedto "goProject/internal/dto/email_code"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (r *Repo) CreateEmailCode(ctx context.Context, req emailcodedto.CreateEmailCodeDto) error {
	const op = "postgresemailcode.create"

	const q = `
		INSERT INTO email_codes (email, hash_code, status, attempts, expiration_date, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(
		ctx,
		q,
		req.Email,
		req.HashCode,
		true,
		req.Attempts,
		req.ExpirationDate,
		req.UserID,
	)
	if err != nil {
		return richerror.New(op).WithErr(err).
			WithKind(richerror.KindConflict).
			WithMessage(errmsg.ErrorMsg_CantScanQueryResult)
	}

	return nil
}
