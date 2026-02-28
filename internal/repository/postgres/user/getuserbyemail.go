package postgresuser

import (
	"context"
	"database/sql"
	"fmt"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	const op = "postgres.getUserByEmail"
	query := fmt.Sprintf("SELECT %s FROM users WHERE email = $1", UserColumns)

	userCache ,exists ,err := r.userCache.GetByEmail(ctx, email)
	if err == nil && exists {
		return userCache, nil
	}

	row := r.db.QueryRowContext(ctx, query, email)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).
				WithErr(err).
				WithMessage(errmsg.ErrorMsg_UserNotFound).
				WithKind(richerror.KindNotFound)

		}

		return entity.User{}, richerror.New(op).
			WithErr(err).WithMessage(errmsg.ErrorMsg_CantScanQueryResult).WithKind(richerror.KindInvalid)
	}

	err = r.userCache.Set(ctx, user)
	if err != nil {
		return user, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsg_UserRedisSetError).
			WithKind(richerror.KindUnexpected)
	}

	return user, nil
}
