package common

import "github.com/chanprogo/somemodule/pkg/conf/iconf"

// GetImagePath get save path
func GetImagePath() string {
	return iconf.AppSetting.ImageSavePath
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	return iconf.AppSetting.RuntimeRootPath + GetImagePath()
}

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return iconf.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}
