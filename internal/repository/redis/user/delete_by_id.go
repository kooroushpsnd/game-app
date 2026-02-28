package redisuser

import (
	"context"
	"fmt"
	"goProject/internal/pkg/errmsg"
	"log"
)

func (c *UserCache) DeleteByID(ctx context.Context, userID uint) error {
	key := fmt.Sprintf(c.config.UserCacheKeyID, userID)
	if err := c.client.Del(ctx, key).Err(); err != nil {
		log.Println(errmsg.ErrorMsg_UserRedisDeleteIDError ,err)
		
		return err
	}

	return nil
}