package sqlc

import (
	"ecom/db/util"
	"testing"
)

func TestCreateUser(t *testing.T){
	CreateRandomUser(t)
}

func CreateRandomUser(t *testing.T){
	username := util.RandomUsername()
	full_name := util.RandomFullName(6)
	hashed_password := util.RandomFullName(6)
		
}
