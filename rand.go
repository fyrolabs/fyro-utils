package utils

import (
	"encoding/base64"
	"io"
	"time"

	cryptoRand "crypto/rand"
	mathRand "math/rand"
)

var (
	alphaNumLetters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numericLetters  = []rune("0123456789")
)

func RandNumeric(n int) string {
	randSrc := mathRand.NewSource(time.Now().UnixNano())
	random := mathRand.New(randSrc)

	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumLetters[random.Intn(len(numericLetters))]
	}
	return string(b)
}

func RandAlphaNumeric(n int) string {
	randSrc := mathRand.NewSource(time.Now().UnixNano())
	random := mathRand.New(randSrc)

	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumLetters[random.Intn(len(alphaNumLetters))]
	}
	return string(b)
}

type VerifyFunc func(str string) (bool, error)

func RandUniqueAlphaNumeric(
	n int, verify VerifyFunc,
) (string, error) {
	randStr := RandAlphaNumeric(n)
	ok, err := verify(randStr)
	if err != nil {
		return "", err
	}

	if ok {
		return randStr, nil
	}

	return RandUniqueAlphaNumeric(n, verify)
}

func GenerateAESKey(len int) string {
	key := make([]byte, len)
	_, _ = io.ReadFull(cryptoRand.Reader, key)
	return base64.StdEncoding.EncodeToString(key)
}
