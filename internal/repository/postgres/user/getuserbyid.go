package postgresuser

import (
	"context"
	"database/sql"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (r *Repo) GetUserByID(ctx context.Context, userID uint) (entity.User, error) {
	const op = "postgres.getUserByID"

	const q = "select * from users where id = ?"

	row := r.db.QueryRowContext(ctx ,q ,userID)
	user ,err := scanUser(row)
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