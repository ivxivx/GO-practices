package transaction

import (
	"context"

	"github.com/goccy/go-json"

	"github.com/gofrs/uuid/v5"
)

type Transaction struct {
	ID           uuid.UUID        `json:"id"`
	InternalData json.RawMessage  `json:"internal_data"`
	ExternalData *json.RawMessage `json:"external_data"`
}

func NewID() uuid.UUID {
	return uuid.Must(uuid.NewV7())
}

type Repo interface {
	CreateTransaction(ctx context.Context, payload *Transaction) (*Transaction, error)
	GetTransaction(ctx context.Context, id uuid.UUID) (*Transaction, error)
}
