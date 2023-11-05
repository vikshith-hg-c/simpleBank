package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqurstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// generate random number
func RandomInt(min, max int64) int64 {
	return (min + rand.Int63n(max-min+1))
}

// generate Random String
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// generate randoem name
func RandomOwner() string {
	return RandomString(8)
}

// generate randoem Money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currency := []string{"INR", "USD", "EUR", "GBP"}
	n := len(currency)
	return currency[rand.Intn(n)]
}
