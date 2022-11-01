package db

import (
	"database/sql"
)

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

// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
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
