package api

import (
	"database/sql"
	"ecom/db/sqlc"
	"ecom/token"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createAddressRequest struct {
	FullName     string `json:"full_name" binding:"required"`
	CountryCode  string `json:"country_code" binding:"required,max=3"`
	City         string `json:"city" binding:"required,max=10"`
	Street       string `json:"street" binding:"required"`
	LandMark     string `json:"landmark" binding:"required"`
	MobileNumber string `json:"mobile_number" binding:"required,numeric,len=10"`
}

type createAddressResponse struct {
	message   string
	addressId int
	addrResponse
}
type addrResponse struct {
	FullName     string
	CountryCode  string
	City         string
	Street       string
	LandMark     string
	MobileNumber string
}

func (s *Server) addAddress(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var req createAddressRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	msisdn, err := strconv.ParseInt(req.MobileNumber, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid mobile number")))
		return
	}

	param := sqlc.AddAddressParams{
		Username:     authPayload.Username,
		FullName:     req.FullName,
		CountryCode:  req.CountryCode,
		City:         req.City,
		Street:       req.Street,
		Landmark:     req.LandMark,
		MobileNumber: msisdn,
	}
	address, err := s.store.AddAddress(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res := createAddressResponse{
		message:   "address created succesfully",
		addressId: int(address.ID),
		addrResponse: addrResponse{
			FullName:     address.FullName,
			CountryCode:  address.CountryCode,
			City:         address.City,
			Street:       address.Street,
			LandMark:     address.Landmark,
			MobileNumber: fmt.Sprint(address.MobileNumber),
		},
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetAddress(ctx *gin.Context) {
	authpayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	id := ctx.Param("id")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("addressId is not provided")))
		return
	}
	addressId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid addressId")))
		return
	}
	address, err := s.store.GetaddressById(ctx, int32(addressId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("no address with addressId: %v", addressId)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if address.Username != authpayload.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("unauthorized to access the address")))
		return
	}
	res := addrResponse{
		FullName:     address.FullName,
		CountryCode:  address.CountryCode,
		City:         address.City,
		Street:       address.Street,
		LandMark:     address.Landmark,
		MobileNumber: fmt.Sprint(address.MobileNumber),
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Server) deleteAddress(ctx *gin.Context) {
	authpayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	id := ctx.Param("id")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("address id is not provided")))
		return
	}
	addressId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid address id")))
		return
	}
	addr, err := s.store.GetaddressById(ctx, int32(addressId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("no address associated with address id: %v", addressId)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if addr.Username != authpayload.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("unauthorized to delete the address")))
		return
	}
	deleteaddr,err := s.store.DeleteAddress(ctx,int32(addressId))
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	res := addrResponse{
		FullName: deleteaddr.FullName,
		CountryCode: deleteaddr.CountryCode,
		City: deleteaddr.City,
		Street: deleteaddr.Street,
		LandMark: deleteaddr.Landmark,
		MobileNumber: fmt.Sprint(deleteaddr.MobileNumber),
	}
	ctx.JSON(http.StatusOK,res)
}

