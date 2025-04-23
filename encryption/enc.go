package encryption

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	crypto_rand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

func EncryptAES(key string, data string) (string, error) {
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

func DecryptAES(key string, data string) (string, error) {
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

func EncryptRSA(toEncrypt, filePath string) (string, error) {
	publicKeyFile, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}
	block, _ := pem.Decode(publicKeyFile)
	if block == nil {
		return "", fmt.Errorf("Failed to parse public key PEM")
	}

	pubKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("Failed to parse RSA public key: %v", err)
	}

	publicKey, ok := pubKeyInterface.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("Invalid RSA public key")
	}

	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(toEncrypt))
	if err != nil {
		return "", fmt.Errorf("Encryption failed: %v", err)
	}

	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

func DecryptRSA(toDecrypt, filePath string) (string, error) {
	privateKeyFile, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}

	block, _ := pem.Decode(privateKeyFile)
	if block == nil {
		return "", fmt.Errorf("Failed to parse private key PEM")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("Failed to parse RSA private key: %v", err)
	}

	encryptedData, err := base64.StdEncoding.DecodeString(toDecrypt)
	if err != nil {
		return "", fmt.Errorf("Failed to decode base64 input: %v", err)
	}

	decryptedBytes, err := rsa.DecryptOAEP(crypto.SHA256.New(), rand.Reader, privateKey, encryptedData, nil)
	if err != nil {
		return "", fmt.Errorf("Decryption failed: %v", err)
	}

	return string(decryptedBytes), nil
}
