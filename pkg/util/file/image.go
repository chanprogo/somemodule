package file

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// GetImageName get image name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = getMd5(fileName)
	return fileName + ext
}

func getMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string, imageAllowExts []string) bool {
	ext := GetExt(fileName)
	for _, allowExt := range imageAllowExts {
		if strings.EqualFold(allowExt, ext) {
			return true
		}
	}
	return false
}

// CheckImageSize check image size
func CheckImageSize(f multipart.File, imageMaxSize int) bool {
	size, err := GetSize(f)
	if err != nil {
		return false
	}
	return size <= imageMaxSize
}

// CheckImage check if the file exists
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	perm := CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	return nil
}
