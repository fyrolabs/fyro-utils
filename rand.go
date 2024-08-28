package util

import (
	"time"

	"math/rand"
)

var (
	alphaNumLetters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandAlphaNumeric(n int) string {
	randSrc := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSrc)

	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumLetters[random.Intn(len(alphaNumLetters))]
	}
	return string(b)
}
