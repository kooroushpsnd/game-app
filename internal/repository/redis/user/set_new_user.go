package redisuser

import (
	"context"
	"encoding/json"
	"fmt"
	"goProject/internal/entity"
	"goProject/internal/pkg/errmsg"
	"log"
)

func (c *UserCache) Set(ctx context.Context, user entity.User) error {
	b, err := json.Marshal(user)
	if err != nil {
		log.Println(errmsg.ErrorMsg_UserRedisSetError ,err)
		return err
	}

	idKey := fmt.Sprintf(c.config.UserCacheKeyID, user.ID)
	emailKey := fmt.Sprintf(c.config.UserCacheKeyEmail, user.Email)

	pipe := c.client.Pipeline()
	pipe.Set(ctx, idKey, b, c.config.UserCacheTTL)
	pipe.Set(ctx, emailKey, b, c.config.UserCacheTTL)
	_, err = pipe.Exec(ctx)
	return err
}