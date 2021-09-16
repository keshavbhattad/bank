package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//Generates a random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//Generates a random name
func RandomName() string {
	return RandomString(int(RandomInt(8, 30)))
}

//Generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

//Generates a random currency
func RandomCurrency() string {
	currencies := []string{"INR", "USD", "EUR", "YEN"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
