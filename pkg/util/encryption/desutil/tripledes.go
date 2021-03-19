package common

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

// 加密
func TripleDesEncrypt(orig, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 3DES 的秘钥长度必须为24位
	block, _ := des.NewTripleDESCipher(k)

	origData = PKCS5Padding(origData, block.BlockSize()) // 补全码
	blockMode := cipher.NewCBCEncrypter(block, k[:8])    // 设置加密方式
	crypted := make([]byte, len(origData))               // 创建密文数组
	blockMode.CryptBlocks(crypted, origData)             // 加密

	return base64.StdEncoding.EncodeToString(crypted)
}

// 解密
func TipleDesDecrypt(crypted string, key string) string {

	cryptedByte, _ := base64.StdEncoding.DecodeString(crypted)

	k := []byte(key)

	block, _ := des.NewTripleDESCipher(k)
	blockMode := cipher.NewCBCDecrypter(block, k[:8])
	origData := make([]byte, len(cryptedByte))
	blockMode.CryptBlocks(origData, cryptedByte)
	origData = PKCS5UnPadding(origData)

	return string(origData)
}

func PKCS5Padding(orig []byte, size int) []byte {
	length := len(orig)
	padding := size - length%size
	paddintText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(orig, paddintText...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
