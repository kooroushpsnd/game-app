package usercontroller

import (
	authservice "goProject/internal/service/auth"
	userservice "goProject/internal/service/user"
)

type Controller struct {
	authConfig authservice.Config
	authSvc    authservice.Service
	userSvc    userservice.Service
}

func New(authConfig authservice.Config, authSvc authservice.Service, userSvc userservice.Service) Controller {
	return Controller{
		authConfig: authConfig,
		authSvc:    authSvc,
		userSvc:    userSvc,
	}
}
