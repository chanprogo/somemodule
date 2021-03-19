package errors

import (
	"fmt"
	"math/rand"
	"time"

	goErrors "errors"

	"github.com/chanprogo/somemodule/pkg/constant"
	"github.com/chanprogo/somemodule/pkg/util/mathutil"
)

// 系统错误实现
type SysError struct {
	status    int
	simpleMsg string
	fullMsg   string
}

type SysErrorInterface interface {
	Error() string
	String() string
	Status() int
}

// 创建系统错误
func NewSys(options ...interface{}) error {
	var (
		simpleMsg string
		fullMsg   string
	)
	for _, v := range options {
		switch opt := v.(type) {
		default:
			simpleMsg = fmt.Sprintf("系统错误[%s]", "0x"+mathutil.DecimalToAny(rand.New(rand.NewSource(time.Now().UnixNano())).Int(), 16))
			fullMsg = fmt.Sprintf("%s: %v", simpleMsg, opt)
		}
	}
	return &SysError{
		status:    constant.RESPONSE_CODE_SYSTEM,
		simpleMsg: simpleMsg,
		fullMsg:   fullMsg,
	}
}

func (e *SysError) Error() string {
	return e.simpleMsg
}
func (e *SysError) String() string {
	return e.fullMsg
}
func (e *SysError) Status() int {
	return e.status
}

// 普通错误实现
type NormalError struct {
	status int
	msg    string
}

// 普通错误接口
type NormalErrorInterface interface {
	Error() string
	Status() int
}

// 创建普通错误
func NewNormal(options ...interface{}) error {
	var (
		status int
		msg    string
	)
	for _, v := range options {
		switch opt := v.(type) {
		case int:
			status = opt
		case int32:
			status = int(opt)
		case int64:
			status = int(opt)
		default:
			msg = fmt.Sprintf("%v", opt)
		}
	}
	if status == 0 {
		status = constant.RESPONSE_CODE_ERROR
	}
	return &NormalError{
		status: status,
		msg:    msg,
	}
}

func (e *NormalError) Error() string {
	return e.msg
}
func (e *NormalError) Status() int {
	return e.status
}

func GetErrStatus(err interface{}) int32 {
	switch v := err.(type) {
	case SysErrorInterface:
		return int32(v.Status())
	case NormalErrorInterface:
		return int32(v.Status())
	default:
		return int32(constant.RESPONSE_CODE_ERROR)
	}
}

func GetErrMsg(err interface{}) string {
	switch v := err.(type) {
	case SysErrorInterface:
		return v.String()
	case error:
		return v.Error()
	case NormalErrorInterface:
		return v.Error()
	default:
		return ""
	}
}

// 兼容golang errors对象
func New(text string) error {
	return goErrors.New(text)
}
