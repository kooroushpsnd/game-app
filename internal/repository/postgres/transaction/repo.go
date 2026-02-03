package postgrestransaction

import "goProject/internal/repository/postgres"

type Repo struct {
	db postgres.DBTX
}

func New(db postgres.DBTX) *Repo{
	return &Repo{db: db}
}