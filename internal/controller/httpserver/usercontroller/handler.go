package usercontroller

import (
	authservice "goProject/internal/service/auth"
	// emailservice "goProject/internal/service/email"
	userservice "goProject/internal/service/user"
)

type Controller struct {
	authConfig authservice.Config
	authSvc    *authservice.Service
	userSvc    *userservice.Service
	// emailSvc   *emailservice.Service
}

func New(authConfig authservice.Config ,authSvc *authservice.Service, userSvc *userservice.Service) *Controller {
	return &Controller{
		authConfig: authConfig,
		authSvc:    authSvc,
		userSvc:    userSvc,
		// emailSvc:   emailSvc,
	}
}
