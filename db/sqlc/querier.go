// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUser(ctx context.Context, username string) (User, error)
}

var _ Querier = (*Queries)(nil)
