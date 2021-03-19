package juconf

import (
	"encoding/json"

	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type jsonConfig struct {
	SQLType       string `json:"mysqlType"`
	SQLAddress    string `json:"mysqlAddress"`
	SQLWriteQueue uint32 `json:"mysqlQueue"`

	BoolValue bool `json:"boolValue"`

	MyName string `json:"myName"`
}

// Setting ...
type Setting struct {
	SQLType       string // 数据库类型
	SQLAddress    string // 数据库地址
	SQLWriteQueue uint32 // 数据缓存队列长度

	boolValue bool // 是否打印debug

	myname string // 标识本程序的名字
}

var setting *Setting

// GetSetting ..
func GetSetting() *Setting {
	if setting == nil {
		setting = new(Setting)
	}
	return setting
}

// LoadFormFile ...
func (thi *Setting) LoadFormFile(filePath string) error {

	fileSuffix := path.Ext(filePath)[1:]
	if fileSuffix == "" {
		return errors.New("can not find file suffix")
	}
	fileSuffix = strings.ToLower(fileSuffix)

	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	switch fileSuffix {

	case "json":
		var config jsonConfig
		err = json.Unmarshal(content, &config)
		if err != nil {
			return err
		}

		err = thi.setValues(config.SQLAddress, config.SQLType, config.SQLWriteQueue, config.BoolValue, config.MyName)
		return err

	default:
		return errors.New("Unknow configure file type")

	}
}

// GetSQLAddress ...
func (thi *Setting) GetSQLAddress() string {
	return thi.SQLAddress
}

// GetSQLType ...
func (thi *Setting) GetSQLType() string {
	return thi.SQLType
}

// GetSQLWriteQueue ...
func (thi *Setting) GetSQLWriteQueue() uint32 {
	return thi.SQLWriteQueue
}

// GetBoolValue ...
func (thi *Setting) GetBoolValue() bool {
	return thi.boolValue
}

// GetMyName ...
func (thi *Setting) GetMyName() string {
	return thi.myname
}

func (thi *Setting) setValues(args ...interface{}) error {
	var index int

	thi.SQLAddress = args[index].(string)
	index++
	thi.SQLType = args[index].(string)
	index++

	thi.SQLWriteQueue = args[index].(uint32)
	index++

	thi.boolValue = args[index].(bool)
	index++

	thi.myname = args[index].(string)
	index++
	return nil
}
