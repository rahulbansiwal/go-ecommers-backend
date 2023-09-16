package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	alphabat = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.NewSource(int64(time.Now().UnixNano()))
}

func randomString(n int) string {
	var s strings.Builder
	k := len(alphabat)
	for i := 0; i < n; i++ {
		c := alphabat[rand.Intn(k)]
		s.WriteByte(c)
	}
	return s.String()
}

func RandomFullName(n int) string {
	return randomString(n)
}

func RandomUsername() string{
	return fmt.Sprintf("%v@%v.com",randomString(6),randomString(5))
}

//func RandomMobileNumber()
