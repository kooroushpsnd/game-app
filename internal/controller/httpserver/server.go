package httpserver

import (
	"fmt"
	"goProject/internal/config"
	"goProject/internal/controller/httpserver/usercontroller"
	applicatioDto "goProject/internal/dto/application"
	"goProject/internal/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config         config.Config
	userController usercontroller.Controller
	Router         *echo.Echo
}

func New(config config.Config, services applicatioDto.SetupServiceDTO) Server {
	e := echo.New()
	e.Validator = validator.New()

	return Server{
		config: config,
		Router: e,
	}
}

func (s Server) Serve() {
	s.Router.Use(middleware.RequestID())
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

	s.Router.GET("/health-check", s.healthCheck)

	address := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("http server error (start)", err)
	}
}
