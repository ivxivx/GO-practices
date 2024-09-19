package repo

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ivxivx/go-practices/database/domain/transaction"
	"github.com/ivxivx/go-practices/util"
)

func Test_Transaction(t *testing.T) {
	t.Parallel()

	repo, err := createDatabase(t.Name())
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	t.Cleanup(func() {
		repo.pool.Close()
	})

	ctx := context.Background()

	testCases := []struct {
		name          string
		request       *transaction.Transaction
		expectedError error
	}{
		{
			name: "no external data",
			request: &transaction.Transaction{
				ID:           transaction.NewID(),
				InternalData: json.RawMessage(`{"key": "value"}`),
			},
		},
		{
			name: "with external data",
			request: &transaction.Transaction{
				ID:           transaction.NewID(),
				InternalData: json.RawMessage(`{"key": "value"}`),
				ExternalData: util.ToPointer(json.RawMessage(`{"key": "value"}`)),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, actualError := repo.CreateTransaction(ctx, tc.request)

			if tc.expectedError != nil {
				if actualError == nil {
					t.Errorf("expected error %v, got nil", tc.expectedError)
				} else {
					if actualError.Error() != tc.expectedError.Error() {
						t.Errorf("expected error %v, got %v", tc.expectedError, actualError)
					}
				}
			} else {
				if actualError != nil {
					t.Errorf("unexpected error: %v", actualError)
				}

				txn, err := repo.GetTransaction(ctx, tc.request.ID)
				if err != nil {
					t.Fatalf("failed to get transaction: %v", err)
				}

				if !reflect.DeepEqual(tc.request, txn) {
					t.Errorf("expected %v, got %v", tc.request, txn)
				}
			}
		})
	}
}
