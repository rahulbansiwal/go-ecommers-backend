package sqlc

import (
	"context"
	"ecom/db/util"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddAddressesTx(t *testing.T) {
	store := NewStore(testdb)
	fmt.Printf("from inside the test %+v", testdb)
	randAddress := util.RandomAddressDetails()
	user := CreateRandomUser(t)
	args := AddAddressesTxParams{
		AddAddressParams: AddAddressParams{
			Username:     user.Username,
			FullName:     user.FullName,
			CountryCode:  randAddress.CountryCode,
			City:         randAddress.City,
			Street:       randAddress.Street,
			Landmark:     randAddress.Landmark,
			MobileNumber: user.MobileNumber.Int64,
		},
	}
	addr, err := store.AddAddressTx(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, addr)
	require.Equal(t, args.FullName, addr.address.FullName)
	require.Equal(t, args.City, addr.address.City)
	require.Equal(t, args.CountryCode, addr.address.CountryCode)
	require.Equal(t, args.Landmark, addr.address.Landmark)
	require.Equal(t, args.Username, addr.address.Username)
	require.Equal(t, args.Street, addr.address.Street)
	require.Equal(t, args.MobileNumber, addr.address.MobileNumber)
}
