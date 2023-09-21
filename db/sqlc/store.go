package sqlc

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
	AddAddressTx(ctx context.Context, arg AddAddressesTxParams) (AddAddressesTxResult, error)
	// update new transaction functions
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
