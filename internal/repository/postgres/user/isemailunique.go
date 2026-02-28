package postgresuser

import (
	"context"
	"database/sql"
	"fmt"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (r *Repo) IsEmailUnique(ctx context.Context, email string) (bool, error) {
	const op = "postgres.IsEmailUnique"
	query := fmt.Sprintf("SELECT %s FROM users WHERE email = $1", UserColumns)

	_ ,exists ,err := r.userCache.GetByEmail(ctx, email)
	if err == nil && exists {
		return false, nil
	}

	row := r.db.QueryRowContext(ctx, query, email)
	_, err = scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsg_CantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return false, nil
}
