package postgresemailcode

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
)

func (r *Repo) CheckEmailCodeReadyToResend(ctx context.Context, email string) (bool, error) {
	const op = "postgresemailcode.CheckEmailCodeReadyToResend"
	
	codeExist ,err := r.GetLatestEmailCode(ctx ,email)
	if(err != nil){
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, err
	}


	log.Println(r.config.Application.EmailCodeResendMinute)
	resendAfter := r.config.Application.EmailCodeResendMinute
	if time.Since(codeExist.CreatedAt) >= resendAfter {
		return true, nil
	}

	return false , nil
}