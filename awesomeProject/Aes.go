/*
Program to show off crypto features of golang's crypto library
I actual used this for a class I couldn't get pycrypto installed in time so I used go.
 */
package main
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	key := []byte("13123402445888447145697345789957")
	plaintext := []byte("Built in crypto")

	fmt.Printf("%s\n", plaintext)
	ciphertext, error := encrypt(key, plaintext)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("%0x\n", ciphertext)
	result, error := decrypt(key, ciphertext)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("%s\n", result)
}
func encrypt(key, text []byte) ([]byte, error) {
	block, error := aes.NewCipher(key)
	if error != nil {
		return nil, error
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, error := aes.NewCipher(key)
	if error != nil {
		return nil, error
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, error := base64.StdEncoding.DecodeString(string(text))
	if error != nil {
		return nil, error
	}
	return data, nil
}