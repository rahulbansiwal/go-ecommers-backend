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

type addItemToCartReqeust struct {
	ItemId   int `json:"item_id" binding:"required,numeric"`
	Quantity int `json:"quantity" binding:"numeric"`
}

type addItemToCartResponse struct {
	Message string `json:"message"`
	Item    int    `json:"item"`
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
	item, err := s.store.GetItemById(ctx, int32(req.ItemId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("no item with itemid: %v", req.ItemId)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	cart, err := s.store.GetCart(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	s.store.CreateCartItem(ctx, sqlc.CreateCartItemParams{
		CartID: cart.ID,
		ItemID: int32(req.ItemId),
		Quantity: sql.NullInt32{
			Int32: int32(req.Quantity),
			Valid: true,
		},
	})
	price, _ := strconv.ParseInt(item.Price, 10, 64)
	amount := int64(cart.TotalValue) + (price)*(int64((req.Quantity)))
	_, err = s.store.UpdateCartAmount(ctx, sqlc.UpdateCartAmountParams{
		TotalValue: int32(amount),
		Username:   authPayload.Username,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, addItemToCartResponse{
		Message: "added to cart",
		Item:    req.ItemId,
	})
}

//TODO : Add create cart item check to modify the quantity in db
