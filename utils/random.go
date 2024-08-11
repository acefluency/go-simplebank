package utils

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// random int (min,max)
func RandomInit(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInit(0, 1000)
}

func RandomMoneyTransfer(balance int64) int64 {
	return RandomInit(0, balance-1)
}

func RandomCurrency() string {
	currencies := []string{"USD", "JPY", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
