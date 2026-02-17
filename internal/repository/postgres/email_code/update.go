package postgresemailcode

import (
	"context"
	"fmt"
	emailcodedto "goProject/internal/dto/email_code"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
	"strings"
)

func (r *Repo) UpdateEmailCode(ctx context.Context, email string, req emailcodedto.UpdateEmailCodeRequestDto) (entity.EmailCode, error) {
	const op = "postgresemailcode.UpdateEmailCode"
	codeExist, err := r.GetLatestEmailCode(ctx, email)
	if err != nil {
		return entity.EmailCode{}, err
	}

	sets := make([]string, 0, 2)
	args := make([]any, 0, 3)
	idx := 1

	if req.Attempts != nil {
		sets = append(sets, fmt.Sprintf("attempts = $%d", idx))
		args = append(args, *req.Attempts)
		idx++
	}
	if req.Status != nil {
		sets = append(sets, fmt.Sprintf("status = $%d", idx))
		args = append(args, *req.Status)
		idx++
	}

	if len(sets) == 0 {
		return codeExist, nil
	}

	args = append(args, codeExist.ID)

	q := fmt.Sprintf(`
        UPDATE email_codes
        SET %s
        WHERE id = $%d
        RETURNING id, email, hash_code, status, attempts, expiration_date, user_id
    `, strings.Join(sets, ", "), idx)

	row := r.db.QueryRowContext(ctx, q, args...)
	emailCode, err := scanEmailCode(row)
	if err != nil {
		return entity.EmailCode{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsg_CantScanQueryResult).
			WithKind(richerror.KindInvalid)
	}

	return emailCode, nil
}
