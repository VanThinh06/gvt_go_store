package db

import (
	"context"
	"database/sql"
	"fmt"

	"gopkg.in/guregu/null.v4"
)

// import "database/sql"
// Định nghĩa
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

type TxOptions struct {
	Isolation sql.IsolationLevel
	ReadOnly  bool
}

func (store *Store) exectTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()

}

type PopulateProductTxResult struct {
	Category    Category    `json:"category"`
	Name        string      `json:"name"`
	Price       null.Int    `json:"price"`
	Image       null.String `json:"image"`
	ListImage   []string    `json:"list_image"`
	Description null.String `json:"description"`
	Sold        null.Int    `json:"sold"`
	Status      null.Int    `json:"status"`
	Sale        null.Int    `json:"sale"`
}

func (store *Store) TransferTx(ctx context.Context) (PopulateProductTxResult, error) {
	result := PopulateProductTxResult{}

	err := store.exectTx(ctx, func(q *Queries) error {
		// var err error
		// var i Category
		// result.Category, err = q.GetAllCategory(ctx)

		return nil
	})
	return result, err
}
