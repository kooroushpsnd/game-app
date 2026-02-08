package postgresuser

import (
	"context"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
)

func (r *Repo) Register(ctx context.Context, u entity.User) (entity.User, error) {
	const op = "postgres.Register"

	const q = `
		INSERT INTO users (email, name, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, email, name, password, role, status, created_at, updated_at
	`

	row := r.db.QueryRowContext(
		ctx,
		q,
		u.Email,
		u.Name,
		u.Password,
		u.Role.String(),
	)
	user, err := scanUser(row)
	if err != nil {
		return entity.User{}, richerror.New(op).
			WithErr(err).WithMessage(errmsg.ErrorMsg_UserCreation)
	}

	return user, nil
}
