package applicatioDto

import (
	// actionservice "goProject/internal/service/action"
	// itemservice "goProject/internal/service/item"
	// transactionservice "goProject/internal/service/transaction"
	userservice "goProject/internal/service/user"
)

type SetupServiceDTO struct {
	UserService        *userservice.Service
	// ItemService        *itemservice.Service
	// ActionService      *actionservice.Service
	// TransactionService *transactionservice.Service
}
