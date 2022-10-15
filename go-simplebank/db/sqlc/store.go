package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides  all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		rberr := tx.Rollback()
		if rberr != nil {
			return fmt.Errorf("tx err: %v,rb err: %v", err, rberr)
		}
		return err
	}
	return tx.Commit()
}

// TranserTx performs a money transfer one account to another
// func (store *Store) TranserTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
// 	return nil
// }

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Entry    `json:"to_account_"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"To_entry"`
}
