package util

import (
	"encoding/base64"
	"io"
	"time"

	cryptoRand "crypto/rand"
	mathRand "math/rand"
)

var (
	alphaNumLetters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandAlphaNumeric(n int) string {
	randSrc := mathRand.NewSource(time.Now().UnixNano())
	random := mathRand.New(randSrc)

	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumLetters[random.Intn(len(alphaNumLetters))]
	}
	return string(b)
}

func GenerateAESKey(len int) string {
	key := make([]byte, len)
	_, _ = io.ReadFull(cryptoRand.Reader, key)
	return base64.StdEncoding.EncodeToString(key)
}
