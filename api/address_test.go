package api

import (
	"bytes"
	mockdb "ecom/db/mock"
	"ecom/db/sqlc"
	"ecom/db/util"
	"ecom/token"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateAddress(t *testing.T) {
	addrparam := randomAddressParam(t)
	testcases := []struct {
		name          string
		buildstubs    func(store *mockdb.MockStore)
		setupAuth     func(paesto *token.PasetoMaker, req *http.Request, t *testing.T)
		request       gin.H
		checkResponse func(recorder *httptest.ResponseRecorder, t *testing.T)
	}{
		{
			name: "OK",
			request: gin.H{
				"full_name":     addrparam.FullName,
				"country_code":  addrparam.CountryCode,
				"city":          addrparam.City,
				"street":        addrparam.Street,
				"landmark":      addrparam.Landmark,
				"mobile_number": fmt.Sprint(addrparam.MobileNumber),
			},
			buildstubs: func(store *mockdb.MockStore) {
				store.EXPECT().AddAddress(gomock.Any(), *addrparam).Times(1).Return(sqlc.Address{
					ID:           int32(util.RandomAmount(100)),
					CountryCode:  addrparam.CountryCode,
					City:         addrparam.City,
					Street:       addrparam.Street,
					FullName:     addrparam.FullName,
					Username:     addrparam.Username,
					Landmark:     addrparam.Landmark,
					MobileNumber: addrparam.MobileNumber,
				}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder, t *testing.T) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
			setupAuth: func(paesto *token.PasetoMaker, req *http.Request, t *testing.T) {
				token, paylaod, err := paesto.CreateToken(addrparam.Username, time.Minute)
				require.NoError(t, err)
				require.NotEmpty(t, token)
				require.NotEmpty(t, paylaod)
				req.Header.Set(authorizationHeaderKey, fmt.Sprint(authorizationTypeBearer, " ", token))
			},
		},
		{
			name:       "badRequest",
			buildstubs: func(store *mockdb.MockStore) {},
			setupAuth: func(paesto *token.PasetoMaker, req *http.Request, t *testing.T) {
				token, paylaod, err := paesto.CreateToken(addrparam.Username, time.Minute)
				require.NoError(t, err)
				require.NotEmpty(t, token)
				require.NotEmpty(t, paylaod)
				req.Header.Set(authorizationHeaderKey, fmt.Sprint(authorizationTypeBearer, " ", token))
			},
			request: gin.H{
				"full_name":     addrparam.FullName,
				"country_code":  addrparam.CountryCode,
				"city":          addrparam.City,
				"street":        addrparam.Street,
				"landmark":      addrparam.Landmark,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder, t *testing.T) {
				require.Equal(t,http.StatusBadRequest,recorder.Code)
			},
		},
	}
	for i := range testcases {
		tc := testcases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctr := gomock.NewController(t)
			store := mockdb.NewMockStore(ctr)
			server := newTestServer(t, store)
			url := "/address"
			tc.buildstubs(store)
			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tc.request)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
			require.NoError(t, err)
			tc.setupAuth(server.paseto, req, t)
			server.router.ServeHTTP(recorder, req)
			res, err := io.ReadAll(recorder.Body)
			require.NoError(t, err)
			fmt.Println(string(res))
			tc.checkResponse(recorder, t)

		})
	}
}

func randomAddressParam(t *testing.T) *sqlc.AddAddressParams {
	user, _ := randomUser(t)
	return &sqlc.AddAddressParams{
		Username:     user.Username,
		FullName:     user.FullName,
		CountryCode:  util.RandomFullName(3),
		City:         util.RandomFullName(10),
		Street:       util.RandomFullName(10),
		Landmark:     util.RandomFullName(10),
		MobileNumber: util.RandomMobileNumber(),
	}
}
