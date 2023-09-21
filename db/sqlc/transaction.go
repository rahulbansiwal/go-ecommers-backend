package sqlc

import (
	"context"
	"fmt"
)

func (s *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	//defer tx.Rollback()
	//qtx := s.WithTx(tx)
	q := New(tx)
	err = fn(q)
	if err != nil {
		rberr := tx.Rollback()
		if rberr != nil {
			return fmt.Errorf("rollback error is %v , transaction error is %v", rberr, err)
		}
		return err
	}
	return tx.Commit()
}
