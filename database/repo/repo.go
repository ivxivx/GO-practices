package repo

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/induzo/gocom/database/pginit/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ivxivx/go-practices/domain/transaction"
)

type Repo interface {
	transaction.Repo
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(ctx context.Context, connString string) (*Repository, error) {
	pgi, err := pginit.New(
		connString,
		pginit.WithDecimalType(),
		pginit.WithGoogleUUIDType(),
		pginit.WithLogger(slog.Default(), "pgx"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize PGInit: %w", err)
	}

	pool, err := pgi.ConnPool(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate connection pool: %w", err)
	}

	return &Repository{
		pool: pool,
	}, nil
}

func (repo *Repository) GetShutdownFuncs() map[string]func(ctx context.Context) error {
	return map[string]func(ctx context.Context) error{
		"postgres": func(_ context.Context) error {
			repo.pool.Close()

			return nil
		},
	}
}

func (repo *Repository) RunInTransaction(
	ctx context.Context,
	txOptions pgx.TxOptions,
	fn func(tx pgx.Tx) error,
) (err error) {
	tx, err := repo.pool.BeginTx(ctx, txOptions)
	if err != nil {
		return fmt.Errorf("failed to begin database transaction: %w", err)
	}

	defer func() {
		if err == nil {
			errC := tx.Commit(ctx)
			if errC != nil {
				err = fmt.Errorf("failed to commit database transaction: %w", errC)
			}
		} else {
			errR := tx.Rollback(ctx)
			if errR != nil {
				err = fmt.Errorf("failed to rollback database transaction: %w", errR)
			}
		}
	}()

	err = fn(tx)

	return err
}
