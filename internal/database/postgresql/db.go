package postgres

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Port     int    `koanf:"port"`
	Host     string `koanf:"host"`
	DBName   string `koanf:"db_name"`
	SSLMode  string `koanf:"ssl_mode"`
}

type PostgreSQLDB struct {
	config Config
	db     *gorm.DB
}

func (p *PostgreSQLDB) Conn() *gorm.DB {
	return p.db
}

func New(config Config) *PostgreSQLDB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host,
		config.Username,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode,
	)

	gormDB ,err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic((fmt.Errorf("can't open postgres db: %v" ,err)))
	}

	sqlDB ,err := gormDB.DB()
	if err != nil {
		panic((fmt.Errorf("can't get sql db from gorm db: %v" ,err)))
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	return &PostgreSQLDB{config: config ,db: gormDB}
}