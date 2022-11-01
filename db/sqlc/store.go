package db

import (
	"database/sql"
)

type Store interface {
	Querier
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		rbErr := tx.Rollback()
// 		if rbErr != nil {
// 			return fmt.Errorf("TxErr: %v, RBErr: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }
