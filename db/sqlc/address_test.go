package sqlc

import (
	"context"
	"ecom/db/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateRandomAddress(t *testing.T) *Address {
	user := CreateRandomUser(t)
	addressdetails := util.RandomAddressDetails()
	address, err := testQueries.AddAddress(context.Background(), AddAddressParams{
		Username:     user.Username,
		FullName:     user.FullName,
		MobileNumber: user.MobileNumber.Int64,
		Street:       addressdetails.Street,
		CountryCode:  addressdetails.CountryCode,
		City:         addressdetails.City,
		Landmark:     addressdetails.Landmark,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, address)
	assert.Equal(t, address.City, addressdetails.City)
	assert.Equal(t, address.CountryCode, addressdetails.CountryCode)
	assert.Equal(t, address.Street, addressdetails.Street)
	assert.Equal(t, address.Landmark, addressdetails.Landmark)
	assert.Equal(t, address.Username, user.Username)
	assert.Equal(t, address.FullName, user.FullName)
	assert.Equal(t, address.MobileNumber, user.MobileNumber.Int64)
	return &address
}

func TestDeleteAddress(t *testing.T){
	a := CreateRandomAddress(t)
	deladdr, err := testQueries.DeleteAddress(context.Background(), a.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, deladdr)
	assert.Equal(t, deladdr.Street, a.Street)
	assert.Equal(t, deladdr.City, a.City)
	assert.Equal(t, deladdr.CountryCode, a.CountryCode)
	assert.Equal(t, deladdr.Landmark, a.Landmark)
	assert.Equal(t, deladdr.Username, a.Username)
	assert.Equal(t, deladdr.MobileNumber, a.MobileNumber)
	assert.Equal(t, deladdr.FullName, a.FullName)
	assert.Equal(t, deladdr.ID, a.ID)
	getaddr, err := testQueries.GetAddresses(context.Background(), a.Username)
	assert.NoError(t, err)
	assert.Len(t, getaddr, 0)
}

func TestAddAndGetAddress(t *testing.T) {
	addressdetails := CreateRandomAddress(t)
	addr, err := testQueries.GetAddresses(context.Background(), addressdetails.Username)
	assert.NoError(t, err)
	assert.NotEmpty(t, addr)
	assert.Len(t, addr, 1)
	a := addr[0]
	assert.Equal(t, addressdetails.Street, a.Street)
	assert.Equal(t, addressdetails.City, a.City)
	assert.Equal(t, addressdetails.CountryCode, a.CountryCode)
	assert.Equal(t, addressdetails.Landmark, a.Landmark)
	assert.Equal(t, addressdetails.Username, a.Username)
	assert.Equal(t, addressdetails.MobileNumber, a.MobileNumber)
	assert.Equal(t, addressdetails.FullName, a.FullName)


}
