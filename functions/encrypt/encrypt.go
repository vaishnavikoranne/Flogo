package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"golang.org/x/crypto/pbkdf2"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func deriveKey(passphrase []byte, salt []byte) []byte {

	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha1.New)
}

func AESEncrypt(src string, sPassphrase string, saltString string) string {

	passphrase := []byte(sPassphrase)
	salt := []byte(saltString)

	key := deriveKey(passphrase, salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	//initialVector := make([]byte, aes.BlockSize)
	var initialVector = []byte{34, 35, 35, 57, 68, 4, 35, 36, 7, 8, 35, 23, 35, 86, 35, 23}

	ecb := cipher.NewCBCEncrypter(block, []byte(initialVector))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return base64.StdEncoding.EncodeToString(initialVector) + base64.StdEncoding.EncodeToString(crypted)
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

	plainText := params[0].(string)
	passphrase := params[1].(string)
	saltstring := "ac103458-fcb6-41d3-94r0-43d25b4f4ff4"

	encryptedString := AESEncrypt(plainText, passphrase, saltstring)
	fmt.Println(encryptedString)
	return encryptedString, nil

}
