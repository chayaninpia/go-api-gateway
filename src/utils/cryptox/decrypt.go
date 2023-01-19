package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func DecodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) ([]byte, error) {
	bytes := []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
	block, err := aes.NewCipher([]byte(MySecret[:32]))
	if err != nil {
		return nil, err
	}
	cipherText := DecodeBase64(text)
	plainText := make([]byte, len(cipherText))
	// iv := plainText[:aes.BlockSize]

	cfb := cipher.NewCFBDecrypter(block, bytes)
	cfb.XORKeyStream(plainText, cipherText)

	return plainText, nil
}
