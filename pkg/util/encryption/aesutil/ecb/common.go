package ecb

import "bytes"

func PKCS7Padding(origData []byte, blockSize int) []byte {
	padding := blockSize - len(origData)%blockSize // 计算需要补几位数
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
