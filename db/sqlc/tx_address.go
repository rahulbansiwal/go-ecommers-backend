package sqlc

import (
	"context"
	"fmt"
)

type AddAddressesTxParams struct {
	AddAddressParams
}

type AddAddressesTxResult struct {
	address Address
}

func (s *SQLStore) AddAddressTx(ctx context.Context, arg AddAddressesTxParams) (AddAddressesTxResult, error) {
	var result AddAddressesTxResult
	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		result.address, err = q.AddAddress(ctx, arg.AddAddressParams)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	})
	return result, err
}
