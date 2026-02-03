package config

import (
	"goProject/internal/repository/postgres"
	authservice "goProject/internal/service/auth"
	"time"
)

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
	Port                    int           `koanf:"port"`
}

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct {
	Application Application        `koanf:"application"`
	HTTPServer  HTTPServer         `koanf:"http_server"`
	PostgreSQL  postgres.Config    `koanf:"postgresql"`
	Auth        authservice.Config `koanf:"auth"`
}
