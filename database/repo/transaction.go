package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"github.com/induzo/gocom/database/pginit/v2"
	"github.com/jackc/pgx/v5"

	"github.com/ivxivx/go-practices/database/domain"
	"github.com/ivxivx/go-practices/database/domain/transaction"
)

func (r *Repository) CreateTransaction(
	ctx context.Context,
	payload *transaction.Transaction,
) (*transaction.Transaction, error) {
	if payload.ID == uuid.Nil {
		payload.ID = transaction.NewID()
	}

	var record *transaction.Transaction

	if err := r.RunInTransaction(ctx, pgx.TxOptions{}, func(dbtx pgx.Tx) error {
		var errC error

		record, errC = createTransactionInTransaction(ctx, dbtx, payload)

		return errC
	}); err != nil {
		return nil, fmt.Errorf("CreateTransaction: %w", err)
	}

	return record, nil
}

func createTransactionInTransaction(
	ctx context.Context,
	txn pgx.Tx,
	payload *transaction.Transaction,
) (*transaction.Transaction, error) {
	query := `
		INSERT INTO transaction (
			id,
			internal_data,
			external_data
		)
		VALUES (
			@id,
			@internal_data,
			@external_data
		)
		RETURNING (
			JSON_BUILD_OBJECT(
				'id', id,
				'internal_data', internal_data,
				'external_data', external_data
			)
		)
	`

	args := pgx.NamedArgs{
		"id":            payload.ID,
		"internal_data": payload.InternalData,
		"external_data": payload.ExternalData,
	}

	row, err := txn.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("could not create transaction: %w", err)
	}

	var ent *transaction.Transaction

	ent, err = pgx.CollectExactlyOneRow(row, pginit.JSONRowToAddrOfStruct[transaction.Transaction])
	if err != nil {
		return nil, fmt.Errorf("could not create transaction: %w", err)
	}

	return ent, nil
}

func (r *Repository) GetTransaction(
	ctx context.Context,
	id uuid.UUID,
) (*transaction.Transaction, error) {
	query := `
		SELECT
			JSON_BUILD_OBJECT(
				'id', id,
				'internal_data', internal_data,
				'external_data', external_data
			)
		FROM transaction
		WHERE id = @id
	`

	args := pgx.NamedArgs{
		"id": id,
	}

	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("could not get record: %w", err)
	}

	var record *transaction.Transaction

	record, err = pgx.CollectExactlyOneRow(rows, pginit.JSONRowToAddrOfStruct[transaction.Transaction])
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, &domain.NotFoundError{ID: id.String()}
		default:
			return nil, fmt.Errorf("could not get record: %w", err)
		}
	}

	return record, nil
}
