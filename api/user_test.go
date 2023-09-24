package api

import (
	"database/sql"
	mockdb "ecom/db/mock"
	"ecom/db/sqlc"
	"ecom/db/util"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	user := randomUser()
	require.NotEmpty(t, user)

	testcases := []struct {
		name          string
		user          sqlc.User
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		
	}
}

func randomUser() sqlc.User {
	return sqlc.User{
		Username:       util.RandomUsername(),
		FullName:       util.RandomFullName(6),
		HashedPassword: util.RandomString(6),
		MobileNumber: sql.NullInt64{
			Int64: util.RandomMobileNumber(),
			Valid: true,
		},
		CreatedAt:         time.Now(),
		IsEmailVerified:   false,
		PasswordChangedAt: time.Now(),
	}
}
