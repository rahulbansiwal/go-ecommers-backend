// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package sqlc

import (
	"context"
)

type Querier interface {
	AddAddress(ctx context.Context, arg AddAddressParams) (Address, error)
	CreateCart(ctx context.Context, username string) (Cart, error)
	CreateCartItem(ctx context.Context, arg CreateCartItemParams) (CartItem, error)
	CreateItem(ctx context.Context, arg CreateItemParams) (Item, error)
	CreateItemImage(ctx context.Context, arg CreateItemImageParams) (ItemImage, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAddress(ctx context.Context, id int32) (Address, error)
	DeleteCartItem(ctx context.Context, arg DeleteCartItemParams) (CartItem, error)
	DeleteItem(ctx context.Context, id int32) (Item, error)
	DeleteItemImage(ctx context.Context, id int32) (ItemImage, error)
	GetAddresses(ctx context.Context, username string) ([]Address, error)
	GetCart(ctx context.Context, username string) (Cart, error)
	GetItemById(ctx context.Context, id int32) (Item, error)
	GetUser(ctx context.Context, username string) (User, error)
	UpdateAddress(ctx context.Context, arg UpdateAddressParams) (Address, error)
	UpdateCartAmount(ctx context.Context, arg UpdateCartAmountParams) (Cart, error)
	UpdateCartItem(ctx context.Context, arg UpdateCartItemParams) (CartItem, error)
	UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error)
	UpdateItemImageURL(ctx context.Context, arg UpdateItemImageURLParams) (ItemImage, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
