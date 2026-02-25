package config

import (
	authservice "goProject/internal/service/auth"
	"time"
)

type Application struct {
	EmailCodeResendMinute         time.Duration `koanf:"email_code_resend_minute"`
	EmailCodeExpirationDateMinute time.Duration `koanf:"email_code_expiration_date_minute"`
	EmailCodeAttemptsMax          int           `koanf:"email_code_attempts_max"`
	GracefulShutdownTimeout       time.Duration `koanf:"graceful_shutdown_timeout"`
	Port                          int           `koanf:"port"`
}

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct {
	Application Application        `koanf:"application"`
	HTTPServer  HTTPServer         `koanf:"http_server"`
	Auth        authservice.Config `koanf:"auth"`
}
