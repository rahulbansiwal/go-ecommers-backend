package api

import (
	"bytes"
	"database/sql"
	mockdb "ecom/db/mock"
	"ecom/db/sqlc"
	"ecom/db/util"
	"ecom/token"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type EqCreatUserParamsMatcher struct {
	arg      sqlc.CreateUserParams
	password string
}

func (e EqCreatUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func (e EqCreatUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(sqlc.CreateUserParams)
	if !ok {
		return false
	}
	err := util.CheckPassword(e.password, e.arg.HashedPassword)
	if err != nil {
		return false
	}
	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func EqCreatUserParams(arg sqlc.CreateUserParams, password string) gomock.Matcher {
	return EqCreatUserParamsMatcher{arg, password}
}

func TestCreateUser(t *testing.T) {
	user, password := randomUser(t)
	fmt.Printf("actual random user: %+v", user)
	testcases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		setupAuth     func(t *testing.T, req *http.Request, paseto *token.PasetoMaker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":        user.Username,
				"hashed_password": password,
				"full_name":       user.FullName,
				"mobile_number":   fmt.Sprint(user.MobileNumber.Int64),
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := sqlc.CreateUserParams{
					Username:       user.Username,
					FullName:       user.FullName,
					HashedPassword: user.HashedPassword,
					MobileNumber:   user.MobileNumber,
				}
				store.EXPECT().CreateUser(gomock.Any(), EqCreatUserParams(arg, password)).
					Times(1).Return(user, nil)
			},
			setupAuth: func(t *testing.T, req *http.Request, paseto *token.PasetoMaker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				fmt.Print(recorder.Flushed)
				requireBodyMatcherFunc(t, recorder.Body, user)

			},
		},
		{
			name: "InternalServerError",
			body: gin.H{
				"username":        user.Username,
				"hashed_password": password,
				"full_name":       user.FullName,
				"mobile_number":   fmt.Sprint(user.MobileNumber.Int64),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateUser(gomock.Any(),gomock.Any()).Times(1).
				Return(sqlc.User{},sql.ErrConnDone)
			},
			setupAuth: func(t *testing.T, req *http.Request, paseto *token.PasetoMaker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t,http.StatusInternalServerError,recorder.Code)
			},
		},
		{
			name:"BadRequest",
			body: gin.H{},
			buildStubs: func(store *mockdb.MockStore) {},
			setupAuth: func(t *testing.T, req *http.Request, paseto *token.PasetoMaker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t,http.StatusBadRequest,recorder.Code)
			},
		},
		{
			name:"BadRequest2",
			body: gin.H{
				"username":        user.Username,
				"hashed_password": password,
				"mobile_number":   fmt.Sprint(user.MobileNumber.Int64),
			},
			buildStubs: func(store *mockdb.MockStore) {},
			setupAuth: func(t *testing.T, req *http.Request, paseto *token.PasetoMaker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t,http.StatusBadRequest,recorder.Code)
			},
		},
		{
			name:"BadRequest2",
			body: gin.H{
				"username":        user.Username,
				"hashed_password": password,
				"mobile_number":   "abc122423",
			},
			buildStubs: func(store *mockdb.MockStore) {},
			setupAuth: func(t *testing.T, req *http.Request, paseto *token.PasetoMaker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t,http.StatusBadRequest,recorder.Code)
			},
		},
		{
			name:"BadRequest2",
			body: gin.H{
				"username":        user.Username,
				"hashed_password": "asc",
				"mobile_number":   "1234",
			},
			buildStubs: func(store *mockdb.MockStore) {},
			setupAuth: func(t *testing.T, req *http.Request, paseto *token.PasetoMaker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t,http.StatusBadRequest,recorder.Code)
			},
		},

	}
	for i := range testcases {
		tc := testcases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			defer ctr.Finish()
			store := mockdb.NewMockStore(ctr)
			tc.buildStubs(store)
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()
			url := "/user"
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			tc.setupAuth(t, req, server.paseto)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomUser(t *testing.T) (sqlc.User, string) {
	password := util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	user := sqlc.User{
		Username:       util.RandomUsername(),
		FullName:       util.RandomFullName(6),
		HashedPassword: hashedPassword,
		MobileNumber: sql.NullInt64{
			Int64: util.RandomMobileNumber(),
			Valid: true,
		},
		PasswordChangedAt: time.Now(),
		IsEmailVerified:   false,
		CreatedAt:         time.Now(),
	}
	return user, password
}

func requireBodyMatcherFunc(t *testing.T, body *bytes.Buffer, user sqlc.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)
	var gotUser CreateUserRequest
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Empty(t, gotUser.HashedPassword)
	require.Equal(t, fmt.Sprint(user.MobileNumber.Int64), gotUser.MobileNumber)

}
