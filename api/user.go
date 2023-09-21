package api

import (
	"database/sql"
	"ecom/db/sqlc"
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
	msisdn, err := strconv.ParseUint(req.MobileNumber, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid mobile_number passed %w", err)))
		return
	}
	args := sqlc.CreateUserParams{
		Username:       req.Username,
		FullName:       req.FullName,
		HashedPassword: req.HashedPassword,
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
	res := userResponse(user)
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetUser(ctx *gin.Context) {
	username := ctx.Param("id")
	fmt.Printf("%+v", ctx)
	user, err := s.store.GetUser(ctx, username)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("user not found for the %v", username)))
		return
	}
	res := userResponse(user)
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) UpdateUser(ctx *gin.Context) {
	olduser := ctx.Param("username")
	if len(olduser) == 0 {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("invalid user"))
		return
	}
	fields, ok := ctx.GetQueryArray("fields")
	if !ok {
		fields = append(fields, "username")
	}
	var req sqlc.UpdateUserParams
	req.OldUser = olduser
	for _, v := range fields {
		switch v {
		case "username":
			// TODO: update username such taht dependent tables will also update
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
			req.HashedPassword.String = r.HashedPassword
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
	user, err := s.store.UpdateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, fmt.Errorf("error while updating error %w", err))
		return
	}
	res := userResponse(user)
	ctx.JSON(http.StatusOK, res)
}
