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

type randomaddress struct {
	CountryCode string
	City        string
	Street      string
	Landmark    string
}

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

func RandomUsername() string {
	return fmt.Sprintf("%v@%v.com", randomString(6), randomString(5))
}

func RandomMobileNumber() int64 {
	var result int64
	for i := 0; i < 10; i++ {
		result = result * 10
		n := rand.Int63n(9)
		result += n
	}
	return result
}

func RandomAddressDetails() randomaddress {
	result := randomaddress{
		CountryCode: randomString(3),
		City:        randomString(6),
		Street:      randomString(10),
		Landmark:    randomString(10),
	}
	return result
}
