package sqlc

import (
	"context"
	"ecom/db/util"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	CreateRandomItem(t)
}

func TestGetItem(t *testing.T) {
	item := CreateRandomItem(t)
	getitem, err := testQueries.GetItemById(context.Background(), item.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, getitem)
	assert.Equal(t, getitem.Category, item.Category)
	assert.Equal(t, getitem.CreatedBy, item.CreatedBy)
	assert.Equal(t, getitem.Name, item.Name)
	assert.Equal(t, getitem.Price, item.Price)
	assert.Equal(t, getitem.ID, item.ID)
	assert.WithinDuration(t, getitem.CreatedAt, item.CreatedAt, time.Minute)
}

func TestDeleteItem(t *testing.T){
	item := CreateRandomItem(t)
	del,err := testQueries.DeleteItem(context.Background(),item.ID)
	assert.NoError(t,err)
	assert.NotEmpty(t,del)
	assert.Equal(t,del.ID,item.ID)
	assert.Equal(t,del.Category,item.Category)
	assert.Equal(t,del.CreatedBy,item.CreatedBy)
	assert.Equal(t,del.Name,item.Name)
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
	assert.NoError(t, err)
	assert.Equal(t, item.Name, req.Name)
	assert.Equal(t, item.Category, req.Category)
	assert.Equal(t, item.CreatedBy, req.CreatedBy)
	assert.Equal(t, item.Price, req.Price)
	return &item
}
