package applicatioDto

import (
	actionService "goProject/service/action"
	itemService "goProject/service/item"
	transactionService "goProject/service/transaction"
	userService "goProject/service/user"
)

type SetupServiceDTO struct {
	UserService        *userService.Service
	ItemService        *itemService.Service
	ActionService      *actionService.Service
	TransactionService *transactionService.Service
}
