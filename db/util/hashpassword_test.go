package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	hashpassword(t, RandomString(6))
}

func hashpassword(t *testing.T, password string) string {
	hp, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hp)
	require.GreaterOrEqual(t, len(hp), len(password))
	fmt.Println(hp)
	return hp
}

func TestCheckPassword(t *testing.T) {
	password := RandomString(6)
	hp := hashpassword(t, password)
	err := CheckPassword(password, hp)
	require.NoError(t, err)

}
