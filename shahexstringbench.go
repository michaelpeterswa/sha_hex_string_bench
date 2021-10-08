package shahexstringbench

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ShaHexStringWithHexLib(a string) string {
	shaBytes := sha256.Sum256([]byte(a))
	return hex.EncodeToString(shaBytes[:])
}

func ShaHexStringWithSprintf(a string) string {
	shaBytes := sha256.Sum256([]byte(a))
	return fmt.Sprintf("%x", shaBytes)
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
