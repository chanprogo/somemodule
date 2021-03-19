package middleware

import (
	"fmt"

	"github.com/chanprogo/somemodule/pkg/log"

	"github.com/gin-gonic/gin"
)

func ShowRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Logger.Debug("ShowRequest:", fmt.Sprintf("\n%+v\n", c.Request))
		c.Next()
	}
}
