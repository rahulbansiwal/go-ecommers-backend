package sqlc

import (
	"context"
	"database/sql"
	"ecom/db/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)

	getusr, err := testQueries.GetUser(context.Background(), user.Username)
	assert.Empty(t, err)
	assert.NotEmpty(t, getusr)
	assert.Equal(t, getusr.Username, user.Username)
	assert.Equal(t, getusr.FullName, user.FullName)
	assert.Equal(t, getusr.HashedPassword, user.HashedPassword)
	assert.Equal(t, getusr.CreatedAt, user.CreatedAt)
	assert.Equal(t, getusr.PasswordChangedAt, user.PasswordChangedAt)
	assert.Equal(t, getusr.MobileNumber, user.MobileNumber)
	assert.Equal(t, getusr.IsEmailVerified, user.IsEmailVerified)
}

func TestUpdateUser(t *testing.T) {
	user := CreateRandomUser(t)
	newdetails := struct {
		NewUsername    string
		FullName       string
		HashedPassword string
		MobileNumber   int64
	}{
		NewUsername:    util.RandomUsername(),
		FullName:       util.RandomFullName(6),
		HashedPassword: util.RandomFullName(6),
		MobileNumber:   util.RandomMobileNumber(),
	}
	newuser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		OldUser: user.Username,
		Username: sql.NullString{
			String: newdetails.NewUsername,
			Valid:  true,
		},
		FullName: sql.NullString{
			String: newdetails.FullName,
			Valid:  true,
		},
		HashedPassword: sql.NullString{
			String: newdetails.HashedPassword,
			Valid:  true,
		},
		MobileNumber: sql.NullInt64{
			Int64: newdetails.MobileNumber,
			Valid: true,
		},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, newuser)
	assert.Equal(t, newdetails.NewUsername, newuser.Username)
	assert.Equal(t, newdetails.FullName, newuser.FullName)
	assert.Equal(t, newdetails.HashedPassword, newuser.HashedPassword)
	assert.Equal(t, newdetails.MobileNumber, newuser.MobileNumber.Int64)

}

func CreateRandomUser(t *testing.T) *User {
	args := CreateUserParams{
		Username:       util.RandomUsername(),
		FullName:       util.RandomFullName(6),
		HashedPassword: util.RandomFullName(6),
		MobileNumber: sql.NullInt64{
			Int64: util.RandomMobileNumber(),
			Valid: true,
		},
	}
	user, err := testQueries.CreateUser(context.Background(), args)
	//fmt.Printf("%+v", user)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
	assert.Equal(t, user.Username, args.Username)
	assert.Equal(t, user.FullName, args.FullName)
	assert.Equal(t, user.MobileNumber, args.MobileNumber)
	assert.Equal(t, user.HashedPassword, args.HashedPassword)
	assert.WithinDuration(t, user.CreatedAt, time.Now(), time.Minute)
	assert.WithinDuration(t, user.PasswordChangedAt, time.Now(), time.Minute)
	return &user
}
