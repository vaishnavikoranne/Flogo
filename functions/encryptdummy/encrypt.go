 package encryptdummy

import (
	/*"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"fmt"*/
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	/*"golang.org/x/crypto/pbkdf2"*/
	
)

type AESEncrypt struct {
}

func init() {
	function.Register(&AESEncrypt{})
}

func (s *AESEncrypt) Name() string {
	return "AESEncrypt"
}



func (s *AESEncrypt) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString, data.TypeString}, false
}

func (s *AESEncrypt) Eval(params ...interface{}) (interface{}, error) {
	plainText := params[0].(string)
	sPassphrase := params[1].(string)
	saltString:=params[2].(string)
	return plainText+sPassphrase+saltString, nil
}


