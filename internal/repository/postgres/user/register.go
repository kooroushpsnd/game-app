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
		INSERT INTO users (email, name, password, role, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, email, name, password, role, status, created_at, updated_at
	`

	var roleStr string

	row := r.db.QueryRowContext(
		ctx,
		q,
		u.Email,
		u.Name,
		u.Password,
		u.Role.String(),
		u.Status,
	)
	user ,err := scanUser(row)
	if err != nil {
		return entity.User{} ,richerror.New(op).
			WithErr(err).WithMessage(errmsg.ErrorMsgUserCreation)
	}

	user.Role = entity.MapToRoleEntity(roleStr)
	return user, nil
}
