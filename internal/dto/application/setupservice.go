package applicatioDto

import (
	// actionservice "goProject/internal/service/action"
	// itemservice "goProject/internal/service/item"
	// transactionservice "goProject/internal/service/transaction"
	"goProject/internal/adapter/redis"
	authservice "goProject/internal/service/auth"
	emailservice "goProject/internal/service/email"
	userservice "goProject/internal/service/user"
)

type SetupServiceDTO struct {
	UserService  *userservice.Service
	AuthService  *authservice.Service
	EmailService *emailservice.Service
	RedisAdaptor *redis.Adapter
	// ItemService        *itemservice.Service
	// ActionService      *actionservice.Service
	// TransactionService *transactionservice.Service
}
