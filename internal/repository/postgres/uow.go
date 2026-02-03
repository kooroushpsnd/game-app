package postgres

import (
	"context"
	"database/sql"
)

type UnitOfWork struct {
	db *sql.DB
}

func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

func (u *UnitOfWork) WithTx(
	ctx context.Context,
	opts *sql.TxOptions,
	fn func(tx *sql.Tx) error,
) error {
	tx, err := u.db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	// اگر fn خطا داد یا زودتر برگشت، rollback انجام می‌شه
	defer func() { _ = tx.Rollback() }()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}
