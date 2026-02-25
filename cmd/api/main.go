package main

import (
	"context"
	"fmt"
	mailer "goProject/internal/adapter/email"
	"goProject/internal/config"
	"goProject/internal/controller/httpserver"
	applicatioDto "goProject/internal/dto/application"
	"goProject/internal/repository/postgres"

	// postgresaction "goProject/internal/repository/postgres/action"
	// postgresitem "goProject/internal/repository/postgres/item"
	// postgrestransaction "goProject/internal/repository/postgres/transaction"
	postgresemailcode "goProject/internal/repository/postgres/email_code"
	postgresuser "goProject/internal/repository/postgres/user"

	// actionservice "goProject/internal/service/action"
	authservice "goProject/internal/service/auth"
	emailservice "goProject/internal/service/email"

	// itemservice "goProject/internal/service/item"
	// transactionservice "goProject/internal/service/transaction"
	userservice "goProject/internal/service/user"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.Load("config.yml")

	go func() {
		http.ListenAndServe(fmt.Sprintf(":%d", cfg.Application.Port), nil)
	}()

	services := setupServices(cfg)

	server := httpserver.New(cfg, services)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("starting to shutdown gracefully")

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, cfg.Application.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}

	time.Sleep(cfg.Application.GracefulShutdownTimeout)
	<-ctxWithTimeout.Done()
}

func setupServices(cfg config.Config) (setupServiceDto applicatioDto.SetupServiceDTO) {
	authSvc := authservice.New(cfg.Auth)

	postgresRepo, err := postgres.New()
	if err != nil {
		fmt.Println("postgres error", err)
		return
	}

	sqlDB := postgresRepo.Conn()

	userRepo := postgresuser.New(sqlDB)
	userSvc := userservice.New(authSvc, userRepo)

	emailRepo := postgresemailcode.New(sqlDB ,cfg)
	mailerAdaptor := mailer.NewSMTPEmailAdapter()
	emailSvc := emailservice.New(userSvc ,emailRepo ,mailerAdaptor ,cfg.Application)

	// transactionRepo := postgrestransaction.New(sqlDB)
	// transactionSvc := transactionservice.New(transactionRepo)

	// itemRepo := postgresitem.New(sqlDB)
	// itemSvc := itemservice.New(itemRepo)

	// actionRepo := postgresaction.New(sqlDB)
	// actionSvc := actionservice.New(actionRepo)

	return applicatioDto.SetupServiceDTO{
		UserService:  userSvc,
		AuthService:  authSvc,
		EmailService: emailSvc,
		// ItemService:        itemSvc,
		// ActionService:      actionSvc,
		// TransactionService: transactionSvc,
	}
}
