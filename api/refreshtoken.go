package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiredAt time.Time `json:"access_token_expired_at"`
}

func (s *Server) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	refreshPayload, err := s.paseto.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	session, err := s.store.GetSessionFromId(ctx, refreshPayload.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	if session.IsBlocked {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("session blocked")))
		return
	}
	if session.Username != refreshPayload.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("incorrect session user")))
		return
	}
	if time.Now().After(session.ExpiredAt) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("token expired")))
		return
	}
	accessToken, accessPayload, err := s.paseto.CreateToken(session.Username, s.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiredAt: accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, response)
}
