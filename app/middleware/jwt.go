package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/chanprogo/somemodule/pkg/constant"
	"github.com/chanprogo/somemodule/pkg/util/jwtutil"
)

func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {

		var data interface{}
		var code int = constant.RESPONSE_CODE_OK

		token := c.Query("token")

		if token == "" {
			code = constant.INVALID_PARAMS

		} else {

			_, err := jwtutil.ParseToken(token)
			if err != nil {

				switch err.(*jwt.ValidationError).Errors {

				case jwt.ValidationErrorExpired:
					code = constant.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = constant.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != constant.RESPONSE_CODE_OK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  constant.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
