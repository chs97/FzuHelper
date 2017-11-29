package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"strconv"
)

// PKCS5Padding DesCBCpadding
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// DesEncryp2Base64 descbc plaintext
func DesEncryp(plainText, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

// GenerateSeed get seed
func GenerateSeed(stdno, date string) (string, error) {
	year := date[0:4]
	month := date[4:6]
	day := date[6:8]
	hour := date[8:10]
	min := date[10:12]
	sec := date[12:14]
	dateStr := []string{year, month, day, hour, min, sec}
	retStr := ""
	m, err := strconv.Atoi(stdno)
	if err != nil {
		return "", err
	}
	m = m%63 + 1
	for i := 0; m > 0; m >>= 1 {
		if m%2 != 0 {
			retStr = dateStr[i] + retStr
		}
		i++
	}
	return retStr, nil
}

// GenerateToken genereate jwt token
func GenerateToken(stdno, date string) (string, error) {
	seed, err := GenerateSeed(stdno, date)
	if err != nil {
		return "", nil
	}
	body := fmt.Sprintf("%s_%s_%s", stdno, date, seed)
	key := []byte("n&1P)J^A")
	iv := key
	plainText := []byte(body)
	result, err := DesEncryp(plainText, key, iv)
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(result), nil

}

// FixedLeadingZero fixed leading zero "1", 2 -> "01"
func FixedLeadingZero(src string, length int) string {
	for i := len(src); i < length; i++ {
		src = "0" + src
	}
	return src
}
