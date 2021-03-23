package app

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取 get、post 提交的参数
// 参数存在时第二个参数返回 true，即使参数的值为空字符串，参数不存在时第二个参数返回 false
func (c *Controller) GetParam(ctx *gin.Context, key string) (string, bool) {
	var param string
	var ok bool
	switch ctx.Request.Method {
	case "POST":
		fallthrough
	case "PUT":
		param, ok = ctx.GetPostForm(key)
	default:
		param, ok = ctx.GetQuery(key)
	}
	return param, ok
}

// 获取 get、post 提交的 string 类型的参数，def 表示默认值，取第一个，多余的丢弃
func (c *Controller) GetString(ctx *gin.Context, key string, def ...string) string {
	param, _ := c.GetParam(ctx, key)
	if len(param) == 0 && len(def) > 0 {
		return def[0]
	}
	return param
}

// 获取 get、post 提交的 int 类型的参数，def 表示默认值，取第一个，多余的丢弃
func (c *Controller) GetInt(ctx *gin.Context, key string, def ...int) (int, error) {
	param, _ := c.GetParam(ctx, key)
	if len(param) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.Atoi(param)
}

// 获取 get、post 提交的 int64 类型的参数，def 表示默认值，取第一个，多余的丢弃
func (c *Controller) GetInt64(ctx *gin.Context, key string, def ...int64) (int64, error) {
	param, _ := c.GetParam(ctx, key)
	if len(param) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseInt(param, 10, 64)
}

// 获取 get、post 提交的 float64 类型的参数，def 表示默认值，取第一个，多余的丢弃
func (c *Controller) GetFloat64(ctx *gin.Context, key string, def ...float64) (float64, error) {
	param, _ := c.GetParam(ctx, key)
	if len(param) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseFloat(param, 64)
}

// 获取 get、post 提交的 float64 类型的参数，def 表示默认值，取第一个，多余的丢弃
func (c *Controller) GetBool(ctx *gin.Context, key string, def ...bool) (bool, error) {
	param, _ := c.GetParam(ctx, key)
	if len(param) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseBool(param)
}
