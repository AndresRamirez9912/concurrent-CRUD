package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LimitGoroutines() gin.HandlerFunc {
	// The maximum amount of goroutines is the chan size x endpoint
	channel := make(chan bool, 100)

	return func(c *gin.Context) {
		select {
		case channel <- true: // Check randoml
		default:
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		defer func() { <-channel }()

		c.Next()
	}
}
