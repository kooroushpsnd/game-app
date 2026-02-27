package postgresuser

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"goProject/internal/pkg/richerror"
	"log"

	"github.com/redis/go-redis/v9"
)

func (r *Repo) GetUserByID(ctx context.Context, userID uint) (entity.User, error) {
	const op = "postgres.getUserByID"

	if r.redis != nil && r.redis.Client() != nil {
		key := fmt.Sprintf(r.redis.Config().UserCacheKeyID, userID)

		val, err := r.redis.Client().Get(ctx, key).Bytes()
		if err == nil {
			var u entity.User
			if jsonErr := json.Unmarshal(val, &u); jsonErr == nil {
				return u, nil
			}
			_ = r.redis.Client().Del(ctx, key).Err()
		} else if err != redis.Nil {
			log.Println(op ,err)
		}
	}

	query := fmt.Sprintf("SELECT %s FROM users WHERE id = $1", UserColumns)

	row := r.db.QueryRowContext(ctx, query, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).
				WithErr(err).WithMessage(errmsg.ErrorMsg_UserNotFound).WithKind(richerror.KindNotFound)

		}

		return entity.User{}, richerror.New(op).
			WithErr(err).WithMessage(errmsg.ErrorMsg_CantScanQueryResult).WithKind(richerror.KindInvalid)
	}

	if r.redis != nil && r.redis.Client() != nil {
		emailKey := fmt.Sprintf(r.redis.Config().UserCacheKeyEmail, user.Email)
		idKey := fmt.Sprintf(r.redis.Config().UserCacheKeyID, user.ID)

		if b, mErr := json.Marshal(user); mErr == nil {
			_ = r.redis.Client().Set(ctx, emailKey, b, r.redis.Config().UserCacheTTL).Err()
			_ = r.redis.Client().Set(ctx, idKey, b, r.redis.Config().UserCacheTTL).Err()
		}
	}

	return user, nil
}
