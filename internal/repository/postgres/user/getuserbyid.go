package postgresuser

import (
	"context"
	"database/sql"
	"fmt"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (r *Repo) GetUserByID(ctx context.Context, userID uint) (entity.User, error) {
	const op = "postgres.getUserByID"

	userCache ,exists ,err := r.userCache.GetByID(ctx, userID)
	if err == nil && exists {
		return userCache, nil
	}

	query := fmt.Sprintf("SELECT %s FROM users WHERE id = $1", UserColumns)

	row := r.db.QueryRowContext(ctx, query, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).
				WithErr(err).WithMessage(errmsg.ErrorMsg_UserNotFound).WithKind(richerror.KindNotFound)

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
