package api

import (
	"ecom/token"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func addAuthorizationToken(
	t *testing.T,
	request *http.Request,
	paesto *token.PasetoMaker,
	authorizationType string,
	username string,
	duration time.Duration,
) {
	token, payload, err := paesto.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	testcases := []struct {
		name          string
		setupAuth     func(t *testing.T, req *http.Request, paesto *token.PasetoMaker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker) {
				addAuthorizationToken(t, req, paesto, authorizationTypeBearer, "user", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NoAuth",
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker) {

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "unsupportedAuthType",
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker) {
				addAuthorizationToken(t, req, paesto, "unsupported", "user", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "invalidauthorizationtype",
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker) {
				addAuthorizationToken(t, req, paesto, "", "user", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "expiredToken",
			setupAuth: func(t *testing.T, req *http.Request, paesto *token.PasetoMaker) {
				addAuthorizationToken(t, req, paesto, authorizationTypeBearer, "user", -time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}
	for i := range testcases {
		tc := testcases[i]
		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t, nil)
			authPath := "/auth"
			server.router.GET(authPath,
				AuthMiddleware(server.paseto),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK,gin.H{})
				},
			)
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet,authPath,nil)
			tc.setupAuth(t,request,server.paseto)
			server.router.ServeHTTP(recorder,request)
			tc.checkResponse(t,recorder)
			
		})
	}

}
