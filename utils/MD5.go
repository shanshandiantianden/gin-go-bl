package utils

import (
	"crypto/md5"
	"fmt"
)

func GetMd5(str string) (sMd5 string) {
	data := []byte(str)
	has := md5.Sum(data)
	sMd5 = fmt.Sprintf("%x", has)
	//sMd5 = strings.ToUpper(md5str)
	return
}
func EncrptrMd5(str string) (sMd5 string) {
	data := []byte(str)
	has := md5.Sum(data)
	sMd5 = fmt.Sprintf("%x", has)
	//sMd5 = strings.ToUpper(md5str)
	return
}
