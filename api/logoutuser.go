package api

import (
	"ecom/token"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutUserResponse struct {
	SessionId string `json:"session_id"`
	Message   string `json:"message"`
}

type LogoutUserFromAllDevicesResponse struct {
	Message string `json:"message"`
}

func (s *Server) LogoutUser(ctx *gin.Context) {
	fmt.Print("inside the next handler")
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload ==nil{
		return
	}
	session, err := s.store.DeleteSessionById(ctx, authPayload.ID)
	fmt.Print(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	req := LogoutUserResponse{
		SessionId: session.ID.String(),
		Message:   "logout from current device",
	}
	ctx.JSON(http.StatusOK, req)
}

func (s *Server) LogoutUserFromAllDevice(ctx *gin.Context) {
	authpaylaod := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	_, err := s.store.DeleteSessionByUsername(ctx, authpaylaod.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	req := LogoutUserFromAllDevicesResponse{
		Message: "logout from all devices",
	}
	ctx.JSON(http.StatusOK, req)
}
