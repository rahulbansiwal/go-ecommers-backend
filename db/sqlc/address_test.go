package sqlc

import (
	"context"
	"ecom/db/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAddress(t *testing.T) {
	user := CreateRandomUser(t)
	addressdetails := util.RandomAddressDetails()
	address,err := testQueries.AddAddress(context.Background(),AddAddressParams{
		Username: user.Username,
		FullName: user.FullName,
		MobileNumber: user.MobileNumber.Int64,
		Street: addressdetails.Street,
		CountryCode: addressdetails.CountryCode,
		City: addressdetails.City,
		Landmark: addressdetails.Landmark,
	})
	assert.NoError(t,err)
	assert.NotEmpty(t,address)
	assert.Equal(t,address.City,addressdetails.City)
	assert.Equal(t,address.CountryCode,addressdetails.CountryCode)
	assert.Equal(t,address.Street,addressdetails.Street)
	assert.Equal(t,address.Landmark,addressdetails.Landmark)
	assert.Equal(t,address.Username,user.Username)
	assert.Equal(t,address.FullName,user.FullName)
	assert.Equal(t,address.MobileNumber,user.MobileNumber.Int64)
}
