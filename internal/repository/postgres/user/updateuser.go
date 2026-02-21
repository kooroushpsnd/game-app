package postgresuser

import (
	"context"
	"fmt"
	userdto "goProject/internal/dto/user"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
	"log"
	"strings"
)

func (r *Repo) UpdateUser(ctx context.Context ,userID uint ,req userdto.UserUpdatePatch) (entity.User, error) {
	const op = "postgres.UpdateUser"
    log.Println("first update")
	userExist ,err := r.GetUserByID(ctx ,userID)
	if err != nil {
        log.Println(err)
		return entity.User{} ,err
	}

	sets := make([]string, 0, 5)
    args := make([]any, 0, 6)
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
    if req.EmailVerify != nil {
        sets = append(sets, fmt.Sprintf("email_verify = $%d", idx))
        args = append(args, *req.EmailVerify)
        idx++
    }

    if len(sets) == 0 {
        log.Println("zero" ,sets)
        return userExist ,nil
    }

    sets = append(sets, "updated_at = NOW()")
    args = append(args, userID)

    q := fmt.Sprintf(`
        UPDATE users
        SET %s
        WHERE id = $%d
        RETURNING id, email, name, password, role, status, created_at, updated_at, email_verify
    `, strings.Join(sets, ", "), idx)

    row := r.db.QueryRowContext(ctx, q, args...)
	user ,err := scanUser(row)
    log.Println(user)
	if err != nil {
        log.Println(err)
		return entity.User{}, richerror.New(op).
			WithErr(err).WithMessage(errmsg.ErrorMsg_CantScanQueryResult).WithKind(richerror.KindInvalid)
	}

    return user, err
}