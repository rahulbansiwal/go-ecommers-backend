package api

import (
	"ecom/token"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(paseto *token.PasetoMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {

			ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("authorization header is not provided")))
			return
		}
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 2 {

			ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid authorization header value")))
			return
		}
		authorizationType := strings.ToLower(fields[0])

		if authorizationType != authorizationTypeBearer {

			ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid authorization type")))
			return
		}
		accessToken := fields[1]
		payload, err := paseto.VerifyToken(accessToken)
		if err != nil {

			ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("unable to verify token")))
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		fmt.Print("passing to next handler")
		ctx.Next()
	}
}
