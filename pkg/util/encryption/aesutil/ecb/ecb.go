package ecb

import (
	"crypto/aes"
	"crypto/cipher"

	"encoding/base64"
	"fmt"
)

// func Base64URLDecode(data string) ([]byte, error) {
// 	var missing = (4 - len(data)%4) % 4
// 	data += strings.Repeat("=", missing)
// 	return base64.URLEncoding.DecodeString(data)
// }

// func Base64UrlSafeEncode(source []byte) string {
// 	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
// 	bytearr := base64.StdEncoding.EncodeToString(source)
// 	safeurl := strings.Replace(string(bytearr), "/", "_", -1)
// 	safeurl = strings.Replace(safeurl, "+", "-", -1)
// 	safeurl = strings.Replace(safeurl, "=", "", -1)
// 	return safeurl
// }

func AesDecrypt(crypted, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("err is:", err)
	}

	origData := make([]byte, len(crypted))

	blockMode := NewECBDecrypter(block)
	blockMode.CryptBlocks(origData, crypted)

	origData = PKCS7UnPadding(origData)
	return origData
}

func AesEncrypt(content, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}

	content = PKCS7Padding(content, block.BlockSize())
	fmt.Printf("补码后字节：%v \n", content)

	crypted := make([]byte, len(content))

	ecb := NewECBEncrypter(block)
	ecb.CryptBlocks(crypted, content)

	return crypted
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func main() {
	oriStr := "1922-001-001-00257"
	oriData := []byte(oriStr)
	fmt.Printf("明文字符串：%v\n", oriStr)
	fmt.Printf("明文字节：%v \n", oriData)

	keyByte := []byte("IJbZX$rLHbNNCDPO") //密钥

	cryptedData := AesEncrypt(oriData, keyByte)
	fmt.Printf("加密后字节：%v \n", cryptedData)
	afterBase64 := base64.StdEncoding.EncodeToString(cryptedData)
	fmt.Printf("加密后（普通base64编码后）：%v\n", afterBase64)
	//fmt.Printf("加密后（Base64UrlSafe）：%v\n", Base64UrlSafeEncode(cryptedData))

	decryptedData := AesDecrypt(cryptedData, keyByte)
	fmt.Printf("解密后字节：%v\n", decryptedData)
	fmt.Printf("解密后字符串：%v\n", string(decryptedData))

	fmt.Println("=================================================================================")
	fmt.Printf("加密后（普通base64编码后）：%v\n", afterBase64)
	decodeBase64ByteSlice, _ := base64.StdEncoding.DecodeString(afterBase64)
	fmt.Printf("普通Base64解码后的字节：%v \n", decodeBase64ByteSlice)

}
