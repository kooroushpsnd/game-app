package postgresuser

import (
	"context"
	"database/sql"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (r *Repo) IsEmailUnique(ctx context.Context, email string) (bool, error) {
	const op = "postgres.IsEmailUnique"

	const q = "select * from users where email = $1"

	row := r.db.QueryRowContext(ctx, q, email)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsg_CantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return false, nil
}
