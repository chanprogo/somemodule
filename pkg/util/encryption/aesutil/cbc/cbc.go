package cbc

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// var key = []byte("yashikejigitcrpy") // 密钥长度只能是 128bit、192bit、256bit 中的一个

//加密
func AesEncryptByKey(key, orig string) string {
	origData := []byte(orig)
	ivkey := []byte("a610a3285c883de2")

	block, _ := aes.NewCipher([]byte(key))                        // 分组秘钥
	blockSize := block.BlockSize()                                // 获取秘钥块的长度
	origData = PKCS7Padding(origData, blockSize)                  // 补全码
	blockMode := cipher.NewCBCEncrypter(block, ivkey[:blockSize]) // 加密模式
	cryted := make([]byte, len(origData))                         // 创建数组
	blockMode.CryptBlocks(cryted, origData)                       // 加密
	return base64.StdEncoding.EncodeToString(cryted)
}

//解密
func AesDecryptByKey(key, cryted string) string {
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	ivkey := []byte("a610a3285c883de2")

	block, _ := aes.NewCipher([]byte(key))                        // 分组秘钥
	blockSize := block.BlockSize()                                // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, ivkey[:blockSize]) // 加密模式
	orig := make([]byte, len(crytedByte))                         // 创建数组
	blockMode.CryptBlocks(orig, crytedByte)                       // 解密
	orig = PKCS7UnPadding(orig)                                   // 去补全码
	return string(orig)
}
