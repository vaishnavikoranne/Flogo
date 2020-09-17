package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)
var (
	//passphrase = "47cef24b-2b82-4ac4-a27c-fb0aca32baea"
	saltstring = "ac103458-fcb6-41d3-94r0-43d25b4f4ff4"

	ErrInvalidBlockSize = errors.New("invalid blocksize")

	// ErrInvalidPKCS7Data indicates bad input to PKCS7 pad or unpad.
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data (empty or not padded)")

	// ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")

	salt = []byte(saltstring)
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func deriveKey(passphrase []byte, salt []byte) []byte {

	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha1.New)
}
func AESEncrypt(src string, passphrase []byte) ([]byte, []byte) {
	key := deriveKey(passphrase, salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	initialVector := make([]byte, aes.BlockSize)

	ecb := cipher.NewCBCEncrypter(block, []byte(initialVector))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted, initialVector
}

func init() {
	_ = function.Register(&aesencrypt{})
}

type aesencrypt struct {
}

func (s *aesencrypt) Name() string {
	return "aesencrypt"
}

func (aesencrypt) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (aesencrypt) Eval(params ...interface{}) (interface{}, error) {
	
	plainText:=params[0].(string)
	passphrase:=params[1].(string)	

	encryptedData, iv := AESEncrypt(plainText, []byte(passphrase))
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)
	fmt.Println(base64.StdEncoding.EncodeToString(iv) + encryptedString)
	return base64.StdEncoding.EncodeToString(iv) + encryptedString ,nil;


}