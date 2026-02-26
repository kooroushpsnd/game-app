package postgresemailcode

import (
	"context"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (r *Repo) IncrementEmailCodeAttempts(ctx context.Context, email string) (entity.EmailCode, error) {
	const op = "postgresemailcode.IncrementEmailCodeAttempts"
	emailCode, err := r.GetLatestEmailCode(ctx, email)
	if err != nil {
		return entity.EmailCode{}, err
	}
	emailCode.Attempts++

	const q = `
		UPDATE email_codes
		SET attempts = $1
		WHERE id = $2
		RETURNING id, email, hash_code, status, attempts, expiration_date, user_id, created_at;
	`

	rows := r.db.QueryRowContext(ctx, q, emailCode.Attempts, emailCode.ID)
	emailCode, err = scanEmailCode(rows)
	if err != nil {
		return entity.EmailCode{}, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindConflict).
			WithMessage(errmsg.ErrorMsg_CantScanQueryResult)
	}
	return emailCode, nil
}