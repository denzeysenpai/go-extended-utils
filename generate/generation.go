package generate

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

func Generate_number_string(limit int) string {
	rand.Seed(time.Now().UnixNano())
	result := "6"
	for i := 0; i < limit; i++ {
		result += fmt.Sprintf("%d", rand.Intn(10))
	}

	return result
}
func Generate_SHA1(fields ...string) string {
	concatenated := ""
	for _, field := range fields {
		concatenated += field
	}
	hasher := sha1.New()
	hasher.Write([]byte(concatenated))
	hash := hasher.Sum(nil)
	return fmt.Sprintf("%x", hash)
}
