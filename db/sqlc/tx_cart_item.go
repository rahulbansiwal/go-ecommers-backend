package sqlc

import (
	"context"
	"database/sql"
	"strconv"
)

type AddCartItemTxRequest struct {
	ItemId   int32
	Username string
	Quantity int
}

type AddCartItemTxResponse struct {
	CartValue int
}

func (s *SQLStore) AddCartItemTx(ctx context.Context, param AddCartItemTxRequest) (AddCartItemTxResponse, error) {
	var result AddCartItemTxResponse
	if param.Quantity < 1 {
		param.Quantity = 1
	}

	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		// validate item id for a valid item
		item, err := s.GetItemById(ctx, param.ItemId)
		if err != nil {
			return err
		}
		// get cart from username
		cart, err := s.GetCart(ctx, param.Username)
		if err != nil {
			return err
		}
		// get cart items from cart id
		cartItem, err := s.GetCartItemFromCartIDAndItemID(ctx, GetCartItemFromCartIDAndItemIDParams{
			CartID: cart.ID,
			ItemID: param.ItemId,
		})
		updatedCartItem := false
		if err != nil {
			if err == sql.ErrNoRows {
				updatedCartItem = true
				_, err := s.CreateCartItem(ctx, CreateCartItemParams{
					CartID: cart.ID,
					ItemID: param.ItemId,
					Quantity: sql.NullInt32{
						Int32: int32(param.Quantity),
						Valid: true,
					},
				})
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
		if !updatedCartItem {
			_, err := s.UpdateCartItem(ctx, UpdateCartItemParams{
				Quantity: sql.NullInt32{
					Int32: cartItem.Quantity.Int32 + int32(param.Quantity),
					Valid: true,
				},
				ItemID: param.ItemId,
				CartID: cart.ID,
			})
			if err != nil {
				return err
			}
		}
		itemPrice, err := strconv.ParseInt(item.Price, 10, 64)
		if err != nil {
			return err
		}
		amount := cart.TotalValue + int32(param.Quantity)*int32(itemPrice)
		finalCart, err := s.UpdateCartAmount(ctx, UpdateCartAmountParams{
			TotalValue: amount,
			Username:   param.Username,
		})
		if err != nil {
			return err
		}
		result.CartValue = int(finalCart.TotalValue)
		return nil
	})
	return result, err
}
