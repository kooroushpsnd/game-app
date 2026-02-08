package middleware

import (
	cfg "goProject/internal/config"
	authservice "goProject/internal/service/auth"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo-jwt/v4"
)

func Auth(service *authservice.Service ,config authservice.Config) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
			ContextKey: cfg.AuthMiddlewareContextKey,
			SigningKey: []byte(config.SignKey),
			SigningMethod: config.SigningMethod,
			ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
				claims, err := service.ParseToken(auth)
				if err != nil {
					return nil, err
				}

				return claims, nil
			},
		},
	)
}