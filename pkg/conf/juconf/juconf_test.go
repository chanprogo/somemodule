package juconf

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestReadJsonConfig(t *testing.T) {

	var settingFilePath string

	if 1 == len(os.Args) {
		settingFilePath = filepath.Join(filepath.Dir(os.Args[0]), "config.json")
	} else {
		settingFilePath = os.Args[1]
	}

	settingFilePath = "/home/chan/Desktop/workspace/config.json"
	t.Log(settingFilePath)

	setting := GetSetting()
	err := setting.LoadFormFile(settingFilePath)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log("Run successful: " + GetSetting().GetMyName())
	time.Sleep(time.Duration(1) * time.Second)
}
