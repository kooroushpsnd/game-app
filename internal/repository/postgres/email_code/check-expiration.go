package postgresemailcode

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

func (r *Repo) CheckEmailCodeExpiration(ctx context.Context, email string) (bool, error) {
	codeExist ,err := r.GetLatestEmailCode(ctx ,email)
	if(err != nil){
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	exp := time.Minute * time.Duration(r.config.Application.EmailCodeExpirationDateMinute)

	return time.Now().After(codeExist.CreatedAt.Add(exp)), nil
}