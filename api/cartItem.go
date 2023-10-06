package api

import (
	"ecom/db/sqlc"
	"ecom/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addItemToCartReqeust struct {
	ItemId   int `json:"item_id" binding:"required,numeric"`
	Quantity int `json:"quantity" binding:"numeric"`
}

type addItemToCartResponse struct {
	Message    string `json:"message"`
	Item       int    `json:"item"`
	TotalValue int    `json:"total_value"`
}

func (s *Server) addItemToCart(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var req addItemToCartReqeust
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if req.Quantity < 1 {
		req.Quantity = 1
	}
	cart, err := s.store.AddCartItemTx(ctx, sqlc.AddCartItemTxRequest{
		ItemId:   int32(req.ItemId),
		Quantity: req.Quantity,
		Username: authPayload.Username,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, addItemToCartResponse{
		Message:    "added to cart",
		Item:       req.ItemId,
		TotalValue: cart.CartValue,
	})
}
