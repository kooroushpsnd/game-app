package postgresuser

import (
	"context"
	"fmt"
	userdto "goProject/internal/dto/user"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
	"strings"
)

func (r *Repo) GetAllUsers(ctx context.Context, req userdto.GetAllRequestUserDto) ([]entity.User, error) {
	const op = "postgresuser.GetAllUsers"
	const userColumns = `id, email, name, password, role, status, created_at, updated_at`

	conds := make([]string, 0, 3)
	args := make([]any, 0, 5)
	i := 1

	if req.Email != nil {
		conds = append(conds, fmt.Sprintf("email ILIKE $%d", i))
		args = append(args, "%"+strings.TrimSpace(*req.Email)+"%")
		i++
	}
	if req.Status != nil {
		conds = append(conds, fmt.Sprintf("status = $%d", i))
		args = append(args, *req.Status)
		i++
	}
	if req.Role != nil {
		conds = append(conds, fmt.Sprintf("role = $%d", i))
		args = append(args, strings.TrimSpace(*req.Role))
		i++
	}

	where := ""
	if len(conds) > 0 {
		where = " WHERE " + strings.Join(conds, " AND ")
	}

	// ----- Count query -----
	countQ := "SELECT COUNT(*) FROM users" + where
	var total int
	if err := r.db.QueryRowContext(ctx, countQ, args...).Scan(&total); err != nil {
		return []entity.User{}, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsg_SomethingWentWrong)
	}
	
	limit := req.PaginationDto.LimitOr()
	offset := req.PaginationDto.OffsetOr()

	selectQ := fmt.Sprintf(`
		SELECT %s
		FROM users
		%s
		ORDER BY id DESC
		LIMIT $%d OFFSET $%d
	`,userColumns , where, i, i+1)

	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, selectQ, args...)
	if err != nil {
		return []entity.User{}, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsg_SomethingWentWrong)
	}
	defer rows.Close()

	users := make([]entity.User, 0, limit)
	for rows.Next() {
		u, err := scanUser(rows)
		if err != nil {
			return []entity.User{}, richerror.New(op).
				WithErr(err).
				WithKind(richerror.KindUnexpected).
				WithMessage(errmsg.ErrorMsg_SomethingWentWrong)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return []entity.User{}, richerror.New(op).
			WithErr(err).
        	WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsg_SomethingWentWrong)
	}

	return users ,nil
}