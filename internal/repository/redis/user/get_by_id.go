package redisuser

import (
	"context"
	"encoding/json"
	"fmt"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"log"

	"github.com/redis/go-redis/v9"
)

func (c *UserCache) GetByID(ctx context.Context, userID uint) (entity.User, bool, error) {
	key := fmt.Sprintf(c.config.UserCacheKeyID, userID)

	b, err := c.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		log.Println(errmsg.ErrorMsg_UserRedisGetIDError ,err)
		return entity.User{}, false, nil
	}
	if err != nil {
		log.Println(errmsg.ErrorMsg_UserRedisGetIDError ,err)
		return entity.User{}, false, err
	}

	var u entity.User
	if err := json.Unmarshal(b, &u); err != nil {
		log.Println(errmsg.ErrorMsg_UserRedisGetIDError ,err)

		_ = c.client.Del(ctx, key).Err()
		return entity.User{}, false, nil
	}
	return u, true, nil
}