package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	crypto_rand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func Encrypt(key string, data string) (string, error) {
	final_key := pad_key(key)
	byteMsg := []byte(shuffle_string(data))
	block, err := aes.NewCipher(final_key)
	if err != nil {
		return "", fmt.Errorf("Could not create new cipher: %v", err)
	}
	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(crypto_rand.Reader, iv); err != nil {
		return "", fmt.Errorf("Could not encrypt: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(key string, data string) (string, error) {
	final_key := pad_key(key)

	cipherText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("Could not base64 decode: %v", err)
	}

	block, err := aes.NewCipher(final_key)
	if err != nil {
		return "", fmt.Errorf("Could not create new cipher: %v", err)
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return shuffle_string(string(cipherText)), nil
}

func pad_key(key string) []byte {
	final_key := key
	if len(key) > 32 {
		final_key = key[:31]
	}
	for len(final_key) < 32 {
		final_key = final_key + "#"
	}
	return []byte(final_key)
}

func shuffle_string(input string) string {
	var shuffle_string string = ""
	for i := len(input) - 1; i > len(input)/2; i-- {
		shuffle_string = shuffle_string + string(input[i])
	}
	for i := (len(input) / 2); i >= 0; i-- {
		shuffle_string = shuffle_string + string(input[i])
	}
	return shuffle_string
}
