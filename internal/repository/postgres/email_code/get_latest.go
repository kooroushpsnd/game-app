package postgresemailcode

import (
	"context"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (r *Repo) GetLatestEmailCode(ctx context.Context, email string) (entity.EmailCode, error) {
	const op = "postgresemailcode.GetLeastEmailCode"

	const q = `SELECT * FROM email_codes where email = $1 ORDER BY id DESC LIMIT 1`

	row := r.db.QueryRowContext(ctx, q, email)
	emailCode, err := scanEmailCode(row)
	if err != nil {
		return entity.EmailCode{}, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindConflict).
			WithMessage(errmsg.ErrorMsg_CantScanQueryResult)
	}

	return emailCode, nil
}
