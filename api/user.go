package api

import (
	"database/sql"
	"ecom/db/sqlc"
	"ecom/db/util"
	"ecom/token"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CreateUserRequest struct {
	Username       string `json:"username" binding:"required,email"`
	HashedPassword string `json:"hashed_password" binding:"required,min=6"`
	FullName       string `json:"full_name" binding:"required,min=6"`
	MobileNumber   string `json:"mobile_number" binding:"numeric"`
}

type UpdateUserUsername struct {
	Username string `json:"username" binding:"required,email"`
}
type UpdateUserFullName struct {
	FullName string `json:"full_name" binding:"required,min=6"`
}
type UpdateUserPassword struct {
	HashedPassword string `json:"hashed_password" binding:"required,min=6"`
}
type UpdateUserMobileNumber struct {
	MobileNumber string `json:"mobile_number" binding:"numeric"`
}

type LoginUserReqeust struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	SessionId             string             `json:"session_id"`
	AccessToken           string             `json:"access_token"`
	AccessTokenExpiredAt  time.Time          `json:"access_token_expired_at"`
	RefreshToken          string             `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time          `json:"refresh_token_expired_at"`
	User                  CreateUserResponse `json:"user"`
}

type CreateUserResponse struct {
	Username     string    `json:"username"`
	FullName     string    `json:"full_name"`
	MobileNumber string    `json:"mobile_number"`
	CreatedAt    time.Time `json:"created_at"`
}

func userResponse(user sqlc.User) CreateUserResponse {
	return CreateUserResponse{
		FullName:     user.FullName,
		Username:     user.Username,
		MobileNumber: fmt.Sprint(user.MobileNumber.Int64),
		CreatedAt:    user.CreatedAt,
	}
}

func (s *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var msisdn uint64
	if len(req.MobileNumber) > 0 {
		var err error
		msisdn, err = strconv.ParseUint(req.MobileNumber, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid mobile_number passed %w", err)))
			return
		}
	}
	hashedPassword, err := util.HashPassword(req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	args := sqlc.CreateUserParams{
		Username:       req.Username,
		FullName:       req.FullName,
		HashedPassword: hashedPassword,
		MobileNumber: sql.NullInt64{
			Int64: int64(msisdn),
			Valid: true,
		},
	}
	if len(req.MobileNumber) == 0 {
		args.MobileNumber.Int64 = 0
		args.MobileNumber.Valid = false
	}
	user, err := s.store.CreateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("error while creating user %w", err)))
		return
	}
	_,err = s.store.CreateCart(ctx,req.Username)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	res := userResponse(user)
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetUser(ctx *gin.Context) {
	username := ctx.Param("id")
	fmt.Printf("%+v", ctx)
	authpayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if username != authpayload.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("unauthorized")))
		return
	}
	user, err := s.store.GetUser(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("user not found for the %v", username)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}

	res := userResponse(user)
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) UpdateUser(ctx *gin.Context) {
	olduser := ctx.Param("username")
	if len(olduser) == 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid user")))
		return
	}
	authpayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	
	if authpayload.Username != olduser{
		ctx.JSON(http.StatusUnauthorized,errorResponse(fmt.Errorf("unauthorized for the current user")))
	}

	fields, ok := ctx.GetQueryArray("fields")
	if !ok {
		fields = append(fields, "username")
	}
	user, err := s.store.GetUser(ctx, olduser)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotAcceptable, errorResponse(fmt.Errorf("invalid user")))
			return
		}
		ctx.JSON(http.StatusNotAcceptable, errorResponse(err))
		return
	}
	var req sqlc.UpdateUserParams
	req.OldUser = olduser
	for _, v := range fields {
		switch v {
		case "username":
			var r UpdateUserUsername
			if err := ctx.ShouldBindBodyWith(&r, binding.JSON); err != nil {
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
			req.Username.String = r.Username
			req.Username.Valid = true
		case "password":
			var r UpdateUserPassword
			if err := ctx.ShouldBindBodyWith(&r, binding.JSON); err != nil {
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
			err := util.CheckPassword(r.HashedPassword, user.HashedPassword)
			fmt.Print(err)
			if err == nil {
				ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("password is same as before")))
				return
			}
			hp, err := util.HashPassword(r.HashedPassword)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			}
			req.HashedPassword.String = hp
			req.HashedPassword.Valid = true
		case "name":
			var r UpdateUserFullName
			if err := ctx.ShouldBindBodyWith(&r, binding.JSON); err != nil {
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
			req.FullName.String = r.FullName
			req.FullName.Valid = true
		case "msisdn":
			var r UpdateUserMobileNumber
			if err := ctx.ShouldBindBodyWith(&r, binding.JSON); err != nil {
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
			msisdn, err := strconv.ParseUint(r.MobileNumber, 10, 64)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid mobile_number passed %w", err)))
				return
			}
			req.MobileNumber.Int64 = int64(msisdn)
			req.MobileNumber.Valid = true
		}
	}
	fmt.Printf("%+v", req)
	user, err = s.store.UpdateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, fmt.Errorf("error while updating error %w", err))
		return
	}
	res := userResponse(user)
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserReqeust
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := s.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("user not found")))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("incorrect password")))
		return
	}

	token, payload, err := s.paseto.CreateToken(user.Username, s.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	refreshToken, refreshPayload, err := s.paseto.CreateToken(user.Username, s.config.RefreshTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	_, err = s.store.CreateSession(ctx, sqlc.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     req.Username,
		RefreshToken: refreshToken,
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	res := LoginUserResponse{
		SessionId:             payload.ID.String(),
		AccessToken:           token,
		AccessTokenExpiredAt:  payload.ExpiredAt,
		User:                  userResponse(user),
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: refreshPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, res)
}
