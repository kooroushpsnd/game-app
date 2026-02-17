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

func (r *Repo) UpdateUser(ctx context.Context ,userID uint ,req userdto.UserUpdatePatch) (entity.User, error) {
	const op = "postgres.UpdateUser"
	userExist ,err := r.GetUserByID(ctx ,userID)
	if err != nil {
		return entity.User{} ,err
	}

	sets := make([]string, 0, 4)
    args := make([]any, 0, 5)
    idx := 1

    if req.Email != nil {
        sets = append(sets, fmt.Sprintf("email = $%d", idx))
        args = append(args, *req.Email)
        idx++
    }
    if req.Name != nil {
        sets = append(sets, fmt.Sprintf("name = $%d", idx))
        args = append(args, *req.Name)
        idx++
    }
    if req.Role != nil {
        sets = append(sets, fmt.Sprintf("role = $%d", idx))
        args = append(args, *req.Role)
        idx++
    }
    if req.Status != nil {
        sets = append(sets, fmt.Sprintf("status = $%d", idx))
        args = append(args, *req.Status)
        idx++
    }

    if len(sets) == 0 {
        return userExist ,nil
    }

    sets = append(sets, "updated_at = NOW()")
    args = append(args, userID)

    q := fmt.Sprintf(`
        UPDATE users
        SET %s
        WHERE id = $%d
        RETURNING id, email, name, password, role, status, created_at, updated_at
    `, strings.Join(sets, ", "), idx)

    row := r.db.QueryRowContext(ctx, q, args...)
	user ,err := scanUser(row)
	if err != nil {
		return entity.User{}, richerror.New(op).
			WithErr(err).WithMessage(errmsg.ErrorMsg_CantScanQueryResult).WithKind(richerror.KindInvalid)
	}

    return user, err
}