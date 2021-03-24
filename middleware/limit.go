package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
)

func LimitMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		limit.NewRateLimiter(func(c *gin.Context) string {
			return "request" // limit all request
		}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
			return rate.NewLimiter(rate.Every(time.Minute), 120), 5 * time.Minute // limit 10 qps/clientIp and permit bursts of at most 10 tokens, and the limiter liveness time duration is 5minute
		}, func(c *gin.Context) {
			c.AbortWithStatus(429) // handle exceed rate limit request
		})

	}

}

func LimitClientIpMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		limit.NewRateLimiter(func(c *gin.Context) string {
			return c.ClientIP() // limit all request
		}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
			return rate.NewLimiter(rate.Every(time.Minute), 120), 5 * time.Minute // limit 10 qps/clientIp and permit bursts of at most 10 tokens, and the limiter liveness time duration is 5minute
		}, func(c *gin.Context) {
			c.AbortWithStatus(429) // handle exceed rate limit request
		})

	}

}
