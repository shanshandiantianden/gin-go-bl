package utils

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

const (
	hash = 10
	N    = 16384
	r    = 8
	p    = 1
)

func ScryptPassword(password string) (out string) {
	salt := make([]byte, 8)
	salt = []byte{1, 2, 3, 4, 5, 6, 7, 8}

	hapw, err := scrypt.Key([]byte(password), salt, N, r, p, hash)
	if err != nil {
		fmt.Println(err.Error())
	}

	eas := base64.StdEncoding.EncodeToString(hapw)
	md5_str := EncrptrMd5(eas)
	out = CBCEncrypter(md5_str, key, iv)

	return
}
