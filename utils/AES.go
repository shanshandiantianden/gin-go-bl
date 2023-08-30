package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var (
	//初始密钥
	key []byte
	// 初始向量
	iv []byte
)

func init() {
	src := EncrptrMd5("h1w7j2n1k3h9j8b2v5a4h1w7q1n4x0t9k7a3")
	key = []byte(src[:16])
	sre := EncrptrMd5("h3a3i4h4q1u6w6b2i6y6v5d5w1d8d5x2e6x0")
	iv = []byte(sre[:16])
}

/*
		CBC 加密
		text 待加密的明文
	    key 秘钥
*/
func CBCEncrypter(data string, key []byte, iv []byte) string {
	text := []byte(data)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	// 填充

	paddText := PKCS7Padding(text, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 加密
	result := make([]byte, len(paddText))
	blockMode.CryptBlocks(result, paddText)
	// 返回密文
	//fmt.Println(base64.StdEncoding.EncodeToString(result))
	return base64.StdEncoding.EncodeToString(result)
}

/*
CBC 解密
encrypter 待解密的密文
key 秘钥
*/
func CBCDecrypter(data string, key []byte, iv []byte) string {
	encrypter, _ := base64.StdEncoding.DecodeString(data)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	result := make([]byte, len(encrypter))
	blockMode.CryptBlocks(result, encrypter)
	// 去除填充
	result = UnPKCS7Padding(result)
	return string(result)
}

/*
PKCS7Padding 填充模式
text：明文内容
blockSize：分组块大小
*/
func PKCS7Padding(text []byte, blockSize int) []byte {
	// 计算待填充的长度
	padding := blockSize - len(text)%blockSize
	var paddingText []byte
	if padding == 0 {
		// 已对齐，填充一整块数据，每个数据为 blockSize
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		// 未对齐 填充 padding 个数据，每个数据为 padding
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}
	return append(text, paddingText...)
}

/*
去除 PKCS7Padding 填充的数据
text 待去除填充数据的原文
*/
func UnPKCS7Padding(text []byte) []byte {
	// 取出填充的数据 以此来获得填充数据长度
	unPadding := int(text[len(text)-1])
	return text[:(len(text) - unPadding)]
}
