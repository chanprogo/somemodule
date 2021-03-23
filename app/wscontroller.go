package app

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/chanprogo/somemodule/pkg/constant"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type WsController struct {
	Controller
	*melody.Melody
}

func (c *WsController) WsRespOK(data interface{}) []byte {
	resp := &Response{
		Code: constant.RESPONSE_CODE_OK,
		Msg:  "成功",
		Data: data,
	}
	respByte, _ := jsoniter.Marshal(resp)
	return respByte
}

func (c *WsController) WsRespErr(data interface{}, options ...interface{}) []byte {

	resp := &Response{
		Code: constant.RESPONSE_CODE_ERROR,
		Msg:  "",
		Data: data,
	}

	for _, v := range options {
		switch opt := v.(type) {
		case int:
			resp.Code = opt
		case string:
			resp.Msg = opt

		case SysErrorInterface: // 系统错误
			resp.Code = opt.Status()

			if gin.Mode() == gin.ReleaseMode {
				resp.Msg = opt.Error()
			} else {
				resp.Msg = opt.String()
			}
		case NormalErrorInterface: // 常规错误
			if opt.Status() != 0 { // 常规错误指定了code并且不为0
				resp.Code = opt.Status()
			}
			resp.Msg = opt.Error()

		case error:
			resp.Msg = opt.Error()
		}
	}

	sysMsg := constant.GetMsg(resp.Code)
	if len(sysMsg) > 0 {
		resp.Msg = sysMsg
	}
	respByte, _ := jsoniter.Marshal(resp)
	return respByte
}
