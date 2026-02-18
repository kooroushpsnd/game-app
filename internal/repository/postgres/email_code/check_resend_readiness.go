package postgresemailcode

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

func (r *Repo) CheckEmailCodeReadyToResend(ctx context.Context, email string) (bool, error) {
	codeExist ,err := r.GetLatestEmailCode(ctx ,email)
	if(err != nil){
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, err
	}

	if time.Since(codeExist.CreatedAt) >= (time.Minute * r.config.Application.EmailCodeExpirationDateMinute) {
		return true, nil
	}
	return true ,nil
}