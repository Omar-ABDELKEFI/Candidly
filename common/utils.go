package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/tekab-dev/tekab-test/models"
	"strconv"
	"strings"
)

func ChoiceIndex(tab []models.QuestionChoices, id uint64) int {
	for i := 0; i < len(tab); i++ {
		if tab[i].ChoiceID == id {
			return i
		}
	}
	return -1
}
func GetTestCandidate(idTestcandidat string) (uint64, error, uint64, error) {
	idTest, errIdTest := strconv.ParseUint(idTestcandidat[strings.Index(idTestcandidat, "=")+1:strings.Index(idTestcandidat, "&")], 10, 64)
	idCandidate, errIdCandidate := strconv.ParseUint(idTestcandidat[strings.LastIndex(idTestcandidat, "=")+1:], 10, 64)
	return idTest, errIdTest, idCandidate, errIdCandidate
}

type ErrorResponse struct {
	StructField string
	Tag         string
	Field       string
	Param       string
	Value       interface{}
}

func ValidateStruct(str interface{}) ([]*ErrorResponse, error) {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(str)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.StructField = err.StructField()
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Param = err.Param()
			element.Value = err.Value()
			errors = append(errors, &element)
		}
	}
	return errors, err
}
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func AesEncrypt(orig string, key string) string {
	// Convert to byte array
	origData := []byte(orig)
	k := []byte(key)

	// Group secret key
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(fmt.Sprintf("key The length must be 16/24/32 length: %s", err.Error()))
	}
	// Gets the length of the secret key block
	blockSize := block.BlockSize()
	// Complement code
	origData = PKCS7Padding(origData, blockSize)
	// Encryption mode
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// Create array
	cryted := make([]byte, len(origData))
	// encryption
	blockMode.CryptBlocks(cryted, origData)
	//Use RawURLEncoding instead of StdEncoding
	//Do not use StdEncoding in the url parameter to cause errors
	return base64.RawURLEncoding.EncodeToString(cryted)

}

func AesDecrypt(cryted string, key string) string {
	//Use RawURLEncoding instead of StdEncoding
	//Do not use StdEncoding in the url parameter to cause errors
	crytedByte, _ := base64.RawURLEncoding.DecodeString(cryted)
	k := []byte(key)

	// Group secret key
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(fmt.Sprintf("key The length must be 16/24/32 length: %s", err.Error()))
	}
	// Gets the length of the secret key block
	blockSize := block.BlockSize()
	// Encryption mode
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// Create array
	orig := make([]byte, len(crytedByte))
	// decrypt
	blockMode.CryptBlocks(orig, crytedByte)
	// De completion code
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//Complement
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//De coding
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
