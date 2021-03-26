package middleware

import (
	"fmt"

	"github.com/chanprogo/somemodule/app"
	"github.com/chanprogo/somemodule/pkg/constant"
	"github.com/chanprogo/somemodule/pkg/module/redis/cache"

	"github.com/gin-gonic/gin"
)

// 检查用户权限（1、超时，2、单点登录）
func AuthValidateMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// 获取所有参数
		ctx.Request.ParseForm()
		ctx.Request.ParseMultipartForm(32 << 20) // 32M

		authtoken := ctx.GetHeader("authtoken")
		keyid := ctx.GetHeader("keyid")
		if authtoken == "" || keyid == "" {
			new(app.Controller).RespErr(ctx, nil, "请检查是否登录,TOKEN 或 keyid 为空")
			ctx.Abort()
			return
		}

		cacheKey := fmt.Sprintf("new_token:uid:%s", keyid)
		cachedToken := cache.Get(cacheKey)
		if cachedToken == "" {
			cacheKey = fmt.Sprintf("new_token:uid:%s", keyid) + "watch_pswd" // 观察密码
			cachedToken = cache.Get(cacheKey)
			if cachedToken == "" {
				new(app.Controller).RespErr(ctx, nil, constant.RESPONSE_CODE_SESSION_INVALID)
				ctx.Abort()
				return
			}
		}

		if cachedToken != authtoken { // token 被替换
			cacheKey = fmt.Sprintf("new_token:uid:%s", keyid) + "watch_pswd" // 观察密码
			cachedToken = cache.Get(cacheKey)
			if cachedToken != authtoken {
				new(app.Controller).RespErr(ctx, nil, constant.RESPONSE_CODE_SESSION_REPLACED)
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
