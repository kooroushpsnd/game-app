package postgresuser

import (
	"context"
	"database/sql"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
	"log"
)

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	const op = "postgres.getUserByEmail"

	row := r.db.QueryRowContext(ctx, "select * from users where email = $1", email)
	user, err := scanUser(row)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).
				WithErr(err).
				WithMessage(errmsg.ErrorMsg_UserNotFound).
				WithKind(richerror.KindNotFound)

		}

		return entity.User{}, richerror.New(op).
			WithErr(err).WithMessage(errmsg.ErrorMsg_CantScanQueryResult).WithKind(richerror.KindInvalid)
	}

	return user, nil
}
