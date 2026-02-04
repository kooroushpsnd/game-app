package postgresuser

import (
	"context"
	"database/sql"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	const op = "postgres.getUserByEmail"

	row := r.db.QueryRowContext(ctx, "select * from users where email = ?", email)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{} ,richerror.New(op).
			WithErr(err).WithMessage(errmsg.ErrorMsgUserNotFound).WithKind(richerror.KindNotFound)
			
		}

		return entity.User{} ,richerror.New(op).
		WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindInvalid)
	}

	return user ,nil
}