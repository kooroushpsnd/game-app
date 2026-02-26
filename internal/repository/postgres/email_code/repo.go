package postgresemailcode

import (
	"goProject/internal/config"
	"goProject/internal/repository/postgres"
)

type Repo struct {
	config config.Config
	db postgres.DBTX
}

func New(db postgres.DBTX, config config.Config) *Repo{
	return &Repo{db: db, config: config}
}