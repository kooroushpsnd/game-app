package redisuser

import (
	"context"
	"fmt"
	"goProject/internal/pkg/errmsg"
	"log"
)

func (c *UserCache) DeleteByEmail(ctx context.Context, email string) error {
	key := fmt.Sprintf(c.config.UserCacheKeyEmail, email)
	if err := c.client.Del(ctx, key).Err(); err != nil {
		log.Println(errmsg.ErrorMsg_UserRedisDeleteEmailError ,err)
		
		return err
	}
	
	return c.client.Del(ctx, key).Err()
}