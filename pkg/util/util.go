package util

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	mathrand "math/rand"
	"time"
)

func Guid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5(base64.URLEncoding.EncodeToString(b))
}

func GenerateRandomString(len int64) string { // GetSjCode

	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	strBytes := []byte(str)

	result := []byte{}
	// time.Sleep(1)
	r := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))

	var i int64
	for i = 0; i < len; i++ {
		// result = append(result, bytes[r.Intn(len(bytes))])  bytes.Count([]byte(str),nil)-1)
		result = append(result, strBytes[r.Intn(bytes.Count(strBytes, nil)-1)])
	}

	return string(result)
}

func GetMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
