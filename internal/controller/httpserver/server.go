package httpserver

import (
	"fmt"
	"goProject/internal/config"
	applicatioDto "goProject/internal/dto/application"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config config.Config
	Router *echo.Echo
}

func New(config config.Config ,services applicatioDto.SetupServiceDTO) Server {
	return Server{
		config: config,
		Router: echo.New(),
	}
}

func (s Server) Serve(){
	s.Router.Use(middleware.RequestID())
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

	s.Router.GET("/health-check" ,s.healthCheck)

	address := fmt.Sprintf(":%d" ,s.config.HTTPServer.Port)
	fmt.Printf("start echo server on %s\n" ,address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("http server error (start)" ,err)
	}
}