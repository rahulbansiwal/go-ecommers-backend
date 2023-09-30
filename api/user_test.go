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
				first :=store.EXPECT().CreateUser(gomock.Any(), EqCreatUserParams(arg, password)).
					Times(1).Return(user, nil)
					store.EXPECT().CreateCart(gomock.Any(),arg.Username).After(first).Times(1).Return(sqlc.Cart{},nil)
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
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(1).
					Return(sqlc.User{}, sql.ErrConnDone)
			},
			setupAuth: func(t *testing.T, req *http.Request, paseto *token.PasetoMaker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:       "BadRequest",
			body:       gin.H{},
			buildStubs: func(store *mockdb.MockStore) {},
			setupAuth:  func(t *testing.T, req *http.Request, paseto *token.PasetoMaker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "BadRequest2",
			body: gin.H{
				"username":        user.Username,
				"hashed_password": password,
				"mobile_number":   fmt.Sprint(user.MobileNumber.Int64),
			},
			buildStubs: func(store *mockdb.MockStore) {},
			setupAuth:  func(t *testing.T, req *http.Request, paseto *token.PasetoMaker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "BadRequest2",
			body: gin.H{
				"username":        user.Username,
				"hashed_password": password,
				"mobile_number":   "abc122423",
			},
			buildStubs: func(store *mockdb.MockStore) {},
			setupAuth:  func(t *testing.T, req *http.Request, paseto *token.PasetoMaker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "BadRequest2",
			body: gin.H{
				"username":        user.Username,
				"hashed_password": "asc",
				"mobile_number":   "1234",
			},
			buildStubs: func(store *mockdb.MockStore) {},
			setupAuth:  func(t *testing.T, req *http.Request, paseto *token.PasetoMaker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
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

func TestGetUser(t *testing.T) {
	user, _ := randomUser(t)

	testcases := []struct {
		name          string
		param         string
		buildStubs    func(store *mockdb.MockStore)
		setupAuth     func(t *testing.T, req *http.Request, paesto *token.PasetoMaker, user sqlc.User)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:  "OK",
			param: user.Username,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), user.Username).Times(1).Return(user, nil)
			},
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker, user sqlc.User) {
				token, payload, err := paesto.CreateToken(user.Username, time.Minute)
				require.NoError(t, err)
				require.NotEmpty(t, token)
				require.NotEmpty(t, payload)
				req.Header.Set(authorizationHeaderKey, fmt.Sprint(authorizationTypeBearer, " ", token))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, recorder.Code, http.StatusOK)
			},
		},
		{
			name:  "BadRequest",
			param: user.Username,
			buildStubs: func(store *mockdb.MockStore) {
				//store.EXPECT().GetUser(gomock.Any(), gomock.Any()).Times(1).Return(user, nil)
			},
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker, user sqlc.User) {
				req.Header.Set(authorizationHeaderKey, fmt.Sprint(authorizationTypeBearer))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, recorder.Code, http.StatusBadRequest)
			},
		},
		{
			name:  "NotFound",
			param: user.Username,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Any()).Times(1).Return(sqlc.User{}, sql.ErrNoRows)
			},
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker, user sqlc.User) {
				token, payload, err := paesto.CreateToken(user.Username, time.Minute)
				require.NoError(t, err)
				require.NotEmpty(t, token)
				require.NotEmpty(t, payload)
				req.Header.Set(authorizationHeaderKey, fmt.Sprint(authorizationTypeBearer, " ", token))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, recorder.Code, http.StatusNotFound)
			},
		},
		{
			name:  "InternalServerError",
			param: user.Username,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Any()).Times(1).Return(sqlc.User{}, sql.ErrConnDone)
			},
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker, user sqlc.User) {
				token, payload, err := paesto.CreateToken(user.Username, time.Minute)
				require.NoError(t, err)
				require.NotEmpty(t, token)
				require.NotEmpty(t, payload)
				req.Header.Set(authorizationHeaderKey, fmt.Sprint(authorizationTypeBearer, " ", token))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, recorder.Code, http.StatusInternalServerError)
			},
		},
		{
			name:  "Unauthorized",
			param: "abcd",
			buildStubs: func(store *mockdb.MockStore) {
				//store.EXPECT().GetUser(gomock.Any(), user.Username).Times(1).Return(user, nil)
			},
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker, user sqlc.User) {
				token, payload, err := paesto.CreateToken("user", time.Minute)
				require.NoError(t, err)
				require.NotEmpty(t, token)
				require.NotEmpty(t, payload)
				req.Header.Set(authorizationHeaderKey, fmt.Sprint(authorizationTypeBearer, " ", token))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, recorder.Code, http.StatusUnauthorized)
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
			url := fmt.Sprint("/user/", tc.param)
			fmt.Print(url)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			tc.setupAuth(t, request, server.paseto, user)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}

}

func TestUpdateUser(t *testing.T) {
	user, password := randomUser(t)
	testcases := []struct {
		name          string
		body          gin.H
		param         string
		queryparam    string
		buildStubs    func(store *mockdb.MockStore)
		setupAuth     func(t *testing.T, req *http.Request, paesto *token.PasetoMaker, username string)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"full_name":       util.RandomFullName(6),
				"hashed_password": util.RandomString(6),
				"mobile_number":   fmt.Sprint(util.RandomMobileNumber()),
			},
			param:      user.Username,
			queryparam: "?fields=name&fields=password&fields=msisdn",
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker, username string) {
				token, _, err := paesto.CreateToken(username, time.Minute)
				require.NoError(t, err)
				req.Header.Set(authorizationHeaderKey, fmt.Sprint(authorizationTypeBearer, " ", token))
			},
			buildStubs: func(store *mockdb.MockStore) {
				firstCall := store.EXPECT().GetUser(gomock.Any(), user.Username).Times(1).Return(user, nil)
				store.EXPECT().UpdateUser(gomock.Any(),gomock.Any()).After(firstCall).Times(1).Return(sqlc.User{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, recorder.Code, http.StatusOK)
			},
		},
		{
			name: "SamePassword",
			body: gin.H{
				"full_name":       util.RandomFullName(6),
				"hashed_password": password,
				"mobile_number":   fmt.Sprint(util.RandomMobileNumber()),
			},
			param:      user.Username,
			queryparam: "?fields=name&fields=password&fields=msisdn",
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker, username string) {
				token, _, err := paesto.CreateToken(username, time.Minute)
				require.NoError(t, err)
				req.Header.Set(authorizationHeaderKey, fmt.Sprint(authorizationTypeBearer, " ", token))
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), user.Username).Times(1).Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, recorder.Code, http.StatusBadRequest)
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
			url := "/user/"
			recorder := httptest.NewRecorder()
			req_url := fmt.Sprint(url, tc.param, tc.queryparam)
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, req_url, bytes.NewReader(data))
			require.NoError(t, err)
			tc.setupAuth(t, req, server.paseto, user.Username)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)

		})
	}

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
