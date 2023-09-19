package sqlc

import (
	"context"
	"ecom/db/util"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	CreateRandomItem(t)
}

func CreateRandomItem(t *testing.T) *Item {
	user := CreateRandomUser(t)
	req := CreateItemParams{
		Name:      util.RandomString(6),
		Price:     fmt.Sprint(util.RandomAmount(100)),
		Category:  util.RandomString(6),
		CreatedBy: user.Username,
	}
	item, err := testQueries.CreateItem(context.Background(), req)
	assert.NoError(t,err)
	assert.Equal(t,item.Name,req.Name)
	assert.Equal(t,item.Category,req.Category)
	assert.Equal(t,item.CreatedBy,req.CreatedBy)
	assert.Equal(t,item.Price,req.Price)
	return &item
}
