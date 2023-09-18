package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomFullName(t *testing.T) {
	val := make(map[string]bool)
	for i := 1; i < 10; i++ {
		n := RandomFullName(i)
		assert.Equal(t, val[n], false)
		assert.Equal(t, len(n), i)
		val[n] = true
	}

}

func TestRandomUsername(t *testing.T) {
	val := make(map[string]bool)
	for i := 1; i < 10; i++ {
		n := RandomUsername()
		assert.Equal(t, val[n], false)
		assert.Equal(t, len(n), 16)
		val[n] = true
	}
}

func TestRandomMobileNumber(t *testing.T) {
	n := RandomMobileNumber()
	fmt.Println(n)
	assert.NotEmpty(t,n)
}
