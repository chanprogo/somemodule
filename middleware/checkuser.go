package middleware

import (
	"fmt"

	"github.com/chanprogo/somemodule/app"
	"github.com/chanprogo/somemodule/pkg/constant"
	"github.com/chanprogo/somemodule/pkg/log"
	"github.com/chanprogo/somemodule/pkg/module/redis/cache"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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

func AccessTokenMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// 获取所有参数
		c.Request.ParseForm()
		c.Request.ParseMultipartForm(32 << 20) // 32M

		params := c.Request.Form

		authtoken := c.GetHeader("authtoken")
		if authtoken == "" {
			authtoken = params.Get("authtoken")
		}

		if authtoken == "" {
			new(app.Controller).RespErr(c, nil, "请检查是否登录,TOKEN 为空")
			c.Abort()
			return
		}

		keys := "auth_token_uid:" + authtoken
		uid, err := cache.RedisClient.Get(keys).Int64()
		if err == redis.Nil {
			uid = 0
		} else if err != nil {
			log.Logger.Error("redis GetUserId err", keys, uid, err)
			uid = 0
		}

		if uid <= 0 {

			keysw := "auth_token_uid:" + authtoken + "watch_pswd"
			uid, err = cache.RedisClient.Get(keysw).Int64()
			if err == redis.Nil {
				uid = 0
			} else if err != nil {
				log.Logger.Error("redis GetUserId err", keysw, uid, err)
				uid = 0
			}

			if uid <= 0 {
				new(app.Controller).RespErr(c, nil, "请检查是否登录,获取 UID 失败")
				c.Abort()
				return
			}
			c.Header("watch_pswd", "true")
		}

		c.Next()
	}
}
