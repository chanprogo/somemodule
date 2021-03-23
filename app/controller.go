package app

import (
	"net/http"

	"github.com/chanprogo/somemodule/pkg/constant"

	"github.com/gin-gonic/gin"
)

const SAVE_DATA_KEY = "save_api_data_key"

type Controller struct {
}

// 设置返回的数据，key-value
// 使用 gin context 的 Keys 保存。gin context 每个请求都会先 reset。
func (c *Controller) Put(ctx *gin.Context, key string, value interface{}) {
	// lazy init
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	if ctx.Keys[SAVE_DATA_KEY] == nil {
		ctx.Keys[SAVE_DATA_KEY] = make(map[string]interface{})
	}
	ctx.Keys[SAVE_DATA_KEY].(map[string]interface{})[key] = value
}

func (c *Controller) RespOK(ctx *gin.Context) {
	resp := &Response{
		Code: constant.RESPONSE_CODE_OK,
		Msg:  "成功",
		Data: ctx.Keys[SAVE_DATA_KEY],
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) RespOKWithMsg(ctx *gin.Context, msg string) {
	resp := &Response{
		Code: constant.RESPONSE_CODE_OK,
		Msg:  msg,
		Data: ctx.Keys[SAVE_DATA_KEY],
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) RespErr(ctx *gin.Context, options ...interface{}) {

	resp := &Response{
		Code: constant.RESPONSE_CODE_ERROR,
		Msg:  "",
		Data: ctx.Keys[SAVE_DATA_KEY],
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
			if opt.Status() != 0 { // 常规错误指定了 code 并且不为0
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
	ctx.JSON(http.StatusOK, resp)
}
