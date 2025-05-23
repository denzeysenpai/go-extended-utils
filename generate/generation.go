package generate

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
)

func GenerateNumberString(limit int) string {
	result := "6"
	for i := 0; i < limit; i++ {
		result += fmt.Sprintf("%d", rand.Intn(10))
	}
	return result
}
func GenerateString(limit int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	runes := make([]rune, limit)

	for i := range runes {
		runes[i] = rune(charset[rand.Intn(len(charset))])
	}

	return string(runes)
}
func GenerateSHA1(fields ...string) string {
	concatenated := ""
	for _, field := range fields {
		concatenated += field
	}
	hasher := sha1.New()
	hasher.Write([]byte(concatenated))
	hash := hasher.Sum(nil)
	return fmt.Sprintf("%x", hash)
}
