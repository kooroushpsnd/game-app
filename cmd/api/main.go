package main

import (
	"context"
	"fmt"
	"goProject/config"
	"goProject/controller/httpserver"
	applicatioDto "goProject/dto/application"
	postgres "goProject/repository/pg"
	postgresaction "goProject/repository/pg/pgaction"
	postgresitem "goProject/repository/pg/pgitem"
	postgrestransaction "goProject/repository/pg/pgtransaction"
	postgresuser "goProject/repository/pg/pguser"
	actionService "goProject/service/action"
	authService "goProject/service/auth"
	itemService "goProject/service/item"
	transactionService "goProject/service/transaction"
	userService "goProject/service/user"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.Load("config.yml")

	go func(){
		http.ListenAndServe(fmt.Sprintf(":%d", cfg.Application.Port), nil)
	}()

	services := setupServices(cfg)

	server := httpserver.New(cfg ,services)
	go func ()  {
		server.Serve()
	}()

	quit := make(chan os.Signal ,1)
	signal.Notify(quit ,os.Interrupt)
	<-quit
	fmt.Println("starting to shutdown gracefully")

	ctx := context.Background()
	ctxWithTimeout ,cancel := context.WithTimeout(ctx ,cfg.Application.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error" ,err)
	}

	time.Sleep(cfg.Application.GracefulShutdownTimeout)
	<-ctxWithTimeout.Done()
}

func setupServices(cfg config.Config) (setupServiceDto applicatioDto.SetupServiceDTO){
	authSvc := authService.New(cfg.Auth)

	postgresRepo := postgres.New(cfg.PostgreSQL)

	userRepo := postgresuser.New(postgresRepo)
	userSvc := userService.New(authSvc ,userRepo)

	transactionRepo := postgrestransaction.New(postgresRepo)
	transactionSvc := transactionService.New(transactionRepo)

	itemRepo := postgresitem.New(postgresRepo)
	itemSvc := itemService.New(itemRepo)

	actionRepo := postgresaction.New(postgresRepo)
	actionSvc := actionService.New(actionRepo)

	return applicatioDto.SetupServiceDTO{
		UserService:        userSvc,
		ItemService:        itemSvc,
		ActionService:      actionSvc,
		TransactionService: transactionSvc,
	}
}