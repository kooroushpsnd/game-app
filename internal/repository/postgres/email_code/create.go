package postgresemailcode

import (
	"context"
	emailcodedto "goProject/internal/dto/email_code"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
	"time"
)

func (r *Repo) CreateEmailCode(ctx context.Context, req emailcodedto.CreateEmailCodeDto) error {
	const op = "postgresemailcode.create"

	exp := time.Now().Add(r.config.Application.EmailCodeExpirationDateMinute)
	
	const q = `
		INSERT INTO email_codes (email, hash_code, status, attempts, expiration_date, user_id ,created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(
		ctx,
		q,
		req.Email,
		req.HashCode,
		entity.EmailCodeStatusActive,
		0,
		exp,
		req.UserID,
		time.Now(),
	)
	if err != nil {
		return richerror.New(op).WithErr(err).
			WithKind(richerror.KindConflict).
			WithMessage(errmsg.ErrorMsg_CantScanQueryResult)
	}

	return nil
}
