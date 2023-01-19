package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"log"
)

func EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(request interface{}, MySecret string) (string, error) {

	bytes := []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
	requestMarshal, err := json.Marshal(&request)
	if err != nil {
		log.Fatalln(err.Error())
	}

	text := string(requestMarshal)

	block, err := aes.NewCipher([]byte(MySecret[:32]))
	if err != nil {
		return "", err
	}

	plainText := []byte(text)
	cipherText := make([]byte, len(plainText))

	cfb := cipher.NewCFBEncrypter(block, bytes)
	cfb.XORKeyStream(cipherText, plainText)
	return EncodeBase64(cipherText), nil
}
